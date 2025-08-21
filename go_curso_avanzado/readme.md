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

## ¿Qué son los patrones de diseño y por qué son importantes?
Los patrones de diseño son conceptos fundamentales en el desarrollo de software. Imagínalos como planos que nos ayudan a resolver problemas comunes de diseño sin necesidad de reinventar la rueda. Estos patrones son esenciales para convertirte en un ingeniero más habilidoso y eficiente. Aunque no son fragmentos de código específicos, proporcionan guías generales que puedes adaptar a tus necesidades específicas.

El término "patrones de diseño" fue introducido por Christopher Alexander, un arquitecto que exploró su aplicación en la construcción y luego fue adoptado en el ámbito del software. Un grupo de cuatro autores, conocidos como el "Gang of Four", adaptaron este concepto al software, especialmente en programación orientada a objetos. En su libro, presentaron 23 patrones de diseño variados en tres categorías principales: creacionales, estructurales y de comportamiento.

## ¿Cuáles son las categorías de los patrones de diseño?

### Patrones creacionales: ¿cómo se crean nuevos objetos?
Los patrones creacionales se centran en mecanismos para crear objetos de forma flexible y reutilizable, lo que es clave en un software maduro. Este tipo de patrones facilita la generación de instancias de nuevas clases de manera que el código sea reutilizable y mantenga las mejores prácticas en ingeniería de software. En este curso, estudiarás dos patrones creacionales: Factory y Singleton.

### Patrones estructurales: ¿cómo se construyen las estructuras de clases?
Los patrones estructurales ayudan a definir cómo crear objetos a partir de clases más grandes, aplicando herencia o composición. Esto se hace siempre buscando maximizar la flexibilidad y reusabilidad del código. En este módulo, analizaremos el patrón Adapter, que es fundamental para establecer cómo deberían configurarse las relaciones entre objetos y clases.

### Patrones de comportamiento: ¿cómo se comunican los objetos?
Los patrones de comportamiento son esenciales para definir cómo los objetos de diferentes clases se comunican de manera efectiva. Aquí la idea es establecer métodos de comunicación que cumplan un objetivo común. En este curso, revisaremos los patrones Observer y Strategy, ambos cruciales para entender y mejorar las interacciones entre objetos.

## ¿Por qué es fundamental implementar patrones de diseño en tu lenguaje de programación?
Implementar patrones de diseño en tu lenguaje de programación, como Go en este caso, te permitirá aprovechar las ventajas y abordar las limitaciones específicas de cada lenguaje. Aprenderás a adaptar estos patrones universales a las particularidades de Go, lo que enriquecerá tu capacidad de diseñar soluciones elegantes y eficientes. Así que, prepárate para aplicar estos valiosos conceptos y elevar tu habilidad como desarrollador. ¡Adelante, el siguiente paso es la implementación práctica!

## Factory
Para comenzar, crearemos una interfaz básica y las estructuras necesarias implementándolas en lenguaje Go. Veamos el proceso paso a paso:

Crear la interfaz base: Primero, definimos una interfaz, llamada iProduct, que exigirá la implementación de varios métodos esenciales:
```go
type iProduct interface {
    setStock(stock int)
    getStock() int
    setName(name string)
    getName() string
}
```
Definir estructuras concretas: Ahora, definimos nuestra estructura Computer y sus métodos para cumplir con la interfaz iProduct.
```go
type Computer struct {
    name  string
    stock int
}

func (c *Computer) setStock(stock int) { c.stock = stock }
func (c *Computer) getStock() int { return c.stock }
func (c *Computer) setName(name string) { c.name = name }
func (c *Computer) getName() string { return c.name }
```
Composición sobre herencia: Creamos subclases como Laptop y Desktop, usando composición en Go en lugar de herencia.
```go
type Laptop struct {
    Computer
}

func newLaptop() iProduct {
    return &Laptop{Computer{name: "Laptop Computer", stock: 25}}
}

type Desktop struct {
    Computer
}

func newDesktop() iProduct {
    return &Desktop{Computer{name: "Desktop Computer", stock: 35}}
}
```
Implementar el Factory: El siguiente paso es crear la función getComputerFactory que utiliza el patrón Factory para devolver el producto adecuado.
```go
func getComputerFactory(computerType string) (iProduct, error) {
    if computerType == "laptop" {
        return newLaptop(), nil
    } else if computerType == "desktop" {
        return newDesktop(), nil
    }
    return nil, fmt.Errorf("Invalid computer type")
}
```
## ¿Por qué usar el patrón Factory?
El patrón Factory nos ofrece múltiples beneficios:

