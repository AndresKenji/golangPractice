package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== EJEMPLOS ADICIONALES DEL MÓDULO OS ===")
	fmt.Println()

	// Ejecutar los ejemplos adicionales
	fmt.Println("Ejecutando ejemplos de archivos avanzados...")
	ejemploArchivosAvanzado()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	fmt.Println("Ejecutando ejemplos de variables de entorno avanzados...")
	ejemploVariablesEntornoAvanzado()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	fmt.Println("Ejecutando ejemplos de flags avanzados...")
	runFlagsExample()
}
