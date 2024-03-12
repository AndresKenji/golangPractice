package main

import (
	"fmt"
	"os"
)

func main() {
	archivo, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	buffer := make([]byte, 1024)
	n, err := archivo.Read(buffer)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	datos := buffer[:n]
	fmt.Printf("Datos leídos del archivo: %s\n", datos)
}
