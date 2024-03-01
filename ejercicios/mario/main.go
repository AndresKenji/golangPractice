package main

import (
	"fmt"
	"os"
	"strconv"
)

func Escalera(altura int) {
	for i := 1; i<= altura; i++{
		for j := 1; j<= i; j++{
			fmt.Print("#")
		}
		fmt.Println()
	}
}


func main() {

	if len(os.Args) < 2 {
		fmt.Println(`
		Uso : main n
		n debe ser un numero entre 5 y 20
		`)
		return
	}

	cantidad, err := strconv.Atoi(os.Args[1])
	if cantidad < 5 || cantidad > 20 {
		fmt.Println("Debes introducir un numero valido para la altura de la escalera")
		return
	}
	if err != nil {
		fmt.Println("Debes introducir un numero valido para la altura de la escalera")
		return
	}
	Escalera(cantidad)
}