- Flexibilidad: Nos permite ampliar el número de productos a crear sin modificar el código existente.
- Polimorfismo: Todas las instancias son tratadas como iProduct, promoviendo el uso del polimorfismo y la reutilización de código.
- Mantenimiento: Centraliza y simplifica la gestión de instancias permitiendo modificaciones en un solo lugar.

## ¿Cómo probar y verificar la implementación?
Crear una función de impresión para visualizar los resultados y verificar el comportamiento polimorfo:
```go
func printNameAndStock(p iProduct) {
    fmt.Printf("Product Name: %s, Stock: %d\n", p.getName(), p.getStock())
}

func main() {
    laptop, _ := getComputerFactory("laptop")
    desktop, _ := getComputerFactory("desktop")

    printNameAndStock(laptop)
    printNameAndStock(desktop)
}
```
Ejecutando este código comprobarás que las instancias de Laptop y Desktop son tratadas adecuadamente como instancias de iProduct.

## ¿Cuáles son los pros y contras del patrón Factory?
El uso del patrón Factory ofrece claridad en la creación de objetos y flexibilidad para añadir nuevos productos. Sin embargo, puede hacer que el código sea más complejo y difícil de leer, especialmente para novatos en programación.

## Singleton
El patrón de diseño Singleton es ampliamente utilizado en programación, especialmente cuando se requiere manejar una única instancia de una clase a lo largo de toda la aplicación. Esto es particularmente relevante cuando hablamos de conexiones a bases de datos, donde tener múltiples instancias podría ser ineficiente o incluso problemático. Este patrón asegura que una clase tenga solo una instancia y proporciona un punto de acceso global a ella.

## ¿Por qué utilizar Singleton en Go?
En Go, el patrón Singleton es crucial para garantizar que solo una instancia de un objeto se cree, lo cual es esencial en situaciones como:

- Conexiones a bases de datos: evitar múltiples conexiones que consumen recursos.
- Control de acceso a recursos compartidos: asegurar que varias partes de un sistema accedan al mismo recurso sin conflicto.
- Mantener estados globales: cuando se necesita persistir valores a través de sesiones.
Go, aunque funciona de manera distinta a otros lenguajes de programación, permite implementar este patrón eficientemente utilizando mecanismos como "locks" para asegurar que las instancias no se dupliquen.

## ¿Cómo implementar Singleton en Go?
Para implementar el patrón Singleton en Go, desarrollamos una función que gestiona la creación y el acceso a la instancia única de la base de datos. Aquí te mostramos un ejemplo simplificado de cómo lograrlo:
```go

package main

import (
	"fmt"
	"sync"
	"time"
)

// Estructura Database como placeholder para la conexión.
type Database struct{}

// Variable para almacenar la instancia única.
var dbInstance *Database
var lock = &sync.Mutex{}

// Función que devuelve la única instancia.
func getDatabaseInstance() *Database {
	// Utilizamos un lock para asegurar que solo un acceso ocurra a la vez.
	lock.Lock()
	defer lock.Unlock()
	
	// Creamos la instancia si aún no existe.
	if dbInstance == nil {
		fmt.Println("Creando la conexión a la base de datos...")
		dbInstance = &Database{}
		createSingleConnection()
	} else {
		fmt.Println("Instancia ya creada, usando la existente.")
	}
	return dbInstance
}

// Simula la creación de una conexión lenta.
func createSingleConnection() {
	fmt.Println("Creando la conexión única para la base de datos...")
	time.Sleep(2 * time.Second)
	fmt.Println("Conexión creada.")
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer waitGroup.Done()
			getDatabaseInstance()
		}()
	}

	waitGroup.Wait()
}
```
## ¿Cómo asegurarse de que solo se cree una instancia?
En el ejemplo dado, utilizamos un sync.Mutex que actúa como un mecanismo de bloqueo para garantizar que solo una goroutine pueda evaluar y crear la instancia al mismo tiempo. Este enfoque es crucial, especialmente en aplicaciones concurrentes como aquellas que usan goroutines en Go.

