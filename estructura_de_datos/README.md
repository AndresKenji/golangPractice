# Estructuras de datos lineales

Al igual que en la mayoria de lenguajes de programación GO permite crear agrupaciones de datos de un mismo tipo en una variable y acceder a ella segun el indice, estas agrupaciones son conocidas como lineales porque sus datos están ordenados segununa sola dimensión.

Go permite crear dos tipos de datos lineales, los **vectores (arrays)** y las **porciones (slice)**, que definen vistas sobre otros vectores y proporcionan mayor dinamismo y manejabilidad


1. Arrays:

Un array en Go es una colección de elementos del mismo tipo con una longitud fija. La longitud del array es parte de su tipo, lo que significa que no puede cambiar después de su creación.

```go
// Declaración y creación de un array de enteros con longitud 3
var miArray [3]int
miArray[0] = 1
miArray[1] = 2
miArray[2] = 3
```

2. Slices:

Un slice es una porción flexible de un array. A diferencia de los arrays, los slices son dinámicos y su tamaño puede cambiar.

```go
// Creación de un slice
miSlice := []int{1, 2, 3, 4, 5}

// Agregar un elemento al final del slice
miSlice = append(miSlice, 6)
```

3. Mapas:

Un mapa es una colección no ordenada de pares clave-valor, donde cada clave es única.

```go
// Creación de un mapa
miMapa := make(map[string]int)
miMapa["uno"] = 1
miMapa["dos"] = 2
```

4. Listas Doblemente Enlazadas:

Go no tiene una implementación estándar de listas doblemente enlazadas, pero puedes implementarlas tú mismo utilizando estructuras personalizadas.

```go
type Nodo struct {
    Valor       int
    Siguiente   *Nodo
    Anterior    *Nodo
}

// Creación de una lista doblemente enlazada
nodo1 := Nodo{Valor: 1}
nodo2 := Nodo{Valor: 2}
nodo1.Siguiente = &nodo2
nodo2.Anterior = &nodo1
```

5. Colas y Pilas:

Puedes implementar colas y pilas utilizando slices y aplicando las operaciones apropiadas.

```go
// Implementación de una cola (FIFO)
miCola := []int{1, 2, 3}
miCola = append(miCola, 4)       // Encolar
primerElemento := miCola[0]       // Desencolar

// Implementación de una pila (LIFO)
miPila := []int{1, 2, 3}
miPila = append(miPila, 4)        // Apilar
ultimoElemento := miPila[len(miPila)-1]  // Desapilar
```

Estas son solo algunas de las estructuras de datos lineales que se utilizan en Go. La elección de la estructura de datos dependerá de los requisitos específicos de la aplicación y de la eficiencia deseada para las operaciones que se realizen. Además de las estructuras de datos mencionadas, Go también proporciona paquetes para listas, conjuntos y otras estructuras de datos en la biblioteca estándar.