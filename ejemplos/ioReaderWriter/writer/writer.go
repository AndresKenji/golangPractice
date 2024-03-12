package main

import (
	"fmt"
	"os"
)

func main() {
	archivo, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer archivo.Close()

	mensaje := []byte("Hola, esto es un ejemplo de io.Writer.")
	_, err = archivo.Write(mensaje)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Datos escritos en el archivo correctamente.")
}