## Consideraciones al usar Singleton
- Eficiencia: Usar Singleton puede mejorar la eficiencia al reducir instancias repetidas de objetos costosos o recursos compartidos.
- Concurrencia: Especialmente en Go, manejar la concurrencia con mecanismos adecuadamente bloqueados es esencial para evitar condiciones de carrera.
- Flexibilidad: Una vez implementado, el patrón debe manejar situaciones donde la inicialización del objeto pueda fallar o requiera reiniciarse.

## Adapter
El patrón de diseño Adapter es un patrón estructural que sirve para resolver problemas de incompatibilidad entre interfaces y estructuras (structs) en programación. Este patrón actúa como un puente entre dos interfaces incompatibles, permitiendo que trabajen juntas sin necesidad de modificar el código existente o agregar líneas innecesarias. Esto es importante porque evita reescribir o refactorizar código que ya está bien definido y probado.

## ¿Cómo se implementa un Adapter en Go?
Para implementar un Adapter en el lenguaje de programación Go, seguiremos los pasos descritos en el siguiente ejemplo:

Definir la interfaz principal: Crear una interfaz que especifique el comportamiento deseado. Por ejemplo:
```go
type Payment interface {
    Pay()
}
```
Crear un struct que implemente esta interfaz: Este struct debe tener la implementación necesaria para la interfaz:
```go
type CashPayment struct{}

func (c CashPayment) Pay() {
    fmt.Println("Pagando utilizando cash")
}
```
Añadir un nuevo struct con comportamiento diferente: Supongamos que se necesita un struct que requiere un parámetro adicional:
```go
type BankPayment struct{}

func (b BankPayment) Pay(accountNumber int) {
    fmt.Printf("Pay using bank account %d\n", accountNumber)
}
```
Crear un adapter para integrar el nuevo struct con la interfaz existente: Se crea un nuevo tipo que adaptará el comportamiento del struct al de la interfaz:
```go
type BankPaymentAdapter struct {
    bankPayment BankPayment
    bankAccount int
}

func (bpa BankPaymentAdapter) Pay() {
    bpa.bankPayment.Pay(bpa.bankAccount)
}
```
Usar el adapter en lugar del nuevo struct cuando sea necesario: Finalmente, incluimos el Adapter en nuestro flujo para garantizar compatibilidad:
```go
func main() {
    cash := CashPayment{}
    processPayment(cash)

    bankAdapter := BankPaymentAdapter{
        bankPayment: BankPayment{},
        bankAccount: 12345,
    }

    processPayment(bankAdapter)
}
```
## ¿Cuál es la utilidad del patrón Adapter?
El uso de un Adapter tiene varias ventajas, entre las que destacan:

-Evitar rehacer el código existente: Adaptar un nuevo componente a una interfaz ya existente sin tener que cambiar el código preexistente.
-Reutilización de código: Permitir el uso de un código nuevo que no implementa interfaces determinadas, favoreciendo la cohesión y flexibilidad.
-Eliminación de redundancias: No será necesario crear métodos duplicados o redefinir interfaces específicas para cada nuevo comportamiento.

Implementar el patrón de diseño Adapter en las situaciones correctas puede simplificar significativamente tu código y mejorar la modularidad. Si todavía no has experimentado con este patrón, te animo a hacerlo. Comienza a adaptarlo en tus proyectos y observa cómo tu código se convierte en una solución más elegante y eficaz. ¡Sigue aprendiendo y explorando patrones de diseño que incrementarán tus habilidades de desarrollo!


