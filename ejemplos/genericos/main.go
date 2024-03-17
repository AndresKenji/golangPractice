package main

import "fmt"

//func PrintList(list ...interface{}){} // Manera de recibir distiontos tipos de dato desde una función
func PrintList(list ...any) {
	for _, value := range list {
		fmt.Printf("Variable de tipo %T con valor %v \n", value, value)
	}
}
// https://pkg.go.dev/golang.org/x/exp/constraints

type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

func Suma[T Numbers](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// func Suma[T ~int | ~float64](nums ...T) T {
// 	var total T
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }

func main() {
	PrintList("Andres", 36, 6.5, true)
	fmt.Println(Suma(4,7,2,5))
	fmt.Println(Suma(46.8,1.7,34.2,14.5))


}
