package main

import "fmt"

func main() {
	fmt.Println(suma(6, 8, 23, 546, 132))
	fmt.Println(factorial(8))

	// Funciones anonimas
	func() {
		fmt.Println("Hola desde una función anonima")
	}()

	saludo := func(name string) {
		fmt.Printf("Hola %s como te va ?\n", name)
	}
	saludo("Andres")

}

// Funcion variadica
func suma(nums ...int) int {
	fmt.Println("Total de parametros:", len(nums))
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Función recursiva
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Printf("%v x %v \n", n, n-1)
	return n * factorial(n-1)
}
