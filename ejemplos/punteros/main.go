package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== EJEMPLOS DE PUNTEROS EN GO ===\n")

	// 1. Declaración básica de punteros
	fmt.Println("1. Declaración básica de punteros:")
	var numero int = 42
	var puntero *int = &numero // & obtiene la dirección de memoria

	fmt.Printf("   Valor de numero: %d\n", numero)
	fmt.Printf("   Dirección de numero: %p\n", &numero)
	fmt.Printf("   Valor del puntero: %p\n", puntero)
	fmt.Printf("   Valor al que apunta el puntero: %d\n", *puntero) // * desreferencia el puntero
	fmt.Println()

	// 2. Modificación a través de punteros
	fmt.Println("2. Modificación a través de punteros:")
	fmt.Printf("   Valor original: %d\n", numero)
	*puntero = 100 // Modificamos el valor a través del puntero
	fmt.Printf("   Valor después de modificar por puntero: %d\n", numero)
	fmt.Println()

	// 3. Punteros nulos (nil)
	fmt.Println("3. Punteros nulos (nil):")
	var punteroNulo *int
	fmt.Printf("   Puntero nulo: %v\n", punteroNulo)
	fmt.Printf("   ¿Es nil? %t\n", punteroNulo == nil)
	
	// Verificación antes de desreferenciar
	if punteroNulo != nil {
		fmt.Printf("   Valor: %d\n", *punteroNulo)
	} else {
		fmt.Println("   No se puede desreferenciar un puntero nil")
	}
	fmt.Println()

	// 4. Creación de punteros con new()
	fmt.Println("4. Creación de punteros con new():")
	punteroNew := new(int) // new() asigna memoria y retorna un puntero
	fmt.Printf("   Dirección: %p\n", punteroNew)
	fmt.Printf("   Valor inicial: %d\n", *punteroNew) // new() inicializa con valor cero
	*punteroNew = 77
	fmt.Printf("   Valor asignado: %d\n", *punteroNew)
	fmt.Println()

	// 5. Punteros a diferentes tipos
	fmt.Println("5. Punteros a diferentes tipos:")
	texto := "Hola Go"
	flotante := 3.14159
	booleano := true

	pTexto := &texto
	pFlotante := &flotante
	pBooleano := &booleano

	fmt.Printf("   String: %s -> Puntero: %p -> Valor: %s\n", texto, pTexto, *pTexto)
	fmt.Printf("   Float: %.2f -> Puntero: %p -> Valor: %.2f\n", flotante, pFlotante, *pFlotante)
	fmt.Printf("   Bool: %t -> Puntero: %p -> Valor: %t\n", booleano, pBooleano, *pBooleano)
	fmt.Println()

	// 6. Punteros en funciones - Paso por referencia
	fmt.Println("6. Punteros en funciones - Paso por referencia:")
	valor := 10
	fmt.Printf("   Valor antes: %d\n", valor)
	duplicarPorValor(valor)
	fmt.Printf("   Valor después de duplicarPorValor: %d\n", valor)
	duplicarPorReferencia(&valor)
	fmt.Printf("   Valor después de duplicarPorReferencia: %d\n", valor)
	fmt.Println()

	// 7. Intercambio de valores usando punteros
	fmt.Println("7. Intercambio de valores usando punteros:")
	a, b := 5, 10
	fmt.Printf("   Antes del intercambio: a=%d, b=%d\n", a, b)
	intercambiar(&a, &b)
	fmt.Printf("   Después del intercambio: a=%d, b=%d\n", a, b)
	fmt.Println()

	// 8. Punteros a estructuras
	fmt.Println("8. Punteros a estructuras:")
	persona := Persona{Nombre: "Juan", Edad: 30}
	punteroPersona := &persona

	fmt.Printf("   Acceso directo: %s, %d años\n", persona.Nombre, persona.Edad)
	fmt.Printf("   Acceso por puntero: %s, %d años\n", punteroPersona.Nombre, punteroPersona.Edad)
	fmt.Printf("   Acceso por puntero (explícito): %s, %d años\n", (*punteroPersona).Nombre, (*punteroPersona).Edad)

	// Modificación a través del puntero
	punteroPersona.Edad = 31
	fmt.Printf("   Después de modificar por puntero: %s, %d años\n", persona.Nombre, persona.Edad)
	fmt.Println()

	// 9. Arreglos y punteros
	fmt.Println("9. Arreglos y punteros:")
	arreglo := [5]int{1, 2, 3, 4, 5}
	punteroArreglo := &arreglo

	fmt.Printf("   Arreglo original: %v\n", arreglo)
	fmt.Printf("   Primer elemento por puntero: %d\n", punteroArreglo[0])
	
	// Modificar a través del puntero
	punteroArreglo[0] = 100
	fmt.Printf("   Arreglo después de modificar: %v\n", arreglo)
	fmt.Println()

	// 10. Slices y punteros
	fmt.Println("10. Slices y punteros:")
	slice := []int{10, 20, 30}
	fmt.Printf("   Slice original: %v\n", slice)
	
	modificarSlice(slice) // Los slices se pasan por referencia por defecto
	fmt.Printf("   Slice después de modificar: %v\n", slice)
	
	// Puntero al slice completo
	punteroSlice := &slice
	*punteroSlice = append(*punteroSlice, 40, 50)
	fmt.Printf("   Slice después de append por puntero: %v\n", slice)
	fmt.Println()

	// 11. Punteros a punteros
	fmt.Println("11. Punteros a punteros:")
	numero2 := 123
	puntero1 := &numero2
	puntero2 := &puntero1 // Puntero a puntero

	fmt.Printf("   Valor original: %d\n", numero2)
	fmt.Printf("   Valor por puntero1: %d\n", *puntero1)
	fmt.Printf("   Valor por puntero2: %d\n", **puntero2)
	fmt.Printf("   Dirección de numero2: %p\n", &numero2)
	fmt.Printf("   Valor de puntero1: %p\n", puntero1)
	fmt.Printf("   Dirección de puntero1: %p\n", &puntero1)
	fmt.Printf("   Valor de puntero2: %p\n", puntero2)
	fmt.Println()

	// 12. Comparación de punteros
	fmt.Println("12. Comparación de punteros:")
	x := 42
	y := 42
	px1 := &x
	px2 := &x
	py := &y

	fmt.Printf("   px1 == px2 (mismo objeto): %t\n", px1 == px2)
	fmt.Printf("   px1 == py (diferentes objetos): %t\n", px1 == py)
	fmt.Printf("   *px1 == *py (mismos valores): %t\n", *px1 == *py)
	fmt.Println()

	// 13. Funciones que retornan punteros
	fmt.Println("13. Funciones que retornan punteros:")
	pEntero := crearEntero(999)
	fmt.Printf("   Valor creado: %d\n", *pEntero)
	
	pPersona := crearPersona("Ana", 25)
	fmt.Printf("   Persona creada: %s, %d años\n", pPersona.Nombre, pPersona.Edad)
	fmt.Println()

	// 14. Métodos con receptores por puntero
	fmt.Println("14. Métodos con receptores por puntero:")
	contador := Contador{valor: 0}
	fmt.Printf("   Valor inicial: %d\n", contador.valor)
	
	contador.Incrementar() // Modifica el receptor
	fmt.Printf("   Después de incrementar: %d\n", contador.valor)
	
	valorSinModificar := contador.ObtenerValor() // No modifica el receptor
	fmt.Printf("   Valor obtenido: %d\n", valorSinModificar)
	fmt.Println()

	// 15. Aritmética de punteros (limitada en Go)
	fmt.Println("15. Información de punteros con unsafe:")
	arr := [3]int{100, 200, 300}
	ptr := &arr[0]
	
	fmt.Printf("   Dirección del primer elemento: %p\n", ptr)
	fmt.Printf("   Tamaño de int: %d bytes\n", unsafe.Sizeof(int(0)))
	
	// En Go no hay aritmética de punteros directa como en C/C++
	// Se debe usar unsafe para operaciones avanzadas (no recomendado)
	fmt.Printf("   Valor en la dirección: %d\n", *ptr)
	fmt.Println()

	// 16. Ejemplo práctico: Lista enlazada simple
	fmt.Println("16. Ejemplo práctico: Lista enlazada simple:")
	lista := &Nodo{Dato: 1, Siguiente: nil}
	lista.Siguiente = &Nodo{Dato: 2, Siguiente: nil}
	lista.Siguiente.Siguiente = &Nodo{Dato: 3, Siguiente: nil}
	
	fmt.Print("   Lista: ")
	imprimirLista(lista)
	fmt.Println()

	// 17. Punteros en maps
	fmt.Println("17. Punteros en maps:")
	mapa := make(map[string]*int)
	num1, num2, num3 := 10, 20, 30
	
	mapa["uno"] = &num1
	mapa["dos"] = &num2
	mapa["tres"] = &num3
	
	for clave, puntero := range mapa {
		fmt.Printf("   %s: %d (dirección: %p)\n", clave, *puntero, puntero)
	}
	fmt.Println()

	fmt.Println("=== FIN DE EJEMPLOS DE PUNTEROS ===")
}

