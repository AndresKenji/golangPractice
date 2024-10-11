package main

import "fmt"

const nums = 3

func Emisor(ch chan<- int) {
	for i := 1; i <= nums; i++ {
		ch <- i
		fmt.Println(i, "Enviado correctamente")
	}
}

func Recerptor(ch <-chan int) {
	for i := 1; i <= nums; i++ {
		num := <-ch
		fmt.Println("Recibido:", num)
	}
}

func main() {
	ch := make(chan int)

	go Emisor(ch)
	Recerptor(ch)
}
