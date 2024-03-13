package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	inventario, err := os.Create("inventario.txt")
	src := "F:/Kenji/Universidad"
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer inventario.Close()

	fileInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println("Error al obtener información del archivo:", err)
		return
	}
	fmt.Print(fileInfo)
	fmt.Println()
	if !fileInfo.IsDir() {
		fmt.Println("La ruta de origen no es una carpeta.")
		return
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		fmt.Println("Error al leer el directorio de origen:", err)
		return
	}

	var archivos = []string{}

	for _, entry := range entries {
		sourcePath := filepath.Join(src, entry.Name())
		if entry.IsDir(){
			GetDirFiles(sourcePath, &archivos)
		} else {
			archivos = append(archivos, sourcePath)
		}

	}

}

func GetDirFiles(src string, archivos *[]string)  {
	var dirFiles = []string{}
	dirEntries, err := os.ReadDir(src)
	fmt.Println("Obteniendo archivos del directorio:",src)
	if err != nil {
		fmt.Println("Error al leer el directorio",err)
		return 
	}
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			dirFiles = append(dirFiles, entry.Name())
		}else {
			fmt.Println("Redireccionando a una nueva carpeta")
		}
	}
	return 	
}