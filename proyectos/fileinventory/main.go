package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type FileData struct {
	Path    string
	Size    int64
	ModTime string
	Mode    string
	IsDir   bool
}

func main() {
	var src string
	var inventory_file string
	flag.StringVar(&src, "dir", "/", "target directory")
	flag.StringVar(&inventory_file, "out-file", "inventory.csv", "output filename")
	flag.Parse()
	inventory, err := os.Create(inventory_file + ".csv")

	writer := csv.NewWriter(inventory)
	defer writer.Flush()

	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer inventory.Close()

	fileInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	if !fileInfo.IsDir() {
		fmt.Println("source route is not a directory")
		return
	}

	files := make(chan FileData)
	var wg sync.WaitGroup

	// Start a goroutine to write to the file
	headers := []string{"name", "size", "timestamp", "permisions"}
	writer.Write(headers)
	go func() {
		for f := range files {
			writer.Write([]string{f.Path, strconv.Itoa(int(f.Size)), f.ModTime, f.Mode})
		}
	}()

	// Start processing the directory
	wg.Add(1)
	go GetDirFiles(src, files, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
	close(files)

	fmt.Println("Inventory done.")
}

func GetDirFiles(src string, archivos chan<- FileData, wg *sync.WaitGroup) {
	defer wg.Done()

	entries, err := os.ReadDir(src)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", src, err)
		return
	}

	for _, entry := range entries {
		fullPath := filepath.Join(src, entry.Name())

		// Get detailed file info
		fileInfo, err := entry.Info()
		if err != nil {
			fmt.Printf("Error getting info for %s: %v\n", fullPath, err)
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
			// Recurse into subdirectories
			wg.Add(1)
			go GetDirFiles(fullPath, archivos, wg)
		}
	}
}
