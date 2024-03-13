package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hola Mundo!")
	src := "F:/Kenji/Libros"
	dest := "C:/Users/andre/Documents"

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

	err = os.MkdirAll(dest, fileInfo.Mode())
	if err != nil {
		fmt.Println("Error al crear el archivo en el destino")
		return
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		fmt.Println("Error al leer el directorio de origen:", err)
		return
	}
	fmt.Print(entries)
	

}