## Observer
El Observer es un patrón de diseño de comportamiento que se utiliza para permitir que un objeto informe a otros objetos sobre cambios en su estado. Este patrón resulta muy útil para situaciones donde se necesita notificar a múltiples objetos cuando ocurre un evento específico, todo sin causar problemas de rendimiento.

## ¿Por qué es útil este patrón?
La aplicación de Observer es particularmente beneficiosa al evitar que los objetos tengan que preguntar repetidamente por el estado de un evento. En lugar de ello, el objeto que sufre el cambio notifica directamente a los observadores interesados, mejorando así la eficiencia.

## ¿Cómo implementar el patrón Observer en Go?
Para demostrar la implementación de Observer en Go, crearemos un sistema de notificación similar al de una tienda en línea, notificando a los clientes cuando un producto vuelva a estar disponible.

## Definición del tópico y el observador
Vamos a comenzar definiendo una interfaz Topic y un tipo Observer. Aquí está el esqueleto básico:
```go

package main

type Observer interface {
    GetID() string
    UpdateValue(value string)
}

type Topic interface {
    Register(observer Observer)
    Broadcast()
}

type Item struct {
    observers []Observer
    name      string
    available bool
}
```
## Creación de un nuevo ítem
Creamos un constructor para nuestros ítems, representando un producto en la tienda.
```go

func NewItem(name string) *Item {
    return &Item{
        name:      name,
        available: false,
    }
}

func (i *Item) UpdateAvailable() {
    fmt.Printf("El ítem %s está disponible\n", i.name)
    i.available = true
    i.Broadcast()
}
```
## Implementación de métodos del ítem
Ahora, implementamos los métodos Register y Broadcast para gestionar los observadores.
```go

func (i *Item) Register(observer Observer) {
    i.observers = append(i.observers, observer)
}

func (i *Item) Broadcast() {
    for _, observer := range i.observers {
        observer.UpdateValue(i.name)
    }
}
```
## Definición de un observador
Creamos una estructura que represente a los observadores, en este caso, un sistema de notificación por email.
```go

type EmailClient struct {
    id string
}

func (e *EmailClient) GetID() string {
    return e.id
}

func (e *EmailClient) UpdateValue(value string) {
    fmt.Printf("Enviando email: El ítem %s está disponible; Notificado al cliente %s\n", value, e.id)
}
```
## Configuración del main
Finalmente, en la función main, creamos un ítem, agregamos observadores y actualizamos la disponibilidad del ítem.
```go

func main() {
    nvidiaItem := NewItem("RTX 3080")
 
    firstObserver := &EmailClient{id: "1234"}
    secondObserver := &EmailClient{id: "5678"}
    
    nvidiaItem.Register(firstObserver)
    nvidiaItem.Register(secondObserver)
    
    nvidiaItem.UpdateAvailable()
}
```
Ejecutar este código enviará notificaciones a todos los observadores registrados cuando el ítem se vuelva disponible.

## Ventajas de usar Observer
- Reducción de consultas constantes: No es necesario que los observadores pregunten constantemente si un evento ha ocurrido.
- Organización de código: El código se estructura de manera más clara y fácil de mantener, ya que separa claramente los objetos notificados de los observadores.
- Escalabilidad: Permite agregar nuevos observadores sin modificar el objeto observado.

El patrón Observer es una poderosa herramienta para manejar suscripciones y notificaciones en cualquier programa. Su uso puede mejorar significativamente el rendimiento y la estructura de tus aplicaciones. ¡Anímate a implementarlo y observa los beneficios!

## Strategy
En el mundo de la programación, la implementación de patrones de diseño puede marcar una diferencia sustancial en la claridad y eficiencia del código. El patrón de diseño Strategy, conocido por su estilo de comportamiento, ofrece una forma excelente de definir una serie de algoritmos de manera que se puedan intercambiar sin cambiar el código que los utiliza. Este enfoque no solo mejora la organización del código, sino que también facilita su mantenimiento y escalabilidad. Vamos a explorar cómo este patrón puede ser implementado en Go, lo que permite la creación de aplicaciones limpias y modulares.

