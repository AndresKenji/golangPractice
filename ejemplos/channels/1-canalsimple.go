package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "hola mundo enviado a un canal"
	}()
	recibido := <-ch
	fmt.Println("Se ha recibido:", recibido)
}
