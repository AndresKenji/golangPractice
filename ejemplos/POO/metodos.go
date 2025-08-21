package main

import (
	"fmt"
	"math"
	"strings"
)

// ============= MÉTODOS Y RECEIVERS =============

// Punto 2D - para demostrar value vs pointer receivers
type Punto struct {
	X, Y float64
}

// Método con value receiver - no modifica el original
func (p Punto) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}

// Método con value receiver - no modifica el original
func (p Punto) Distancia(otro Punto) float64 {
	dx := p.X - otro.X
	dy := p.Y - otro.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// Método con value receiver - devuelve nuevo punto
func (p Punto) Trasladar(dx, dy float64) Punto {
	return Punto{X: p.X + dx, Y: p.Y + dy}
}

// Método con pointer receiver - modifica el original
func (p *Punto) Mover(dx, dy float64) {
	p.X += dx
	p.Y += dy
}

// Método con pointer receiver - modifica el original
func (p *Punto) Escalar(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Método con pointer receiver - establece nueva posición
func (p *Punto) EstablecerPosicion(x, y float64) {
	p.X = x
	p.Y = y
}

// Vector 3D para métodos más complejos
type Vector3D struct {
	X, Y, Z float64
}

// Métodos con value receiver
func (v Vector3D) String() string {
	return fmt.Sprintf("Vector3D(%.2f, %.2f, %.2f)", v.X, v.Y, v.Z)
}

func (v Vector3D) Magnitud() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector3D) Sumar(otro Vector3D) Vector3D {
	return Vector3D{
		X: v.X + otro.X,
		Y: v.Y + otro.Y,
		Z: v.Z + otro.Z,
	}
}

func (v Vector3D) Restar(otro Vector3D) Vector3D {
	return Vector3D{
		X: v.X - otro.X,
		Y: v.Y - otro.Y,
		Z: v.Z - otro.Z,
	}
}

func (v Vector3D) ProductoEscalar(otro Vector3D) float64 {
	return v.X*otro.X + v.Y*otro.Y + v.Z*otro.Z
}

func (v Vector3D) ProductoVectorial(otro Vector3D) Vector3D {
	return Vector3D{
		X: v.Y*otro.Z - v.Z*otro.Y,
		Y: v.Z*otro.X - v.X*otro.Z,
		Z: v.X*otro.Y - v.Y*otro.X,
	}
}

// Métodos con pointer receiver
func (v *Vector3D) Normalizar() {
	mag := v.Magnitud()
	if mag > 0 {
		v.X /= mag
		v.Y /= mag
		v.Z /= mag
	}
}

func (v *Vector3D) EscalarPor(factor float64) {
	v.X *= factor
	v.Y *= factor
	v.Z *= factor
}

// Contador para demostrar métodos que modifican estado
type Contador struct {
	valor int
	paso  int
}

func NuevoContador(valorInicial, paso int) *Contador {
	return &Contador{valor: valorInicial, paso: paso}
}

// Métodos con value receiver - solo lectura
func (c Contador) Valor() int {
	return c.valor
}

func (c Contador) Paso() int {
	return c.paso
}

func (c Contador) String() string {
	return fmt.Sprintf("Contador: %d (paso: %d)", c.valor, c.paso)
}

// Métodos con pointer receiver - modifican estado
func (c *Contador) Incrementar() {
	c.valor += c.paso
}

func (c *Contador) Decrementar() {
	c.valor -= c.paso
}

func (c *Contador) Reiniciar() {
	c.valor = 0
}

func (c *Contador) EstablecerValor(nuevoValor int) {
	c.valor = nuevoValor
}

func (c *Contador) EstablecerPaso(nuevoPaso int) {
	c.paso = nuevoPaso
}

// Lista enlazada para métodos recursivos
type Nodo struct {
	Valor     int
	Siguiente *Nodo
}

type ListaEnlazada struct {
	cabeza *Nodo
	tamaño int
}

func NuevaListaEnlazada() *ListaEnlazada {
	return &ListaEnlazada{cabeza: nil, tamaño: 0}
}

// Métodos con value receiver
func (le ListaEnlazada) Tamaño() int {
	return le.tamaño
}

func (le ListaEnlazada) EstaVacia() bool {
	return le.cabeza == nil
}

func (le ListaEnlazada) String() string {
	if le.cabeza == nil {
		return "Lista: []"
	}

	var elementos []string
	actual := le.cabeza
	for actual != nil {
		elementos = append(elementos, fmt.Sprintf("%d", actual.Valor))
		actual = actual.Siguiente
	}
	return fmt.Sprintf("Lista: [%s]", strings.Join(elementos, ", "))
}

// Métodos con pointer receiver
func (le *ListaEnlazada) Agregar(valor int) {
	nuevoNodo := &Nodo{Valor: valor, Siguiente: le.cabeza}
	le.cabeza = nuevoNodo
	le.tamaño++
}

func (le *ListaEnlazada) AgregarAlFinal(valor int) {
	nuevoNodo := &Nodo{Valor: valor, Siguiente: nil}

	if le.cabeza == nil {
		le.cabeza = nuevoNodo
	} else {
		actual := le.cabeza
		for actual.Siguiente != nil {
			actual = actual.Siguiente
		}
		actual.Siguiente = nuevoNodo
	}
	le.tamaño++
}

func (le *ListaEnlazada) Eliminar(valor int) bool {
	if le.cabeza == nil {
		return false
	}

	// Si el primer nodo es el que queremos eliminar
	if le.cabeza.Valor == valor {
		le.cabeza = le.cabeza.Siguiente
		le.tamaño--
		return true
	}

	// Buscar en el resto de la lista
	actual := le.cabeza
	for actual.Siguiente != nil {
		if actual.Siguiente.Valor == valor {
			actual.Siguiente = actual.Siguiente.Siguiente
			le.tamaño--
			return true
		}
		actual = actual.Siguiente
	}

	return false
}

func (le *ListaEnlazada) Limpiar() {
	le.cabeza = nil
	le.tamaño = 0
}

// Métodos que demuestran diferencias entre value y pointer receivers
func (le ListaEnlazada) Contiene(valor int) bool {
	actual := le.cabeza
	for actual != nil {
		if actual.Valor == valor {
			return true
		}
		actual = actual.Siguiente
	}
	return false
}

// Stack basado en slice
type Stack struct {
	elementos []int
}

func NuevaStack() *Stack {
	return &Stack{elementos: make([]int, 0)}
}

// Métodos con value receiver
func (s Stack) Tamaño() int {
	return len(s.elementos)
}

func (s Stack) EstaVacia() bool {
	return len(s.elementos) == 0
}

func (s Stack) Tope() (int, bool) {
	if s.EstaVacia() {
		return 0, false
	}
	return s.elementos[len(s.elementos)-1], true
}

func (s Stack) String() string {
	if s.EstaVacia() {
		return "Stack: []"
	}

	var elementos []string
	for i := len(s.elementos) - 1; i >= 0; i-- {
		elementos = append(elementos, fmt.Sprintf("%d", s.elementos[i]))
	}
	return fmt.Sprintf("Stack: [%s] (tope -> base)", strings.Join(elementos, ", "))
}

// Métodos con pointer receiver
func (s *Stack) Push(valor int) {
	s.elementos = append(s.elementos, valor)
}

func (s *Stack) Pop() (int, bool) {
	if s.EstaVacia() {
		return 0, false
	}

	indice := len(s.elementos) - 1
	valor := s.elementos[indice]
	s.elementos = s.elementos[:indice]
	return valor, true
}

func (s *Stack) Limpiar() {
	s.elementos = s.elementos[:0]
}

func ejemploMetodos() {
	fmt.Println("  === PUNTOS Y VECTORES ===")

	// Crear puntos
	punto1 := Punto{X: 3, Y: 4}
	punto2 := Punto{X: 6, Y: 8}

	fmt.Printf("    Punto 1: %s\n", punto1)
	fmt.Printf("    Punto 2: %s\n", punto2)
	fmt.Printf("    Distancia: %.2f\n", punto1.Distancia(punto2))

	// Value receiver - no modifica el original
	punto3 := punto1.Trasladar(1, 1)
	fmt.Printf("    Punto 1 trasladado (nuevo): %s\n", punto3)
	fmt.Printf("    Punto 1 original: %s\n", punto1)

	// Pointer receiver - modifica el original
	fmt.Println("\n  Modificando punto con pointer receiver:")
	punto1.Mover(2, 3)
	fmt.Printf("    Punto 1 después de mover: %s\n", punto1)

	punto1.Escalar(2)
	fmt.Printf("    Punto 1 después de escalar: %s\n", punto1)

	// Vectores 3D
	fmt.Println("\n  === VECTORES 3D ===")
	vector1 := Vector3D{X: 1, Y: 2, Z: 3}
	vector2 := Vector3D{X: 4, Y: 5, Z: 6}

	fmt.Printf("    Vector 1: %s (magnitud: %.2f)\n", vector1, vector1.Magnitud())
	fmt.Printf("    Vector 2: %s (magnitud: %.2f)\n", vector2, vector2.Magnitud())

	suma := vector1.Sumar(vector2)
	fmt.Printf("    Suma: %s\n", suma)

	productoEscalar := vector1.ProductoEscalar(vector2)
	fmt.Printf("    Producto escalar: %.2f\n", productoEscalar)

	productoVectorial := vector1.ProductoVectorial(vector2)
	fmt.Printf("    Producto vectorial: %s\n", productoVectorial)

	// Normalizar vector (pointer receiver)
	fmt.Printf("    Vector 1 antes de normalizar: %s\n", vector1)
	vector1.Normalizar()
	fmt.Printf("    Vector 1 normalizado: %s (magnitud: %.2f)\n", vector1, vector1.Magnitud())

	fmt.Println("\n  === CONTADORES ===")

	// Crear contador
	contador := NuevoContador(10, 3)
	fmt.Printf("    %s\n", contador)

	// Operaciones que modifican estado
	contador.Incrementar()
	fmt.Printf("    Después de incrementar: %s\n", contador)

	contador.Incrementar()
	contador.Incrementar()
	fmt.Printf("    Después de 2 incrementos más: %s\n", contador)

	contador.Decrementar()
	fmt.Printf("    Después de decrementar: %s\n", contador)

	contador.EstablecerPaso(5)
	contador.Incrementar()
	fmt.Printf("    Con paso 5, después de incrementar: %s\n", contador)

	contador.Reiniciar()
	fmt.Printf("    Después de reiniciar: %s\n", contador)

	fmt.Println("\n  === LISTA ENLAZADA ===")

	// Crear lista enlazada
	lista := NuevaListaEnlazada()
	fmt.Printf("    Lista inicial: %s (vacía: %t)\n", lista, lista.EstaVacia())

	// Agregar elementos
	lista.Agregar(10)
	lista.Agregar(20)
	lista.Agregar(30)
	fmt.Printf("    Después de agregar al inicio: %s\n", lista)

	lista.AgregarAlFinal(5)
	lista.AgregarAlFinal(1)
	fmt.Printf("    Después de agregar al final: %s\n", lista)

	// Buscar elementos
	fmt.Printf("    ¿Contiene 20?: %t\n", lista.Contiene(20))
	fmt.Printf("    ¿Contiene 100?: %t\n", lista.Contiene(100))

	// Eliminar elementos
	eliminado := lista.Eliminar(20)
	fmt.Printf("    Después de eliminar 20 (éxito: %t): %s\n", eliminado, lista)

	eliminado = lista.Eliminar(999)
	fmt.Printf("    Intentar eliminar 999 (éxito: %t): %s\n", eliminado, lista)

	fmt.Printf("    Tamaño actual: %d\n", lista.Tamaño())

	fmt.Println("\n  === STACK ===")

	// Crear stack
	stack := NuevaStack()
	fmt.Printf("    Stack inicial: %s (vacía: %t)\n", stack, stack.EstaVacia())

	// Push elementos
	elementos := []int{10, 20, 30, 40, 50}
	for _, elem := range elementos {
		stack.Push(elem)
		fmt.Printf("    Push %d: %s\n", elem, stack)
	}

	// Verificar tope
	if tope, ok := stack.Tope(); ok {
		fmt.Printf("    Tope actual: %d\n", tope)
	}

	// Pop elementos
	fmt.Println("    Operaciones Pop:")
	for !stack.EstaVacia() {
		if valor, ok := stack.Pop(); ok {
			fmt.Printf("      Pop: %d -> %s\n", valor, stack)
		}
	}

	// Intentar pop en stack vacía
	if valor, ok := stack.Pop(); !ok {
		fmt.Println("    Pop en stack vacía: operación fallida (correcto)")
	} else {
		fmt.Printf("    Pop inesperado: %d\n", valor)
	}

	fmt.Printf("    Stack final: %s (tamaño: %d)\n", stack, stack.Tamaño())
}
