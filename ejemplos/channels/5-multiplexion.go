package main

import (
	"fmt"
	"time"
)

func CentralMensajeria(sms, email, carta <-chan string) {
	for {
		select {
		case num := <-sms:
			fmt.Println("Recibido SMS del número", num)
		case dir := <-email:
			fmt.Println("Recibido Email de la dirección", dir)
		case rem := <-carta:
			fmt.Println("Recibida Carta del remitente", rem)
		}
	}
}

func main() {
	sms := make(chan string, 5)
	email := make(chan string, 5)
	carta := make(chan string, 5)

	go CentralMensajeria(sms, email, carta)

	sms <- "3016679447"
	email <- "andres.kenji@mail.com"
	carta <- "shizu"
	sms <- "3016847213"

	time.Sleep(time.Second)

}
