package main

import (
	"fmt"
	"math/rand"
)

func main() {

	numeroMeta := rand.Intn(100)
	intentos := 1
	var numero int
	ganador := false
	fmt.Println("Adivina el numero de 0 a 100 antes de quedarte sin intentos")
	for intentos <= 10 || ganador == true {
		fmt.Println("Cantidad de intentos:", intentos)
		fmt.Print("Adivina el numero:")
		fmt.Scanln(&numero)
		if numero == numeroMeta {
			fmt.Printf("Lo lograste !! en %v intentos, que barbaro ! 😲", intentos)
			ganador = true
			return
		} else {
			if numero < numeroMeta {
				fmt.Println("Un poco mas")
			}
			if numero > numeroMeta {
				fmt.Println("Un poco menos")
			}
		}
		intentos++
	}

	fmt.Println("Que lastima, eso de las matematicas no es lo tuyo 😅")

}