## ¿Cómo se estructura el patrón Strategy en Go?
Para implementar el patrón Strategy en Go, primero creamos un archivo nuevo denominado strategy.go. Comenzamos por definir el paquete main y luego procedemos a crear un nuevo struct llamado PasswordProtector. Este actuaría como un contenedor para diferentes algoritmos de hachís o hashing, plataforma a través de la cual el patrón Strategy muestra todo su potencial. Aquí se muestra cómo se implementa:
```go

package main

type PasswordProtector struct {
    user       string
    passwordName string
    hashAlgo   HashAlgorithm
}

type HashAlgorithm interface {
    Hash(p PasswordProtector)
}

func NewPasswordProtector(user, passwordName string, algo HashAlgorithm) PasswordProtector {
    return PasswordProtector{user: user, passwordName: passwordName, hashAlgo: algo}
}

func (p *PasswordProtector) SetHashAlgorithm(algo HashAlgorithm) {
    p.hashAlgo = algo
}

func (p *PasswordProtector) Hash() {
    p.hashAlgo.Hash(*p)
}
```
## ¿Cómo se implementan algoritmos de hash intercambiables?
Dentro de este marco, cada algoritmo de hashing se implementa como una clase aislada. Utilizaremos ejemplos como SHA y MD5, en donde ambos algoritmos pueden ser intercambiables gracias a la estructura del patrón Strategy.
```go

type SHA struct{}

func (s SHA) Hash(p PasswordProtector) {
    fmt.Printf("Hashing usando SHA para el password: %s\n", p.passwordName)
}

type MD5 struct{}

func (m MD5) Hash(p PasswordProtector) {
    fmt.Printf("Hashing usando MD5 para el password: %s\n", p.passwordName)
}
```
En el ejemplo anterior, hemos creado dos algoritmos distintos (SHA y MD5) que implementan la interfaz HashAlgorithm. Esta estructura permite que los algoritmos se intercambien dinámicamente según la necesidad.

## ¿Cómo se utiliza el patrón Strategy en una función main?
La funcionalidad del patrón Strategy se pone a prueba dentro de la función main, donde los algoritmos son fácilmente intercambiables sin necesidad de modificar el código que los invoca directamente. Eso se logra mediante la capacidad del PasswordProtector de cambiar el algoritmo de hash en pleno vuelo.
```go

func main() {
    sha := SHA{}
    md5 := MD5{}

    passwordProtector := NewPasswordProtector("Nestor", "Gmail", sha)
    passwordProtector.Hash()

    passwordProtector.SetHashAlgorithm(md5)
    passwordProtector.Hash()
}
```

Al ejecutar el código, se observa cómo inicialmente se utiliza el algoritmo de SHA, y posteriormente se cambia a MD5. Esta capacidad de intercambio demuestra la flexibilidad y adaptabilidad que el patrón Strategy confiere al desarrollo del software. Implementar este patrón no solo resulta en un código más limpio y reutilizable, sino que también se adhiere a principios como el de responsabilidad única y la apertura al cambio sin modificar estructuras existentes. Así se cultiva un ciclo de desarrollo robusto y adaptable a futuros cambios o mejoras.

## ¿Cómo escanear puertos utilizando Go?
El lenguaje de programación Go se destaca por su potente librería estándar, razón por la cual es muy popular entre los desarrolladores. En esta guía, exploraremos el paquete net de Go para generar conexiones TCP a diferentes servidores y escanear los puertos disponibles. Esto te permitirá determinar cuáles puertos están abiertos y cuáles no, un conocimiento esencial si te dedicas a la ciberseguridad.

