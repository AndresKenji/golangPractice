package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ============= EJEMPLOS DE GOROUTINES Y CHANNELS =============

func runGoroutineExamples() {
	fmt.Println("=== EJEMPLOS DE GOROUTINES Y CHANNELS EN GO ===")
	fmt.Println()

	// 1. Goroutines básicas
	fmt.Println("1. GOROUTINES BÁSICAS:")
	ejemplosGoroutinesBasicas()
	fmt.Println()

	// 2. Channels básicos
	fmt.Println("2. CHANNELS BÁSICOS:")
	ejemplosChannelsBasicos()
	fmt.Println()

	// 3. Channels con buffer
	fmt.Println("3. CHANNELS CON BUFFER:")
	ejemplosChannelsBuffer()
	fmt.Println()

	// 4. Select statement
	fmt.Println("4. SELECT STATEMENT:")
	ejemplosSelect()
	fmt.Println()

	// 5. WaitGroup
	fmt.Println("5. WAITGROUP:")
	ejemplosWaitGroup()
	fmt.Println()

	// 6. Worker pools
	fmt.Println("6. WORKER POOLS:")
	ejemplosWorkerPools()
	fmt.Println()

	// 7. Fan-in/Fan-out
	fmt.Println("7. FAN-IN/FAN-OUT:")
	ejemplosFanInFanOut()
	fmt.Println()

	// 8. Pipeline pattern
	fmt.Println("8. PIPELINE PATTERN:")
	ejemplosPipeline()
	fmt.Println()

	// 9. Context y cancellation
	fmt.Println("9. CONTEXT Y CANCELLATION:")
	ejemplosContext()
	fmt.Println()

	// 10. Mutex y sincronización
	fmt.Println("10. MUTEX Y SINCRONIZACIÓN:")
	ejemplosMutex()
	fmt.Println()

	fmt.Println("=== FIN DE LOS EJEMPLOS DE GOROUTINES Y CHANNELS ===")
}

// ============= GOROUTINES BÁSICAS =============

func tareaSimple(id int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("    Goroutine %d: iteración %d\n", id, i+1)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("    Goroutine %d terminada\n", id)
}

func ejemplosGoroutinesBasicas() {
	fmt.Println("  Lanzando 3 goroutines:")

	// Lanzar goroutines
	for i := 1; i <= 3; i++ {
		go tareaSimple(i)
	}

	// Esperar un poco para que las goroutines terminen
	time.Sleep(500 * time.Millisecond)
	fmt.Println("  Función principal terminada")
}

// ============= CHANNELS BÁSICOS =============

func enviarDatos(ch chan string, mensaje string) {
	time.Sleep(200 * time.Millisecond)
	ch <- mensaje
}

func ejemplosChannelsBasicos() {
	// Channel simple
	ch := make(chan string)

	go enviarDatos(ch, "Hola desde goroutine")

	// Recibir del channel (bloquea hasta recibir)
	mensaje := <-ch
	fmt.Printf("  Recibido: %s\n", mensaje)

	// Channel para comunicación bidireccional
	numeros := make(chan int)
	resultados := make(chan int)

	// Goroutine que calcula cuadrados
	go func() {
		for num := range numeros {
			resultados <- num * num
		}
		close(resultados)
	}()

	// Enviar números
	go func() {
		for i := 1; i <= 5; i++ {
			numeros <- i
		}
		close(numeros)
	}()

	// Recibir resultados
	fmt.Println("  Cuadrados:")
	for resultado := range resultados {
		fmt.Printf("    %d\n", resultado)
	}
}

// ============= CHANNELS CON BUFFER =============

