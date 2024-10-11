package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func copyFile(src, dst string, wg *sync.WaitGroup) {
	defer wg.Done()

	source, err := os.Open(src)
	if err != nil {
		fmt.Println("Error al abrir el archivo de origen:", err)
		return
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error al crear el archivo de destino:", err)
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		fmt.Println("Error al copiar el archivo:", err)
		return
	}
}

func copyDirectory(src, dst string, wg *sync.WaitGroup) {
	defer wg.Done()

	fileInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println("Error al obtener información del archivo:", err)
		return
	}

	if !fileInfo.IsDir() {
		fmt.Println("La ruta de origen no es un directorio.")
		return
	}

	err = os.MkdirAll(dst, fileInfo.Mode())
	if err != nil {
		fmt.Println("Error al crear directorio de destino:", err)
		return
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		fmt.Println("Error al leer el directorio de origen:", err)
		return
	}

	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		destinationPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			wg.Add(1)
			go copyDirectory(sourcePath, destinationPath, wg)
		} else {
			wg.Add(1)
			go copyFile(sourcePath, destinationPath, wg)
		}
	}
}

func main() {

	sigChan := make(chan os.Signal, 1)
	src := "F:/Kenji/Libros"
	dst := "C:/Users/andre/Documents"

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Iniciando proceso")
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-sigChan:
			fmt.Println("\nCopia de archivos cancelada.")
			os.Exit(1)
		}
	}()

	wg.Add(1)
	go copyDirectory(src, dst, &wg)

	wg.Wait()
	fmt.Println("Archivos copiados exitosamente.")
}
