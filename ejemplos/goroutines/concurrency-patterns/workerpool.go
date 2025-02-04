package main

import (
	"fmt"
	"sync"
	"time"
)

// El patrón de "Worker Pool" consiste en crear un número fijo de goroutines que procesan tareas desde una cola compartida.
// Este patrón es útil para controlar la cantidad de tareas concurrentes, lo cual es crucial para gestionar el uso de recursos.

// worker Esta función representa un "trabajador".
// Recibe un id, un canal de trabajos (jobs), un canal de resultados (results), y un WaitGroup (wg).
// defer wg.Done(): Indica que la función wg.Done() se ejecutará cuando el worker termine su ejecución.
// Esto disminuye el contador del WaitGroup, señalando que el worker ha completado su trabajo.
// for job := range jobs: Un bucle que itera sobre los trabajos recibidos desde el canal jobs.
// La iteración continúa hasta que el canal jobs se cierra.
// fmt.Printf("Worker %d started job %d\n", id, job): Imprime un mensaje indicando que el worker ha comenzado un trabajo.
// time.Sleep(time.Second): Simula la ejecución del trabajo con un retardo de 1 segundo.
// fmt.Printf("Worker %d finished job %d\n", id, job): Imprime un mensaje indicando que el worker ha terminado el trabajo.
// results <- job * 2: Envía el resultado del trabajo (el doble del número del trabajo) al canal results.
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	// Define el número de trabajos y de workers.
	const numJobs = 5
	const numWorkers = 3
	//  Crea un canal jobs con un buffer que puede contener numJobs trabajos.
	jobs := make(chan int, numJobs)
	// Crea un canal results con un buffer que puede contener numJobs resultados.
	results := make(chan int, numJobs)
	// Declara un WaitGroup que se usa para esperar a que todos los workers terminen sus tareas.
	var wg sync.WaitGroup
	// Lanza numWorkers goroutines, cada una ejecutando la función worker.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Aumenta el contador del WaitGroup antes de lanzar cada worker.
		go worker(i, jobs, results, &wg)
	}
	// Envía numJobs trabajos al canal jobs.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// Cierra el canal jobs indicando que no se enviarán más trabajos.
	close(jobs)
	//  Espera a que todos los workers terminen sus tareas.
	wg.Wait()
	// Cierra el canal results indicando que no se enviarán más resultados.
	close(results)
	// Itera sobre los resultados en el canal results e imprime cada resultado.
	for result := range results {
		fmt.Println("Result:", result)
	}
}