func ejemplosChannelsBuffer() {
	// Channel con buffer de 3
	ch := make(chan int, 3)

	// Enviar datos sin bloquear (hasta llenar el buffer)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("  Enviados 3 elementos al channel con buffer")

	// Recibir datos
	fmt.Printf("  Recibido: %d\n", <-ch)
	fmt.Printf("  Recibido: %d\n", <-ch)
	fmt.Printf("  Recibido: %d\n", <-ch)

	// Ejemplo de productor-consumidor con buffer
	buffer := make(chan string, 2)
	done := make(chan bool)

	// Productor
	go func() {
		productos := []string{"producto1", "producto2", "producto3", "producto4"}
		for _, producto := range productos {
			buffer <- producto
			fmt.Printf("    Producido: %s\n", producto)
			time.Sleep(100 * time.Millisecond)
		}
		close(buffer)
	}()

	// Consumidor
	go func() {
		for producto := range buffer {
			fmt.Printf("    Consumido: %s\n", producto)
			time.Sleep(200 * time.Millisecond)
		}
		done <- true
	}()

	<-done
}

// ============= SELECT STATEMENT =============

func ejemplosSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutines que envían a diferentes channels
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "mensaje de channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "mensaje de channel 2"
	}()

	// Select para recibir del primer channel disponible
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("  Recibido de ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("  Recibido de ch2: %s\n", msg2)
		}
	}

	// Select con timeout
	fmt.Println("\n  Select con timeout:")
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(300 * time.Millisecond)
		timeout <- true
	}()

	select {
	case <-timeout:
		fmt.Println("    Operación completada")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("    Timeout! Operación cancelada")
	}

	// Select con default (no bloqueante)
	fmt.Println("\n  Select con default:")
	nonBlocking := make(chan string)

	select {
	case msg := <-nonBlocking:
		fmt.Printf("    Recibido: %s\n", msg)
	default:
		fmt.Println("    No hay mensajes disponibles")
	}
}

// ============= WAITGROUP =============

func trabajador(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Marcar como terminado al salir

	fmt.Printf("    Trabajador %d iniciado\n", id)
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	fmt.Printf("    Trabajador %d terminado\n", id)
}

func ejemplosWaitGroup() {
	var wg sync.WaitGroup
	numTrabajadores := 5

	fmt.Printf("  Iniciando %d trabajadores:\n", numTrabajadores)

	for i := 1; i <= numTrabajadores; i++ {
		wg.Add(1) // Incrementar contador
		go trabajador(i, &wg)
	}

	wg.Wait() // Esperar a que todos terminen
	fmt.Println("  Todos los trabajadores terminaron")
}

// ============= WORKER POOLS =============

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("    Worker %d procesando job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2 // Procesar y enviar resultado
	}
}

func ejemplosWorkerPools() {
	numJobs := 9
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Iniciar workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Enviar jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Recoger resultados
	fmt.Println("  Resultados:")
	for r := 1; r <= numJobs; r++ {
		resultado := <-results
		fmt.Printf("    Resultado: %d\n", resultado)
	}
}

// ============= FAN-IN/FAN-OUT =============

// Fan-out: distribuir trabajo a múltiples goroutines
func fanOut(input <-chan int, workers int) []<-chan int {
	outputs := make([]<-chan int, workers)

	for i := 0; i < workers; i++ {
		output := make(chan int)
		outputs[i] = output

		go func(out chan<- int, workerID int) {
			defer close(out)
			for num := range input {
				// Procesar (simular trabajo)
				processed := num * num
				fmt.Printf("    Worker %d: %d -> %d\n", workerID, num, processed)
				time.Sleep(50 * time.Millisecond)
				out <- processed
			}
		}(output, i+1)
	}

	return outputs
}

// Fan-in: combinar outputs de múltiples goroutines
func fanIn(inputs ...<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	// Función para copiar de input a output
	multiplex := func(input <-chan int) {
		defer wg.Done()
		for value := range input {
			output <- value
		}
	}

	// Iniciar una goroutine para cada input
	wg.Add(len(inputs))
	for _, input := range inputs {
		go multiplex(input)
	}

	// Cerrar output cuando todos los inputs terminen
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func ejemplosFanInFanOut() {
	input := make(chan int)

	// Fan-out a 3 workers
	outputs := fanOut(input, 3)

	// Fan-in de todos los outputs
	result := fanIn(outputs...)

	// Enviar datos
	go func() {
		defer close(input)
		for i := 1; i <= 6; i++ {
			input <- i
		}
	}()

	// Recibir resultados combinados
	fmt.Println("  Resultados combinados:")
	for value := range result {
		fmt.Printf("    Resultado final: %d\n", value)
	}
}

// ============= PIPELINE PATTERN =============

// Etapa 1: generar números
func generarNumeros(nums ...int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for _, num := range nums {
			output <- num
		}
	}()
	return output
}

