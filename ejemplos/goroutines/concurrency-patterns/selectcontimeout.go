package main

import (
	"fmt"
	"time"
)

// El uso de la sentencia select con un timeout te permite evitar bloqueos indefinidos.
// Este patrón es útil cuando quieres realizar una acción o abortar si una operación tarda demasiado.

func main() {
	c := make(chan string) // Crear un canal para la comunicación

	go func() {
		time.Sleep(2 * time.Second) // Simular un trabajo que toma tiempo
		c <- "result"               // Enviar un resultado al canal
	}()

	select {
	case res := <-c: // Caso en el que se recibe un mensaje del canal
		fmt.Println("Received:", res) // Imprimir el resultado recibido
	case <-time.After(1 * time.Second): // Caso en el que el timeout se alcanza
		fmt.Println("Timeout") // Imprimir mensaje de timeout
	}
}