## ¿Qué pasos seguir en la codificación para el escaneo de puertos?
Para comenzar, debes crear una nueva carpeta llamada "net" y un archivo nuevo denominado port.go. Este será el esqueleto de tu programa. El paquete main será utilizado para este ejercicio, ya que es el punto de entrada del programa. A continuación, se detalla el proceso para crear el escaneo de puertos.
```go

package main

import (
    "fmt"
    "net"
)

func main() {
    for i := 0; i < 100; i++ {
        address := fmt.Sprintf("scanme.nmap.org:%d", i)
        conn, err := net.Dial("tcp", address)
        if err != nil {
            continue
        }
        conn.Close()
        fmt.Printf("Puerto %d está abierto\n", i)
    }
}
```
1. Crear un ciclo para escanear puertos: El ciclo se inicializa en cero y continúa hasta llegar al puerto 99.
2. Conexiones TCP: Se establece una conexión TCP a cada puerto posible. Si la conexión es exitosa, el puerto se considera abierto.
3. Manejo de errores: Si no se puede conectar a un puerto, se omite el error para continuar explorando los siguientes puertos.
4. Cierre de conexiones: Es vital cerrar la conexión una vez verificada la apertura del puerto, similar a cerrar archivos abiertos.

## ¿Qué consideraciones tener al escanear puertos de un servidor?
Cuando se escanean puertos, es crucial entender que esta técnica puede ser vista como un ciberataque si no se lleva a cabo de manera adecuada. Por lo tanto, te ofrecemos algunas recomendaciones clave:

- Utilizar sitios de prueba: Escanea sitios específicamente diseñados para pruebas, como scanme.nmap.org.
- Usar con fines didácticos: Esto debería hacerse en tu propio entorno o en servidores de prueba, nunca en servidores ajenos sin permiso.
- Educación continua: Considera inscribirte en cursos de seguridad informática para entender mejor las implicaciones éticas y legales de estas técnicas.

## ¿Cuál es el problema del escaneo de puertos de este ejemplo?
Una de las principales limitaciones del primer método de escaneo de puertos es su lentitud. En este caso, se limitaron a escanear 100 puertos, pero el internet abarca hasta 65,535 puertos posibles. Debido a que este proceso es secuencial, puede ser extremadamente ineficiente y demorado.

En futuras sesiones, se puede mejorar la eficiencia al implementar el escaneo de puertos concurrente utilizando 'Gorutinas'. Esto permitirá realizar múltiples escaneos simultáneamente, reduciendo el tiempo total.

Recuerda siempre que el uso ético y responsable del conocimiento es fundamental, así que aprovecha estas técnicas de manera segura y legal.

## ¿Qué es NetCAD y cómo se implementa?
Crear un proyecto utilizando NetCAD es una oportunidad magnífica para aprender sobre conexiones de red y la transmisión de datos. NetCAD es una herramienta que permite escribir y leer a través de conexiones TCP o UDP, mostrándolo en una consola. En este proyecto, NetCAD funcionará como cliente para un servidor de chat, permitiendo el intercambio de mensajes.

## ¿Cómo se inicia la implementación de NetCAD?
Primero, creamos un nuevo archivo llamado netcad.go dentro de la carpeta net. Definimos el paquete principal con package main y creamos la función main, que establecerá una conexión usando el protocolo TCP.
```go

package main

import (
    "flag"
    "fmt"
    "log"
    "net"
    "os"
    "os/exec"
)

func main() {
    // Definición de flags para host y puerto
    port := flag.Int("p", 3090, "el puerto")
    host := flag.String("h", "localhost", "el host")
    flag.Parse()

    // Creación de la conexión TCP
    connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
    if err != nil {
        log.Fatal(err)
    }
    defer connection.Close()

    // Canal de control para manejo de concurrencia
    done := make(chan struct{})

    go func() {
        io.Copy(os.Stdout, connection) 
        done <- struct{}{}
    }()

    // Llamada a la función copyContent
    copyContent(connection, os.Stdin)
    <-done
}

// Función para copiar contenido entre el lector y el escritor
func copyContent(dst net.Conn, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}
```
## ¿Cómo se configura y utiliza la función main?
En la función main, se definen dos parámetros importantes mediante flags: el host y el port.

- Host: Se representa con el flag -h y por defecto es localhost.
- Puerto: Se configura con -p y su valor por defecto es 3090.

