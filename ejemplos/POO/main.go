package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== PROGRAMACIÓN ORIENTADA A OBJETOS EN GO ===")
	fmt.Println()

	// 1. Ejemplo básico con interfaces
	fmt.Println("1. INTERFACES Y POLIMORFISMO:")
	ejemploInterfaces()
	fmt.Println()

	// 2. Composición vs herencia
	fmt.Println("2. COMPOSICIÓN (EMBEDDING):")
	ejemploComposicion()
	fmt.Println()

	// 3. Estructuras anidadas
	fmt.Println("3. ESTRUCTURAS ANIDADAS:")
	ejemploEstructurasAnidadas()
	fmt.Println()

	// 4. Encapsulación
	fmt.Println("4. ENCAPSULACIÓN:")
	ejemploEncapsulacion()
	fmt.Println()

	// 5. Métodos con diferentes receivers
	fmt.Println("5. MÉTODOS Y RECEIVERS:")
	ejemploMetodos()
	fmt.Println()

	// 6. Interfaces vacías y type assertions
	fmt.Println("6. INTERFACES VACÍAS Y TYPE ASSERTIONS:")
	ejemploInterfazVacia()
	fmt.Println()

	// 7. Patrón Strategy
	fmt.Println("7. PATRÓN STRATEGY:")
	ejemploPatronStrategy()
	fmt.Println()

	// 8. Patrón Observer
	fmt.Println("8. PATRÓN OBSERVER:")
	ejemploPatronObserver()
	fmt.Println()

	// 9. Sistema completo de biblioteca
	fmt.Println("9. SISTEMA DE BIBLIOTECA (EJEMPLO COMPLETO):")
	ejemploSistemaBiblioteca()
	fmt.Println()

	fmt.Println("=== FIN DE LOS EJEMPLOS DE POO ===")
}
