package main

import "fmt"

func cincoVeces(msg string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("(%d de 5) %s\n", i, msg)
	}
}

func main() {
	fmt.Println("Iniciando gorrutina")

	go cincoVeces("Esta gorrutina no siempre se completará")

	cincoVeces("Este mensaje se mostrara cinco veces")

	fmt.Println("Finalizando gorrutina")
}