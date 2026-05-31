package config

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Options struct {
	Source string
	Dest   string

	Threads     int
	Mirror      bool
	Restartable bool
	BackupMode  bool
	Compress    bool

	PreserveSecurity bool
	PreserveOwner    bool

	ExcludeExt     []string
	IncludeExt     []string
	ExcludePattern []string
	MinSize        int64
	MaxSize        int64
	MinAge         time.Duration
	MaxAge         time.Duration

	Retries    int
	RetryDelay time.Duration

	LogFile     string
	ReportJSON  string
	FailOnError bool
	DryRun      bool
}

func Parse(args []string) (Options, error) {
	var opts Options

	fs := flag.NewFlagSet("gobocopy", flag.ContinueOnError)
	fs.StringVar(&opts.Source, "source", "", "Directorio origen")
	fs.StringVar(&opts.Dest, "dest", "", "Directorio destino")

	fs.IntVar(&opts.Threads, "threads", 8, "Numero de workers concurrentes")
	fs.BoolVar(&opts.Mirror, "mirror", false, "Refleja origen en destino eliminando extras")
	fs.BoolVar(&opts.Restartable, "restartable", true, "Reanuda archivos parcialmente copiados")
	fs.BoolVar(&opts.BackupMode, "backup", true, "Modo tolerante a fallos de permisos")
	fs.BoolVar(&opts.Compress, "compress", false, "Guarda archivos comprimidos en formato .gz")

	fs.BoolVar(&opts.PreserveSecurity, "preserve-security", true, "Preserva ACL/seguridad si el sistema lo permite")
	fs.BoolVar(&opts.PreserveOwner, "preserve-owner", runtime.GOOS != "windows", "Preserva uid/gid en sistemas Unix")

	excludeExt := fs.String("exclude-ext", "", "Extensiones a excluir separadas por coma, ejemplo: .tmp,.log")
	includeExt := fs.String("include-ext", "", "Extensiones permitidas separadas por coma")
	excludePattern := fs.String("exclude-pattern", "", "Patrones glob de exclusión separados por coma")
	minSize := fs.String("min-size", "", "Tamano minimo en bytes o con sufijos K,M,G")
	maxSize := fs.String("max-size", "", "Tamano maximo en bytes o con sufijos K,M,G")
	minAge := fs.String("min-age", "", "Antiguedad minima, ejemplo: 24h")
	maxAge := fs.String("max-age", "", "Antiguedad maxima, ejemplo: 720h")

	fs.IntVar(&opts.Retries, "retries", 3, "Reintentos por archivo ante error")
	fs.DurationVar(&opts.RetryDelay, "retry-delay", 2*time.Second, "Espera base entre reintentos")

	fs.StringVar(&opts.LogFile, "log-file", "gobocopy.log", "Archivo de log detallado")
	fs.StringVar(&opts.ReportJSON, "report-json", "", "Ruta opcional para guardar reporte JSON")
	fs.BoolVar(&opts.FailOnError, "fail-on-error", false, "Devuelve error final si algun archivo falla")
	fs.BoolVar(&opts.DryRun, "dry-run", false, "Simula operaciones sin modificar archivos")

	if err := fs.Parse(args); err != nil {
		return opts, err
	}

	if opts.Source == "" || opts.Dest == "" {
		return opts, errors.New("debes indicar -source y -dest")
	}

	var err error
	opts.Source, err = filepath.Abs(opts.Source)
	if err != nil {
		return opts, fmt.Errorf("source invalido: %w", err)
	}
	opts.Dest, err = filepath.Abs(opts.Dest)
	if err != nil {
		return opts, fmt.Errorf("dest invalido: %w", err)
	}

	if opts.Threads < 1 {
		opts.Threads = 1
	}

	if opts.ExcludeExt, err = parseCSVExtensions(*excludeExt); err != nil {
		return opts, err
	}
	if opts.IncludeExt, err = parseCSVExtensions(*includeExt); err != nil {
		return opts, err
	}
	opts.ExcludePattern = parseCSV(*excludePattern)

	if *minSize != "" {
		opts.MinSize, err = parseSize(*minSize)
		if err != nil {
			return opts, fmt.Errorf("min-size invalido: %w", err)
		}
	}
	if *maxSize != "" {
		opts.MaxSize, err = parseSize(*maxSize)
		if err != nil {
			return opts, fmt.Errorf("max-size invalido: %w", err)
		}
	}

	if *minAge != "" {
		opts.MinAge, err = time.ParseDuration(*minAge)
		if err != nil {
			return opts, fmt.Errorf("min-age invalido: %w", err)
		}
	}
	if *maxAge != "" {
		opts.MaxAge, err = time.ParseDuration(*maxAge)
		if err != nil {
			return opts, fmt.Errorf("max-age invalido: %w", err)
		}
	}

	if opts.MinSize > 0 && opts.MaxSize > 0 && opts.MinSize > opts.MaxSize {
		return opts, errors.New("min-size no puede ser mayor que max-size")
	}
	if opts.MinAge > 0 && opts.MaxAge > 0 && opts.MinAge > opts.MaxAge {
		return opts, errors.New("min-age no puede ser mayor que max-age")
	}

	return opts, nil
}

func parseCSV(raw string) []string {
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func parseCSVExtensions(raw string) ([]string, error) {
	items := parseCSV(raw)
	if len(items) == 0 {
		return nil, nil
	}
	out := make([]string, 0, len(items))
	for _, item := range items {
		if !strings.HasPrefix(item, ".") {
			return nil, fmt.Errorf("la extension %q debe iniciar con punto", item)
		}
		out = append(out, strings.ToLower(item))
	}
	return out, nil
}

func parseSize(raw string) (int64, error) {
	raw = strings.TrimSpace(strings.ToUpper(raw))
	if raw == "" {
		return 0, nil
	}

	mult := int64(1)
	switch {
	case strings.HasSuffix(raw, "K"):
		mult = 1024
		raw = strings.TrimSuffix(raw, "K")
	case strings.HasSuffix(raw, "M"):
		mult = 1024 * 1024
		raw = strings.TrimSuffix(raw, "M")
	case strings.HasSuffix(raw, "G"):
		mult = 1024 * 1024 * 1024
		raw = strings.TrimSuffix(raw, "G")
	}

	base, err := strconv.ParseInt(strings.TrimSpace(raw), 10, 64)
	if err != nil {
		return 0, err
	}
	if base < 0 {
		return 0, errors.New("tamano negativo")
	}
	return base * mult, nil
}
