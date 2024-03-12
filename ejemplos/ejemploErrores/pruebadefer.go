package main

import "fmt"

func main() {
    fmt.Println("Inicio de la función principal")

    // Defer se utiliza para posponer la ejecución de la función hasta que main haya terminado
    defer fmt.Println("Defer 1")

    // Otras operaciones
    fmt.Println("Operación 1")
    fmt.Println("Operación 2")

    // Defer se utiliza para posponer la ejecución de la función hasta que main haya terminado
    defer fmt.Println("Defer 2")

    fmt.Println("Fin de la función principal")
}
