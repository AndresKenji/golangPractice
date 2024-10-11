package main

import (
	"fmt"
	//"time"
)

func main() {
	// Crear un canal
	canal := make(chan string)

	// Lanzar una gorutina que envía datos al canal
	go func() {
		canal <- "Hola desde la gorutina."
	}()

	// Leer datos del canal en la gorutina principal
	mensaje := <-canal
	fmt.Println(mensaje)
}
