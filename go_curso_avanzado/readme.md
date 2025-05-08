# SYNC

## ¿Cómo podemos manejar la condición de carrera en Go?
La programación concurrente en Go nos permite realizar múltiples tareas a la vez, incrementando la eficiencia de nuestros programas. Sin embargo, esto también puede llevar a problemas si no se manejan adecuadamente los accesos compartidos a las variables. Uno de estos problemas es la condición de carrera, un fenómeno donde múltiples rutinas acceden y manipulan datos compartidos al mismo tiempo, causando comportamientos incorrectos o inconsistentes.

## ¿Cómo podemos identificar una condición de carrera?
Para empezar, Go ofrece una herramienta para detectar condiciones de carrera. Podemos utilizar la bandera --race al compilar nuestro programa. Esta simple acción nos dará un aviso si el compilador detecta subrutinas accediendo indeciblemente a variables compartidas. La detección temprana es crucial para evitar problemas en producción.

`go run --race main.go`

## ¿Cómo podemos evitar una condición de carrera en Go?
Manejando adecuadamente las condiciones de carrera, podemos asegurar la integridad de los datos. Un enfoque común es utilizar locks o candados. En Go, podemos utilizar la estructura sync.Mutex que actúa como un candado que regula el acceso a las variables compartidas.

### Implementación básica de locks en Go
1. Definir un Mutex: Primero, definimos un sync.Mutex en nuestro programa, que usaremos para controlar el acceso.

`var mu sync.Mutex`

2. Bloquear y desbloquear el acceso: Dentro de las funciones que modifican los datos compartidos, utilizamos los métodos Lock y Unlock de Mutex para bloquear y desbloquear el acceso a las variables.
```go
func depositar(amount int) {
    mu.Lock()
    b := balance
    balance = b + amount
    mu.Unlock()
}
```
>Este código garantiza que solo una go rutina pueda modificar el balance a la vez, evitando por ende la condición de carrera.

## ¿Qué ocurre si solo realizamos lecturas?
Aunque el uso de sync.Mutex soluciona los problemas de concurrencia en operaciones de escritura, puede resultar ineficiente para operaciones que solo requieren lectura. En estos casos, donde no estamos modificando los datos, hay otro tipo de lock en Go más apropiado llamado sync.RWMutex. Este permite múltiples lecturas concurrentes pero restringe el acceso total al realizar una escritura. Lo exploraremos más en profundidad en las siguientes lecciones.

## ¿Cómo manejar lecturas y escrituras concurrentes en Go?
Cuando hablamos de concurrencia en programación, siempre surge la preocupación acerca de las condiciones de carrera, especialmente al manejar variables compartidas. En el contexto de Go, surge una pregunta interesante: ¿Es necesario bloquear las lecturas cuando estamos escribiendo en la misma variable? Es un interrogante esencial al considerar operaciones bancarias, por ejemplo, depósitos simultáneos y consultas de saldo.

Aunque las escrituras deben finalizar para asegurar que los datos sean consistentes, las lecturas no siempre necesitan esperar. Go presenta una solución elegante para manejar múltiples lectores y un único escritor usando un mecanismo llamado RWMutex.

## ¿Cómo implementar un RWMutex en Go?
Un RWMutex es un tipo de log que permite múltiples lectores simultáneos, mientras que asegura que solo un escritor pueda operar en un momento dado. Esto optimiza la concurrencia, ya que no es necesario bloquear las lecturas cuando alguien está escribiendo.

Por ejemplo, al modificar nuestro programa de depósitos, la implementación del RWMutex implica sustituir el mutex en las funciones de lectura y escritura. Veámoslo en un código simplificado:

`var mu sync.RWMutex // Definimos el RWMutex`
```go
func deposit(amount int) {
  mu.Lock() // Bloqueamos la escritura
  balance += amount
  mu.Unlock() // Desbloqueamos después de escribir
}

func getBalance() int {
  mu.RLock() // Bloqueamos solo para lectura
  b := balance
  mu.RUnlock() // Desbloqueamos la lectura
  return b
}
```
Con este enfoque, múltiples goroutines pueden leer el saldo sin bloqueos, aunque solo una pueda escribir a la vez. Esto evita el bloqueo innecesario en las funciones de lectura, mientras que garantiza la seguridad de las escrituras.

