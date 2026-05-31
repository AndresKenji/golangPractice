package engine

import (
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"gobocopy/internal/config"
	"gobocopy/internal/platform"
)

type Report struct {
	StartedAt      time.Time         `json:"startedAt"`
	FinishedAt     time.Time         `json:"finishedAt"`
	Duration       string            `json:"duration"`
	ScannedFiles   int64             `json:"scannedFiles"`
	CopiedFiles    int64             `json:"copiedFiles"`
	SkippedFiles   int64             `json:"skippedFiles"`
	FailedFiles    int64             `json:"failedFiles"`
	DeletedFiles   int64             `json:"deletedFiles"`
	DeletedDirs    int64             `json:"deletedDirs"`
	CopiedBytes    int64             `json:"copiedBytes"`
	FailedList     []string          `json:"failedList,omitempty"`
	ConfigSnapshot map[string]string `json:"configSnapshot"`
}

type Engine struct {
	opts config.Options
	log  *safeLogger

	copiedFiles  atomic.Int64
	skippedFiles atomic.Int64
	failedFiles  atomic.Int64
	deletedFiles atomic.Int64
	deletedDirs  atomic.Int64
	copiedBytes  atomic.Int64
	scannedFiles atomic.Int64

	failedMu   sync.Mutex
	failedList []string
}

type copyTask struct {
	SourcePath string
	DestPath   string
	RelPath    string
	Info       fs.FileInfo
}

func Run(opts config.Options) (Report, error) {
	eng, err := newEngine(opts)
	if err != nil {
		return Report{}, err
	}
	defer eng.log.Close()

	report, runErr := eng.run()
	if writeErr := writeReportJSON(opts.ReportJSON, report); writeErr != nil {
		eng.log.Errorf("no se pudo guardar reporte JSON: %v", writeErr)
		if runErr == nil {
			runErr = writeErr
		}
	}

	return report, runErr
}

func newEngine(opts config.Options) (*Engine, error) {
	if err := os.MkdirAll(opts.Dest, 0o755); err != nil {
		return nil, fmt.Errorf("crear destino: %w", err)
	}

	lg, err := newSafeLogger(opts.LogFile)
	if err != nil {
		return nil, err
	}

	return &Engine{opts: opts, log: lg}, nil
}

func (e *Engine) run() (Report, error) {
	started := time.Now()
	e.log.Infof("inicio de copia source=%s dest=%s threads=%d mirror=%v compress=%v dry-run=%v", e.opts.Source, e.opts.Dest, e.opts.Threads, e.opts.Mirror, e.opts.Compress, e.opts.DryRun)

	tasks, sourceSet, err := e.buildTasks()
	if err != nil {
		return e.makeReport(started), err
	}

	e.log.Infof("archivos detectados para evaluar: %d", len(tasks))

	if e.opts.Mirror {
		if err := e.mirrorDelete(sourceSet); err != nil {
			e.log.Warnf("mirror con errores: %v", err)
		}
	}

	e.executeWorkers(tasks)

	report := e.makeReport(started)
	if e.failedFiles.Load() > 0 && e.opts.FailOnError {
		return report, fmt.Errorf("finalizado con errores en %d archivo(s)", e.failedFiles.Load())
	}

	e.log.Infof("fin de copia scanned=%d copied=%d skipped=%d failed=%d deletedFiles=%d deletedDirs=%d bytes=%d", report.ScannedFiles, report.CopiedFiles, report.SkippedFiles, report.FailedFiles, report.DeletedFiles, report.DeletedDirs, report.CopiedBytes)
	return report, nil
}