Estos flags permiten que los parámetros sean personalizables durante la ejecución del programa. Una vez establecidos, se establece la conexión. En caso de error, el programa termina su ejecución mediante log.Fatal(). Si no hay error, el programa procede a crear un canal que permitirá manejar las operaciones concurrentes.

## ¿Cómo se gestionan las operaciones concurrentes y el control de flujo?
La función anónima dentro de main se encarga de leer todo lo que se recibe a través de la conexión y mostrarlo en la consola. Utiliza io.Copy, que es una función eficiente del paquete io en Go, para copiar datos de un Reader a un Writer. Aquí, la consola actúa como el Writer y la conexión como el Reader.

Además, se envía un struct vacío a un canal para señalar que la lectura ha terminado. Esto es crucial para saber cuándo desbloquear el programa.

## ¿Qué sigue después de implementar NetCAD?
Finalmente, debemos recordar que NetCAD aun no está completo sin un servidor al cual conectarse. En las próximas etapas del proyecto, nos enfocaremos en construir el servidor que complementará a NetCAD, permitiendo una interacción bidireccional con el cliente/NetCAD.

Implementar NetCAD es un paso importante para adquirir experiencia en programación concurrente y manejo de conexiones en Go. No obstante, entender cómo crear y gestionar correctamente tanto el cliente como el servidor es clave para el éxito del proyecto. ¡Sigue adelante, sigue aprendiendo y completa tu proyecto con éxito!


## ¿Cómo construir un servidor de chat en Go?
En esta clase, abordamos la creación de un servidor de chat utilizando Go. Si bien en la lección anterior creamos el cliente, ahora es momento de profundizar en el backend del proyecto para dar vida a nuestro sistema de mensajería. Vamos a explorar el código y entender cada pieza fundamental en este rompecabezas.

## ¿Qué es el tipo client y cómo se utiliza?
Primero, definimos un nuevo tipo llamado client, que es esencialmente un canal que transmite strings. Este tipo es clave, ya que es el medio por el cual enviaremos mensajes dentro del chat. Aquí están los pasos necesarios:

1. Definición inicial:

`type client chan<- string`

2. Variables de canal:
- incomeingClients: canal para los clientes que se conectan.
- livingClients: canal para los que abandonan el chat.
- messages: canal para los mensajes en sí, de tipo string.

## ¿Cómo se gestionan las conexiones del cliente?
La función handleConnection se encarga de manejar las conexiones de los clientes de manera individual. Cada cliente que se conecta es asignado a una instancia de esta función, asegurando un manejo efectivo y seguro de los recursos.

Pasos dentro de `handleConnection`:
1. Asignar un nombre único a cada cliente usando la dirección remota.

`clientName := con.RemoteAddr().String()`

2. Enviar un mensaje de bienvenida al cliente que se conecta y notificar al resto de los clientes.
```go
messages <- fmt.Sprintf("Bienvenido al servidor, %s", clientName)
messages <- fmt.Sprintf("Nuevo cliente ha llegado: %s", clientName)
```
3. Lectura de mensajes: Utilizando un escáner para leer continuamente mensajes desde la terminal:
```go
inputMessage := bufio.NewScanner(con)
for inputMessage.Scan() {
    messages <- fmt.Sprintf("%s: %s", clientName, inputMessage.Text())
}
```
## ¿Cómo funciona la escritura de mensajes?
La función messageWrite es responsable de escribir los mensajes a través de la conexión de cada cliente. Utiliza una goroutine para manejar la concurrencia y mantener el flujo de mensajes constante.
```go

func messageWrite(con net.Conn, message chan string) {
    for msg := range message {
        fmt.Fprintln(con, msg)
    }
}
```
## ¿Cómo manejar la desconexión del cliente?
Cuando un cliente decide desconectar, es importante informar a los demás y liberar recursos:

1. Se finaliza el ciclo de escaneo de mensajes.
2. Se reporta que un cliente ha abandonado el chat.
`messages <- fmt.Sprintf("%s ha dejado el chat.", clientName)`