## ¿Cuáles son los beneficios del RWMutex?
Implementar un RWMutex en un escenario concurrente trae varios beneficios:

- Concurrencia mejorada: Permite que múltiples procesos lean al mismo tiempo sin interferir.
- Evita bloqueos innecesarios: Las lecturas no bloquean las escrituras, permitiendo consultas más rápidas.
- Seguridad en escritura: Solo se permite una escritura a la vez, previniendo condiciones de carrera.
- Escalabilidad: Especialmente útil en aplicaciones donde las lecturas son mucho más frecuentes que las escrituras.

En resumen, el RWMutex es una herramienta poderosa en Go para optimizar operaciones concurrentes, permitiendo múltiples operaciones de lectura y asegurando la integridad de las escrituras. Este paradigma es ideal para sistemas donde las consultas son prioritarias y se requiere eficiencia y consistencia. Invierte tiempo en interiorizar estos conceptos y apliquémoslos en contextos más complejos, como sistemas de caché concurrente, asegurando programas robustos y eficientes.

## ¿Cómo crear un sistema de caché en Go?
En el desarrollo de sistemas, es común enfrentar problemas con cálculos costosos y que demandan gran cantidad de tiempo para ser realizados. Un buen ejemplo de esto es el cálculo de la serie de Fibonacci para números grandes. Para evitar recalcular continuamente estos valores, se utiliza un sistema de caché, donde almacenamos los resultados de estos cálculos. Si el mismo cálculo debe repetirse, simplemente se busca el resultado en el caché. Aquí te mostraré cómo construir un sistema de caché concurrente en Go.

## ¿Cómo definir la función de Fibonacci?
Antes de iniciar con el caché, es fundamental establecer la función de Fibonacci. En esta ocasión, se utiliza una versión recursiva que es computacionalmente intensiva:
```go

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```
## ¿Qué estructuras de datos necesitas en Go?
En Go, las estructuras (structs) son equivalentes a las clases en otros lenguajes. Aquí, la estructura memory utiliza funciones para gestionar los cálculos de Fibonacci. Además, mapeamos enteros a resultados utilizando un mapa (map).
```go
type Function func(int) (interface{}, error)

type FunctionResult struct {
    value interface{}
    err   error
}

type Memory struct {
    cache map[int]FunctionResult
    f     Function
}

```
## ¿Cómo crear un constructor para el caché?
El constructor se encarga de inicializar nuestra estructura de memoria y preparar el mapa para almacenar los resultados.
```go

func NewCache(f Function) *Memory {
    return &Memory{
        cache: make(map[int]FunctionResult),
        f:     f,
    }
}
```
## ¿Cómo implementar el acceso al caché?
El método get del caché verifica si el resultado ya existe y lo devuelve; de lo contrario, calcula, almacena y luego retorna el valor.
```go

func (m *Memory) get(key int) (interface{}, error) {
    if result, exists := m.cache[key]; exists {
        return result.value, result.err
    }

    value, err := m.f(key)
    m.cache[key] = FunctionResult{value, err}
    return value, err
}
```
## ¿Cómo ejecutar y medir el rendimiento?
Se integran funciones adicionales para probar varios valores y medir el rendimiento mediante el tiempo que toma cada cálculo.
```go
func main() {
    cache := NewCache(getFibonacci)
    numbers := []int{42, 40, 41, 42, 38}

    for _, n := range numbers {
        start := time.Now()
        value, err := cache.get(n)
        
        if err != nil {
            fmt.Printf("Hubo un error: %v\n", err)
            continue
        }
        
        fmt.Printf("Fib(%d) = %v, calculated in %v\n", n, value, time.Since(start))
    }
}

```
Puedes observar cómo los cálculos posteriores para el mismo número son prácticamente instantáneos, lo que muestra la eficiencia del sistema de caché. construcción de un sistema de caché como este, no solo mejora el rendimiento, sino que también optimiza recursos, haciendo que tus aplicaciones sean mucho más eficientes.

