package main

import (
	"fmt"
	"sync"
)

// Fan-Out, Fan-In
// Fan-Out se refiere a cuando se inician múltiples goroutines para procesar datos en paralelo.
// Fan-In se refiere a cuando los resultados de estas goroutines se combinan en un único pipeline.
// Este patrón es útil para el procesamiento en paralelo y luego para la recolección de resultados en un solo flujo de trabajo.

// producer Función que actúa como productor
func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()          // Marca el trabajo como hecho al terminar
	for i := 0; i < 5; i++ { // Produce 5 valores
		ch <- i                                        // Envia cada valor al canal de entrada
		fmt.Printf("Producer %d produced %d\n", id, i) // Imprime lo producido
	}
}

// consumer Función que actúa como consumidor
func consumer(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()     // Marca el trabajo como hecho al terminar
	for v := range in { // Lee valores del canal de entrada
		out <- v * 2                                    // Procesa el valor (lo multiplica por 2) y lo envía al canal de salida
		fmt.Printf("Consumer %d processed %d\n", id, v) // Imprime lo procesado
	}
}

func main() {
	numProducers := 2            // Número de productores
	numConsumers := 2            // Número de consumidores
	input := make(chan int, 10)  // Canal de entrada con un buffer de 10
	output := make(chan int, 10) // Canal de salida con un buffer de 10
	var wg sync.WaitGroup        // Grupo de espera para sincronizar las goroutines

	// Inicia las goroutines de los productores
	for i := 1; i <= numProducers; i++ {
		wg.Add(1)                  // Añade un trabajo al grupo de espera
		go producer(i, input, &wg) // Inicia el productor
	}
	wg.Wait()    // Espera a que todos los productores terminen
	close(input) // Cierra el canal de entrada

	// Inicia las goroutines de los consumidores
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)                          // Añade un trabajo al grupo de espera
		go consumer(i, input, output, &wg) // Inicia el consumidor
	}
	wg.Wait()     // Espera a que todos los consumidores terminen
	close(output) // Cierra el canal de salida

	// Lee los resultados del canal de salida
	for result := range output {
		fmt.Println("Result:", result) // Imprime cada resultado procesado
	}
}