// Estructuras para los ejemplos
type Persona struct {
	Nombre string
	Edad   int
}

type Contador struct {
	valor int
}

type Nodo struct {
	Dato      int
	Siguiente *Nodo
}

// Funciones para demostrar paso por valor vs paso por referencia
func duplicarPorValor(n int) {
	n = n * 2 // Solo modifica la copia local
}

func duplicarPorReferencia(n *int) {
	*n = *n * 2 // Modifica el valor original
}

// Función para intercambiar valores
func intercambiar(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// Función para modificar slice
func modificarSlice(s []int) {
	if len(s) > 0 {
		s[0] = 999 // Modifica el slice original
	}
}

// Funciones que retornan punteros
func crearEntero(valor int) *int {
	numero := valor // Variable local
	return &numero  // Go permite retornar punteros a variables locales
}

func crearPersona(nombre string, edad int) *Persona {
	return &Persona{
		Nombre: nombre,
		Edad:   edad,
	}
}

// Métodos con receptor por puntero
func (c *Contador) Incrementar() {
	c.valor++ // Modifica el receptor original
}

// Método con receptor por valor
func (c Contador) ObtenerValor() int {
	return c.valor // No modifica el receptor
}

// Función para imprimir lista enlazada
func imprimirLista(nodo *Nodo) {
	for nodo != nil {
		fmt.Printf("%d", nodo.Dato)
		if nodo.Siguiente != nil {
			fmt.Print(" -> ")
		}
		nodo = nodo.Siguiente
	}
	fmt.Println()
}
