package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hola Mundo!")
	src := "C:/Users/OscarAndresRodriguez/Documents/Notebooks"

	fileInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println("Error al obtener información del archivo:", err)
		return
	}
	fmt.Print(fileInfo)
	if !fileInfo.IsDir() {
		fmt.Println("La ruta de origen no es un directorio.")
		return
	}
}