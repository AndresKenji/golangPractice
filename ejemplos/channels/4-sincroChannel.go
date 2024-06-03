package main

import "fmt"

func main() {
	espera := TareaSincrona()

	<- espera

	fmt.Println("Programa finalizado")

}

func TareaSincrona() <-chan struct{} {
	ch := make(chan struct{})
	go func ()  {
		fmt.Println("Haciendo alguna cosa en paralelo...")
		for i := 0; i < 3; i++ {
			fmt.Println(i,"...")
		}
		fmt.Println("finalizada tarea en paralelo")
	
		close(ch)
		
	}()
	return ch

}