func (e *Engine) buildTasks() ([]copyTask, map[string]struct{}, error) {
	tasks := make([]copyTask, 0, 256)
	sourceSet := make(map[string]struct{}, 512)

	err := filepath.WalkDir(e.opts.Source, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		rel, err := filepath.Rel(e.opts.Source, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)

		if rel == "." {
			sourceSet[rel] = struct{}{}
			return nil
		}

		if d.IsDir() {
			sourceSet[rel] = struct{}{}
			if !e.opts.DryRun {
				dstDir := filepath.Join(e.opts.Dest, rel)
				if err := os.MkdirAll(dstDir, 0o755); err != nil {
					return fmt.Errorf("crear directorio destino %s: %w", dstDir, err)
				}
			}
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		e.scannedFiles.Add(1)
		if !e.shouldInclude(rel, info) {
			e.skippedFiles.Add(1)
			return nil
		}

		dst := e.destinationPath(rel)
		sourceSet[filepath.ToSlash(dst)] = struct{}{}
		tasks = append(tasks, copyTask{
			SourcePath: path,
			DestPath:   filepath.Join(e.opts.Dest, dst),
			RelPath:    rel,
			Info:       info,
		})

		return nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("recorrer source: %w", err)
	}

	return tasks, sourceSet, nil
}

func (e *Engine) shouldInclude(rel string, info fs.FileInfo) bool {
	ext := strings.ToLower(filepath.Ext(rel))
	if len(e.opts.IncludeExt) > 0 {
		allowed := false
		for _, candidate := range e.opts.IncludeExt {
			if ext == candidate {
				allowed = true
				break
			}
		}
		if !allowed {
			return false
		}
	}

	for _, blocked := range e.opts.ExcludeExt {
		if ext == blocked {
			return false
		}
	}

	for _, pattern := range e.opts.ExcludePattern {
		ok, err := filepath.Match(pattern, filepath.Base(rel))
		if err == nil && ok {
			return false
		}
	}

	size := info.Size()
	if e.opts.MinSize > 0 && size < e.opts.MinSize {
		return false
	}
	if e.opts.MaxSize > 0 && size > e.opts.MaxSize {
		return false
	}

	if e.opts.MinAge > 0 || e.opts.MaxAge > 0 {
		age := time.Since(info.ModTime())
		if e.opts.MinAge > 0 && age < e.opts.MinAge {
			return false
		}
		if e.opts.MaxAge > 0 && age > e.opts.MaxAge {
			return false
		}
	}

	return true
}

func (e *Engine) destinationPath(rel string) string {
	if e.opts.Compress {
		return rel + ".gz"
	}
	return rel
}

func (e *Engine) mirrorDelete(sourceSet map[string]struct{}) error {
	e.log.Infof("mirror activo: analizando eliminaciones")
	var mirrorErr error

	_ = filepath.WalkDir(e.opts.Dest, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			mirrorErr = errors.Join(mirrorErr, walkErr)
			return nil
		}

		rel, err := filepath.Rel(e.opts.Dest, path)
		if err != nil {
			mirrorErr = errors.Join(mirrorErr, err)
			return nil
		}
		rel = filepath.ToSlash(rel)
		if rel == "." {
			return nil
		}

		if _, ok := sourceSet[rel]; ok {
			return nil
		}

		if e.opts.DryRun {
			e.log.Infof("dry-run delete %s", path)
			if d.IsDir() {
				e.deletedDirs.Add(1)
			} else {
				e.deletedFiles.Add(1)
			}
			return nil
		}

		if d.IsDir() {
			if err := os.RemoveAll(path); err != nil {
				mirrorErr = errors.Join(mirrorErr, err)
				return nil
			}
			e.deletedDirs.Add(1)
			e.log.Infof("delete dir %s", path)
			return filepath.SkipDir
		}

		if err := os.Remove(path); err != nil {
			mirrorErr = errors.Join(mirrorErr, err)
			return nil
		}
		e.deletedFiles.Add(1)
		e.log.Infof("delete file %s", path)
		return nil
	})

	return mirrorErr
}

func (e *Engine) executeWorkers(tasks []copyTask) {
	ch := make(chan copyTask)
	wg := sync.WaitGroup{}

	for i := 0; i < e.opts.Threads; i++ {
		workerID := i + 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range ch {
				e.processTask(workerID, task)
			}
		}()
	}

	for _, task := range tasks {
		ch <- task
	}
	close(ch)
	wg.Wait()
}

