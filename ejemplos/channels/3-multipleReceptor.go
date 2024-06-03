package main

import (
	"fmt"
	"time"
)

func main() {

	dulces := make(chan string, 10)

	go Engullidor("kenji", dulces)
	go Engullidor("shizu", dulces)
	go Engullidor("chibi", dulces)

	dulces <- "Donnut"
	time.Sleep(time.Second)
	dulces <- "Galleta"
	time.Sleep(time.Second)
	dulces <- "Brownie"
	time.Sleep(time.Second)



}
func Engullidor(nombre string, dulces <-chan string) {
	for dulce := range dulces {
		fmt.Println(nombre,"come",dulce)
	}
}