## ¿Cuáles son las herramientas clave utilizadas?
Para llevar a cabo este proyecto, hemos aprovechado las siguientes herramientas de Go:

- Channels: Para la comunicación entre goroutines.
- Goroutines: Para la ejecución concurrente de las funciones.
- Packages estándar de Go: Como net para manejar las conexiones y fmt para el formateo de salida.

Con este conjunto de herramientas y técnicas, has aprendido a establecer la base de un servidor de chat en Go. Aunque este es solo el comienzo, invita a la exploración y manipulación de conexiones concurrentes y mensajes en un entorno de producción real. Si estás interesado en completar y ejecutar este proyecto, te animo a continuar aprendiendo y descubrir cómo optimizar y expandir este sistema de chat. ¡Buena suerte y sigue adelante!


## ¿Cómo manejar las conexiones de los clientes?
El manejo eficiente de múltiples conexiones de clientes es esencial para un servidor de chat robusto. La clave radica en mapear las conexiones con el estado de cada cliente, permitiendo saber quién está conectado y facilitar la transmisión de mensajes entre ellos.

- Creación del mapa de clientes: Define un mapa en Go que asocie cada cliente con un valor booleano.

`var clients = make(map[Client]bool)`
- Multiplexación de canales: Utiliza select para manejar las diferentes acciones de conexión:
1. Caso de un nuevo mensaje: Recorre cada cliente y envía el mensaje recibido.
2. Conexión de un nuevo cliente: Añade el cliente al mapa para asegurar su participación en futuros mensajes.
3. Desconexión de un cliente: Elimina al cliente del mapa para evitar mensajes innecesarios.

## ¿Cómo gestionar la lógica de transmisión de mensajes?
El componente core de un servidor de chat es la transmisión (broadcast) de mensajes entre los usuarios conectados. Implementar esta función garantiza que cada mensaje se distribuya a todos los clientes de manera efectiva.

- Función broadcast: Realiza una iteración indefinida a través del mapa de clientes, manejando mensajes nuevos, nuevas conexiones y desconexiones.
```go

func broadcast() {
  for {
    select {
    case msg := <-messages:
      for client := range clients {
        client <- msg
      }
    case newClient := <-incomingClients:
      clients[newClient] = true
    case leavingClient := <-leavingClients:
      delete(clients, leavingClient)
      close(leavingClient)
    }
  }
}
```
## ¿Cómo poner en marcha el servidor de chat?
Una vez resuelta la lógica interna, es crucial iniciar y ejecutar el servidor de chat. Este proceso implica la inicialización del servidor, la gestión de errores y la aceptación de conexiones entrantes.

- Función main: Configura el servidor para escuchar conexiones en un puerto específico.
```go

func main() {
  listener, err := net.Listen("tcp", net.JoinHostPort(host, port))
  if err != nil {
    log.Fatal("Error al iniciar el listener: ", err)
  }
  defer listener.Close()

  go broadcast()

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Println("Error al aceptar conexión: ", err)
      continue
    }
    go handleConnection(conn)
  }
}
```
## ¿Cómo interactuar con el chat?
Con el servidor en marcha, la interacción con el chat es el siguiente paso. Utiliza NetCAT para conectarte al chat como cliente, permitiendo enviar y recibir mensajes en tiempo real.

- Construcción y ejecución:
```bash
go build -o netcat NetCAD.go
go build -o chat Chat.go
./chat
```
- Conexión a través de NetCAT:
    - Ejecuta NetCAT en nuevas terminales para simular clientes.
    - Envía y recibe mensajes, y observa cómo el servidor maneja las conexiones.

### Reflexiones finales
Implementar un servidor de chat en Golang no solo es un excelente ejercicio de programación de redes, sino también una manera de apreciar el potencial y la robustez del lenguaje Go. Este proyecto demuestra cómo gestionar eficientemente las conexiones TCP y construir aplicaciones escalables.

Ahora que has aprendido a levantar un servidor de chat básico, ¡continúa explorando las posibilidades que Golang ofrece en el desarrollo de soluciones de red eficientes y potentes!












