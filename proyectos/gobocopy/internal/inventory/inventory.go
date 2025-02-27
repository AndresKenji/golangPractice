package inventory

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

func GetDirFiles(src string, archivos chan<- FileData, wg *sync.WaitGroup) {
    defer wg.Done()

    entries, err := os.ReadDir(src)
    if err != nil {
        log.Printf("Error reading directory %s: %v\n", src, err)
        return
    }

    for _, entry := range entries {
        fullPath := filepath.Join(src, entry.Name())

        fileInfo, err := entry.Info()
        if err != nil {
            log.Printf("Error getting info for %s: %v\n", fullPath, err)
            continue
        }

        data := FileData{
            Path:    fullPath,
            Size:    fileInfo.Size(),
            ModTime: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
            Mode:    fileInfo.Mode().String(),
            IsDir:   fileInfo.IsDir(),
        }

        if !data.IsDir {
            archivos <- data
        } else {
            wg.Add(1)
            go GetDirFiles(fullPath, archivos, wg)
        }
    }
}