package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B float64
	const precision = 2

	fmt.Println("Hola, vamos a calcular la hipotenusa y el area del triangulo rectangulo, para eso ingresa lo siguiente:")
	fmt.Print("Introduce el lado A del triangulo: ")
	fmt.Scanln(&A)
	fmt.Print("Introduce el lado B del triangulo: ")
	fmt.Scanln(&B)

	hipotenusa := math.Sqrt(math.Pow(A, 2) + math.Pow(B, 2))
	area := (A * B) / 2
	perimetro := hipotenusa + A + B

	fmt.Printf("Los valores ingresados son A: %v, B: %v \n", A, B)
	fmt.Printf("La hipotenusa es: %v \n", hipotenusa)
	fmt.Printf("El area es: %.2f \n", area)
	fmt.Printf("El perimetro es: %.*f \n", precision, perimetro)

}
