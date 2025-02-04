package main

import "fmt"

// Primera etapa: simplemente envía los números al canal
func stage1(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out) // Cierra el canal cuando termina
	}()
	return out
}

// Segunda etapa: multiplica cada número por 2
func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out) // Cierra el canal cuando termina
	}()
	return out
}

// Tercera etapa: suma 1 a cada número
func stage3(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 1
		}
		close(out) // Cierra el canal cuando termina
	}()
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5} // Datos iniciales
	c1 := stage1(nums)           // Primera etapa
	c2 := stage2(c1)             // Segunda etapa
	c3 := stage3(c2)             // Tercera etapa
	for result := range c3 {     // Recolecta e imprime los resultados finales
		fmt.Println(result)
	}
}