func (e *Engine) processTask(workerID int, task copyTask) {
	if e.isUpToDate(task) {
		e.skippedFiles.Add(1)
		e.log.Infof("worker=%d skip up-to-date %s", workerID, task.RelPath)
		return
	}

	attempts := e.opts.Retries + 1
	var lastErr error
	for attempt := 1; attempt <= attempts; attempt++ {
		err := e.copyOne(task)
		if err == nil {
			e.copiedFiles.Add(1)
			e.log.Infof("worker=%d copied %s", workerID, task.RelPath)
			return
		}
		lastErr = err
		e.log.Warnf("worker=%d error %s attempt=%d/%d: %v", workerID, task.RelPath, attempt, attempts, err)

		if attempt < attempts {
			time.Sleep(e.opts.RetryDelay * time.Duration(attempt))
		}
	}

	e.failedFiles.Add(1)
	e.failedMu.Lock()
	e.failedList = append(e.failedList, fmt.Sprintf("%s: %v", task.RelPath, lastErr))
	e.failedMu.Unlock()
}

func (e *Engine) isUpToDate(task copyTask) bool {
	info, err := os.Stat(task.DestPath)
	if err != nil {
		return false
	}
	if task.Info.Size() != info.Size() && !e.opts.Compress {
		return false
	}
	if info.ModTime().Before(task.Info.ModTime()) {
		return false
	}
	return true
}

func (e *Engine) copyOne(task copyTask) error {
	if e.opts.DryRun {
		e.log.Infof("dry-run copy %s -> %s", task.SourcePath, task.DestPath)
		e.copiedBytes.Add(task.Info.Size())
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(task.DestPath), 0o755); err != nil {
		return fmt.Errorf("mkdir destino: %w", err)
	}

	tempPath := task.DestPath + ".part"

	written, srcHash, dstHash, err := e.copyFileWithSafety(task.SourcePath, tempPath, task.Info.Size(), e.opts.Compress)
	if err != nil {
		if e.opts.BackupMode {
			_ = os.Chmod(task.DestPath, 0o666)
		}
		return err
	}

	if !e.opts.Compress && srcHash != dstHash {
		return fmt.Errorf("integridad invalida hash src=%s dst=%s", srcHash, dstHash)
	}

	if err := os.Rename(tempPath, task.DestPath); err != nil {
		return fmt.Errorf("rename final: %w", err)
	}

	if err := platform.ApplyMetadata(task.SourcePath, task.DestPath, task.Info, e.opts.PreserveSecurity, e.opts.PreserveOwner); err != nil {
		e.log.Warnf("metadatos no aplicados completamente %s: %v", task.RelPath, err)
	}

	e.copiedBytes.Add(written)
	return nil
}