// Etapa 2: calcular cuadrados
func calcularCuadrados(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			output <- num * num
		}
	}()
	return output
}

// Etapa 3: filtrar pares
func filtrarPares(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			if num%2 == 0 {
				output <- num
			}
		}
	}()
	return output
}

func ejemplosPipeline() {
	// Construir pipeline
	numeros := generarNumeros(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	cuadrados := calcularCuadrados(numeros)
	pares := filtrarPares(cuadrados)

	// Procesar resultados
	fmt.Println("  Pipeline: números -> cuadrados -> filtrar pares")
	for resultado := range pares {
		fmt.Printf("    %d\n", resultado)
	}
}

// ============= CONTEXT Y CANCELLATION =============

func tareaConCancelacion(id int, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Printf("    Tarea %d cancelada\n", id)
			return
		default:
			fmt.Printf("    Tarea %d trabajando...\n", id)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func ejemplosContext() {
	done := make(chan bool)

	// Iniciar varias tareas
	for i := 1; i <= 3; i++ {
		go tareaConCancelacion(i, done)
	}

	// Dejar que trabajen un poco
	time.Sleep(600 * time.Millisecond)

	// Cancelar todas las tareas
	close(done)

	// Esperar un poco para ver las cancelaciones
	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Todas las tareas canceladas")
}

// ============= MUTEX Y SINCRONIZACIÓN =============

type Contador struct {
	mu    sync.Mutex
	valor int
}

func (c *Contador) Incrementar() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func (c *Contador) Valor() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.valor
}

func incrementarContador(contador *Contador, veces int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < veces; i++ {
		contador.Incrementar()
	}
}

func ejemplosMutex() {
	contador := &Contador{}
	var wg sync.WaitGroup

	numGoroutines := 5
	incrementosPorGoroutine := 100

	fmt.Printf("  Incrementando contador con %d goroutines, %d veces cada una\n",
		numGoroutines, incrementosPorGoroutine)

	// Lanzar goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go incrementarContador(contador, incrementosPorGoroutine, &wg)
	}

	wg.Wait()

	esperado := numGoroutines * incrementosPorGoroutine
	actual := contador.Valor()

	fmt.Printf("  Valor esperado: %d\n", esperado)
	fmt.Printf("  Valor actual: %d\n", actual)
	fmt.Printf("  ¿Correcto?: %t\n", esperado == actual)

	// Ejemplo de RWMutex
	fmt.Println("\n  Ejemplo con sync.RWMutex:")
	ejemplosRWMutex()
}

// RWMutex permite múltiples lectores simultáneos
type ContadorLectura struct {
	mu    sync.RWMutex
	valor int
}

func (c *ContadorLectura) Leer() int {
	c.mu.RLock() // Bloqueo de lectura
	defer c.mu.RUnlock()
	time.Sleep(10 * time.Millisecond) // Simular trabajo de lectura
	return c.valor
}

func (c *ContadorLectura) Escribir(nuevo int) {
	c.mu.Lock() // Bloqueo exclusivo
	defer c.mu.Unlock()
	time.Sleep(50 * time.Millisecond) // Simular trabajo de escritura
	c.valor = nuevo
}

func ejemplosRWMutex() {
	contador := &ContadorLectura{}
	var wg sync.WaitGroup

	// Múltiples lectores
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			valor := contador.Leer()
			fmt.Printf("    Lector %d: %d\n", id, valor)
		}(i + 1)
	}

	// Un escritor
	wg.Add(1)
	go func() {
		defer wg.Done()
		contador.Escribir(42)
		fmt.Println("    Escritor: valor actualizado a 42")
	}()

	wg.Wait()
}
