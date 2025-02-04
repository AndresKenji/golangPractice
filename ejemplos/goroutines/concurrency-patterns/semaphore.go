package main

import (
	"fmt"
	"sync"
	"time"
)

// Un Semáforo limita el número de goroutines que pueden acceder a un recurso particular de manera concurrente.
// Este patrón es útil para controlar la concurrencia y evitar la sobrecarga de recursos.

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- struct{}{} // Adquirir semáforo
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simular trabajo
	fmt.Printf("Worker %d done\n", id)
	<-sem // Liberar semáforo
}

func main() {
	const numWorkers = 5
	const maxConcurrent = 2
	sem := make(chan struct{}, maxConcurrent) // Crear semáforo con capacidad máxima de `maxConcurrent`
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, sem, &wg) // Lanzar goroutines con control de semáforo
	}

	wg.Wait() // Esperar a que todos los workers terminen
}