func (e *Engine) copyFileWithSafety(src, dst string, srcSize int64, compress bool) (int64, string, string, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, "", "", fmt.Errorf("abrir source: %w", err)
	}
	defer srcFile.Close()

	var offset int64
	if e.opts.Restartable && !compress {
		if info, err := os.Stat(dst); err == nil {
			offset = info.Size()
			if offset > srcSize {
				offset = 0
			}
		}
	}

	if offset > 0 {
		if _, err := srcFile.Seek(offset, io.SeekStart); err != nil {
			return 0, "", "", fmt.Errorf("seek source: %w", err)
		}
	}

	flags := os.O_CREATE | os.O_WRONLY
	if offset > 0 {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}
	dstFile, err := os.OpenFile(dst, flags, 0o644)
	if err != nil {
		return 0, "", "", fmt.Errorf("abrir temp destino: %w", err)
	}
	defer dstFile.Close()

	srcHash := sha256.New()
	dstHash := sha256.New()
	var writer io.Writer
	var gzWriter *gzip.Writer

	if compress {
		gzWriter = gzip.NewWriter(dstFile)
		writer = gzWriter
	} else {
		writer = io.MultiWriter(dstFile, dstHash)
	}

	if offset > 0 && !compress {
		if err := seedHashesForResume(src, dst, offset, srcHash, dstHash); err != nil {
			return 0, "", "", fmt.Errorf("reanudar hash: %w", err)
		}
	}

	buffer := make([]byte, 8*1024*1024)
	tee := io.TeeReader(srcFile, srcHash)
	written, err := io.CopyBuffer(writer, tee, buffer)
	if err != nil {
		return 0, "", "", fmt.Errorf("copy: %w", err)
	}

	if gzWriter != nil {
		if err := gzWriter.Close(); err != nil {
			return 0, "", "", fmt.Errorf("cerrar gzip: %w", err)
		}
	}

	if err := dstFile.Sync(); err != nil {
		return 0, "", "", fmt.Errorf("sync temp: %w", err)
	}

	totalWritten := written + offset
	return totalWritten, hex.EncodeToString(srcHash.Sum(nil)), hex.EncodeToString(dstHash.Sum(nil)), nil
}

func seedHashesForResume(src, dst string, offset int64, srcHash, dstHash io.Writer) error {
	if offset == 0 {
		return nil
	}
	if err := hashPrefix(src, offset, srcHash); err != nil {
		return err
	}
	if err := hashPrefix(dst, offset, dstHash); err != nil {
		return err
	}
	return nil
}

func hashPrefix(path string, limit int64, w io.Writer) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.CopyN(w, f, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}

func (e *Engine) makeReport(started time.Time) Report {
	ended := time.Now()
	report := Report{
		StartedAt:    started,
		FinishedAt:   ended,
		Duration:     ended.Sub(started).String(),
		ScannedFiles: e.scannedFiles.Load(),
		CopiedFiles:  e.copiedFiles.Load(),
		SkippedFiles: e.skippedFiles.Load(),
		FailedFiles:  e.failedFiles.Load(),
		DeletedFiles: e.deletedFiles.Load(),
		DeletedDirs:  e.deletedDirs.Load(),
		CopiedBytes:  e.copiedBytes.Load(),
		ConfigSnapshot: map[string]string{
			"source":      e.opts.Source,
			"dest":        e.opts.Dest,
			"threads":     fmt.Sprint(e.opts.Threads),
			"mirror":      fmt.Sprint(e.opts.Mirror),
			"restartable": fmt.Sprint(e.opts.Restartable),
			"backup":      fmt.Sprint(e.opts.BackupMode),
			"compress":    fmt.Sprint(e.opts.Compress),
			"dryRun":      fmt.Sprint(e.opts.DryRun),
		},
	}
	if len(e.failedList) > 0 {
		report.FailedList = append(report.FailedList, e.failedList...)
	}
	return report
}

func writeReportJSON(path string, report Report) error {
	if strings.TrimSpace(path) == "" {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(report)
}

type safeLogger struct {
	mu sync.Mutex
	f  *os.File
}

func newSafeLogger(path string) (*safeLogger, error) {
	if path == "" {
		path = "gobocopy.log"
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, fmt.Errorf("abrir log: %w", err)
	}
	return &safeLogger{f: f}, nil
}

func (l *safeLogger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.f.Close()
}

func (l *safeLogger) Infof(format string, args ...any) {
	l.write("INFO", format, args...)
}

func (l *safeLogger) Warnf(format string, args ...any) {
	l.write("WARN", format, args...)
}

func (l *safeLogger) Errorf(format string, args ...any) {
	l.write("ERROR", format, args...)
}

func (l *safeLogger) write(level, format string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	line := fmt.Sprintf(format, args...)
	ts := time.Now().Format(time.RFC3339)
	_, _ = fmt.Fprintf(l.f, "%s [%s] %s\n", ts, level, line)
	_, _ = fmt.Printf("%s [%s] %s\n", ts, level, line)
}
