// go get golang.org/x/exp/constraints
package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// func PrintList(list ...interface{}){} // Manera de recibir distiontos tipos de dato desde una función
func PrintList(list ...any) {
	for _, value := range list {
		fmt.Printf("Variable de tipo %T con valor %v \n", value, value)
	}
}

type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

func Suma[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func Filetered[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}

	return result
}

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	PrintList("Andres", 36, 6.5, true)
	fmt.Println(Suma(4, 7, 2, 5))
	fmt.Println(Suma(46.8, 1.7, 34.2, 14.5))

	strings := []string{"a", "b", "c", "d", "e"}
	numbers := []int{4, 6, 8, 3}
	println(Includes(strings, "a"))
	println(Includes(strings, "k"))
	println(Includes(numbers, 8))
	println(Includes(numbers, 20))

	fmt.Println(Filetered(numbers, func(value int) bool { return value < 8 }))
	fmt.Println(Filetered(strings, func(value string) bool { return value > "b" }))

	pruducto1 := Product[uint]{1, "Zapatos", 35}
	pruducto2 := Product[string]{"DRG-ATVVSFD-SDWRFC", "Zapatos", 35}

	fmt.Println(pruducto1,pruducto2)

}
