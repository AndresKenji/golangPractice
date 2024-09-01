package main

import (
 "fmt"
 "time"
)
// El Rate Limiting controla la tasa a la que se procesan los eventos usando un ticker.
// Este patrón es útil cuando necesitas controlar la frecuencia de ciertas tareas, como solicitudes a una API.
func main() {
 rate := time.Second          // Definir la tasa de limitación (1 segundo)
 ticker := time.NewTicker(rate) // Crear un ticker que emite ticks cada segundo
 defer ticker.Stop()          // Asegurarse de detener el ticker al final

 requests := make(chan int, 5) // Crear un canal para solicitudes con capacidad de 5
 for i := 1; i <= 5; i++ {
  requests <- i // Enviar 5 solicitudes al canal
 }
 close(requests) // Cerrar el canal de solicitudes

 for req := range requests {
  <-ticker.C // Esperar el siguiente tick
  fmt.Println("Processing request", req) // Procesar la solicitud
 }
}