## ¿Cómo resolver problemas de concurrencia en un sistema de caché para la serie de Fibonacci?
La implementación de un sistema de caché eficiente para la serie de Fibonacci puede parecer simple en escenarios básicos, pero se vuelve un desafío cuando se introducen procesos concurrentes. Imagínate una situación donde varios procesos simultáneos solicitan el cálculo de valores de Fibonacci para números grandes que, de no estar bien gestionados, podrían sobrecargar tu sistema. En este contexto, es esencial desarrollar un sistema que no solo gestione la caché eficientemente, sino que también identifique y gestione las operaciones que están en progreso y aquellas que ya están listas para entrega. En esta guía, veremos cómo implementar un sistema de caché con manejo de procesos concurrentes, tomando como referencia el lenguaje de programación Go.

## ¿Cómo simular un cálculo intensivo de Fibonacci?
Para simular un cálculo intensivo, se puede emplear una función que representa un tiempo de ejecución prolongado. La función Fibonacci no solo debe calcular el valor, sino también emular tiempo de procesamiento significativo:
```go
func calcularFibonacci(n int) int {
    fmt.Printf("Calculando Fibonacci costoso para %d\n", n)
    time.Sleep(5 * time.Second) // Simula un cálculo largo
    return n
}

```
## ¿Cómo definir estructuras y mapas para gestionar trabajos en proceso?
Para manejar múltiples procesos concurrentes, es crucial definir estructuras y mapas que gestionen qué trabajos se están procesando actualmente y cuáles están esperando una respuesta:
```go
type Servicio struct {
    enProgreso map[int]bool
    enEspera   map[int][]chan int
    mu         sync.RWMutex // candado para evitar condiciones de carrera
}
```
enProgreso: Mapa que indica si un trabajo está en proceso.
enEspera: Mapa que almacena canales de respuesta para trabajos que esperan resultados.

## ¿Cómo implementar la función de trabajo para gestionar procesos concurrentes?
El método trabajar se encarga de manejar las operaciones concurrentes. Comprueba si un trabajo está en progreso y actúa consecuentemente:
```go

func (s *Servicio) trabajar(trabajo int) {
    s.mu.RLock()
    if s.enProgreso[trabajo] {
        // Si está en progreso, agrega el trabajador al mapa de espera
        ch := make(chan int, 1)
        s.mu.RUnlock()
        
        // Operaciones que involucran bloqueo total
        s.mu.Lock()
        s.enEspera[trabajo] = append(s.enEspera[trabajo], ch)
        s.mu.Unlock()

        // Espera la respuesta
        respuesta := <-ch
        fmt.Printf("Respuesta recibida: %d\n", respuesta)
        return
    }
    s.mu.RUnlock()

    // Si no está en progreso, calcúlelo y notifique a los trabajadores
    s.mu.Lock()
    s.enProgreso[trabajo] = true
    s.mu.Unlock()

    fmt.Printf("Calculando Fibonacci para %d\n", trabajo)
    resultado := calcularFibonacci(trabajo)

    s.mu.Lock()
    for _, ch := range s.enEspera[trabajo] {
        ch <- resultado
        close(ch)
    }
    s.enProgreso[trabajo] = false
    s.enEspera[trabajo] = nil
    s.mu.Unlock()
}
```
## ¿Cómo instanciar el servicio y gestionar trabajos concurrentemente?
Para poner a prueba nuestro sistema, definimos un conjunto de trabajos y los procesamos concurrentemente utilizando sync.WaitGroup:
```go

func main() {
    servicio := &Servicio{
        enProgreso: make(map[int]bool),
        enEspera:   make(map[int][]chan int),
    }

    trabajos := []int{3, 4, 5, 5, 4, 8, 8, 8}
    wg := &sync.WaitGroup{}
    wg.Add(len(trabajos))

    for _, trabajo := range trabajos {
        go func(trabajo int) {
            defer wg.Done()
            servicio.trabajar(trabajo)
        }(trabajo)
    }
    
    wg.Wait()
}
```
Con esta implementación, hemos creado un sistema de caché que maneja de forma efectiva procesos concurrentes al gestionar trabajos en progreso y esperando resultados. Te animo a implementar tus propias mejoras y explorar combinaciones con otras técnicas de caché para crear una solución aún más robusta. ¡Sigue avanzando y experimentando con nuevas ideas!