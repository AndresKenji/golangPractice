package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("=== EJEMPLOS DE FUNCIONES EN GO ===")
	fmt.Println()

	// 1. Funciones básicas
	fmt.Println("1. FUNCIONES BÁSICAS:")
	ejemplosFuncionesBasicas()
	fmt.Println()

	// 2. Funciones con parámetros
	fmt.Println("2. FUNCIONES CON PARÁMETROS:")
	ejemplosParametros()
	fmt.Println()

	// 3. Funciones con retorno simple
	fmt.Println("3. FUNCIONES CON RETORNO SIMPLE:")
	ejemplosRetornoSimple()
	fmt.Println()

	// 4. Funciones con retorno múltiple
	fmt.Println("4. FUNCIONES CON RETORNO MÚLTIPLE:")
	ejemplosRetornoMultiple()
	fmt.Println()

	// 5. Funciones anónimas
	fmt.Println("5. FUNCIONES ANÓNIMAS:")
	ejemplosFuncionesAnonimas()
	fmt.Println()

	// 6. Closures
	fmt.Println("6. CLOSURES:")
	ejemplosClosures()
	fmt.Println()

	// 7. Funciones como first-class citizens
	fmt.Println("7. FUNCIONES COMO FIRST-CLASS CITIZENS:")
	ejemplosFirstClass()
	fmt.Println()

	// 8. Funciones variádicas
	fmt.Println("8. FUNCIONES VARIÁDICAS:")
	ejemplosVariadicas()
	fmt.Println()

	// 9. Defer, panic y recover
	fmt.Println("9. DEFER, PANIC Y RECOVER:")
	ejemplosDeferPanicRecover()
	fmt.Println()

	// 10. Métodos y receivers
	fmt.Println("10. MÉTODOS Y RECEIVERS:")
	ejemplosMetodos()
	fmt.Println()

	// 11. Ejemplos de recursividad
	fmt.Println("11. EJEMPLOS DE RECURSIVIDAD:")
	runRecursiveExamples()
	fmt.Println()

	// 12. Ejemplos de goroutines y channels
	fmt.Println("12. EJEMPLOS DE GOROUTINES Y CHANNELS:")
	runGoroutineExamples()
	fmt.Println()

	fmt.Println("=== FIN DE LOS EJEMPLOS BÁSICOS DE FUNCIONES ===")
}

// ============= FUNCIONES BÁSICAS =============

// Función sin parámetros ni retorno
func saludar() {
	fmt.Println("  ¡Hola desde una función básica!")
}

// Función que solo ejecuta código
func mostrarMensaje() {
	fmt.Println("  Esta función no recibe parámetros ni retorna valores")
}

func ejemplosFuncionesBasicas() {
	saludar()
	mostrarMensaje()

	// Llamar función múltiples veces
	for i := 0; i < 3; i++ {
		fmt.Printf("  Llamada #%d: ", i+1)
		saludar()
	}
}

// ============= FUNCIONES CON PARÁMETROS =============

// Función con un parámetro
func saludarPersona(nombre string) {
	fmt.Printf("  ¡Hola, %s!\n", nombre)
}

// Función con múltiples parámetros
func sumar(a, b int) {
	resultado := a + b
	fmt.Printf("  %d + %d = %d\n", a, b, resultado)
}

// Función con parámetros de diferentes tipos
func presentarPersona(nombre string, edad int, activo bool) {
	estado := "inactivo"
	if activo {
		estado = "activo"
	}
	fmt.Printf("  Persona: %s, %d años, estado: %s\n", nombre, edad, estado)
}

// Función con parámetros del mismo tipo (sintaxis abreviada)
func calcularRectangulo(ancho, alto float64) {
	area := ancho * alto
	perimetro := 2 * (ancho + alto)
	fmt.Printf("  Rectángulo %.1f x %.1f: Área = %.2f, Perímetro = %.2f\n",
		ancho, alto, area, perimetro)
}

func ejemplosParametros() {
	saludarPersona("Ana")
	saludarPersona("Bruno")

	sumar(5, 3)
	sumar(10, 20)

	presentarPersona("Carmen", 25, true)
	presentarPersona("David", 30, false)

	calcularRectangulo(5.0, 3.0)
	calcularRectangulo(7.2, 4.8)
}

// ============= FUNCIONES CON RETORNO SIMPLE =============

// Función que retorna un entero
func multiplicar(a, b int) int {
	return a * b
}

// Función que retorna un string
func obtenerSaludo(nombre string) string {
	return fmt.Sprintf("Hola, %s! ¿Cómo estás?", nombre)
}

// Función que retorna un bool
func esPareja(numero int) bool {
	return numero%2 == 0
}

// Función que retorna un float64
func calcularCircunferencia(radio float64) float64 {
	return 2 * math.Pi * radio
}

// Función con lógica condicional
func clasificarEdad(edad int) string {
	if edad < 13 {
		return "niño"
	} else if edad < 20 {
		return "adolescente"
	} else if edad < 60 {
		return "adulto"
	} else {
		return "adulto mayor"
	}
}

func ejemplosRetornoSimple() {
	resultado := multiplicar(6, 7)
	fmt.Printf("  6 × 7 = %d\n", resultado)

	saludo := obtenerSaludo("Elena")
	fmt.Printf("  %s\n", saludo)

	numero := 8
	if esPareja(numero) {
		fmt.Printf("  %d es par\n", numero)
	} else {
		fmt.Printf("  %d es impar\n", numero)
	}

	radio := 5.0
	circunferencia := calcularCircunferencia(radio)
	fmt.Printf("  Circunferencia de radio %.1f = %.2f\n", radio, circunferencia)

	edades := []int{10, 16, 25, 65}
	for _, edad := range edades {
		categoria := clasificarEdad(edad)
		fmt.Printf("  %d años: %s\n", edad, categoria)
	}
}

// ============= FUNCIONES CON RETORNO MÚLTIPLE =============

// Función que retorna dos valores
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return a / b, nil
}

// Función con múltiples retornos nombrados
func analizarTexto(texto string) (palabras int, caracteres int, vocales int) {
	palabras = len(strings.Fields(texto))
	caracteres = len(texto)

	vocalesSet := "aeiouAEIOU"
	for _, char := range texto {
		if strings.ContainsRune(vocalesSet, char) {
			vocales++
		}
	}
	return // Retorno implícito de variables nombradas
}

// Función que retorna valor y estado de éxito
func buscarEnSlice(slice []int, target int) (int, bool) {
	for i, v := range slice {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

// Función que retorna múltiples cálculos
func estadisticasBasicas(numeros []float64) (min, max, promedio float64, err error) {
	if len(numeros) == 0 {
		return 0, 0, 0, fmt.Errorf("slice vacío")
	}

	min = numeros[0]
	max = numeros[0]
	suma := 0.0

	for _, num := range numeros {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		suma += num
	}

	promedio = suma / float64(len(numeros))
	return min, max, promedio, nil
}

func ejemplosRetornoMultiple() {
	// Ejemplo de división con manejo de errores
	resultado, err := dividir(10, 3)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  10 ÷ 3 = %.2f\n", resultado)
	}

	// Intentar división por cero
	_, err = dividir(10, 0)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	}

	// Análisis de texto
	texto := "Hola mundo, esto es una prueba"
	palabras, chars, vocales := analizarTexto(texto)
	fmt.Printf("  Texto: \"%s\"\n", texto)
	fmt.Printf("  Palabras: %d, Caracteres: %d, Vocales: %d\n", palabras, chars, vocales)

	// Búsqueda en slice
	numeros := []int{10, 20, 30, 40, 50}
	indice, encontrado := buscarEnSlice(numeros, 30)
	if encontrado {
		fmt.Printf("  Número 30 encontrado en índice %d\n", indice)
	} else {
		fmt.Printf("  Número 30 no encontrado\n")
	}

	// Estadísticas
	datos := []float64{1.5, 2.8, 3.1, 4.7, 2.2, 1.9, 3.8}
	min, max, promedio, err := estadisticasBasicas(datos)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  Estadísticas: Min=%.1f, Max=%.1f, Promedio=%.2f\n", min, max, promedio)
	}
}

// ============= FUNCIONES ANÓNIMAS =============

func ejemplosFuncionesAnonimas() {
	// Función anónima simple
	func() {
		fmt.Println("  Esta es una función anónima")
	}()

	// Función anónima con parámetros
	func(nombre string) {
		fmt.Printf("  Hola %s desde función anónima\n", nombre)
	}("Fernando")

	// Función anónima con retorno
	resultado := func(a, b int) int {
		return a*a + b*b
	}(3, 4)
	fmt.Printf("  3² + 4² = %d\n", resultado)

	// Asignar función anónima a variable
	cuadrado := func(x float64) float64 {
		return x * x
	}

	fmt.Printf("  Cuadrado de 5: %.1f\n", cuadrado(5))
	fmt.Printf("  Cuadrado de 7.5: %.2f\n", cuadrado(7.5))

	// Función anónima en slice
	operaciones := []func(int) int{
		func(x int) int { return x * 2 },     // Doble
		func(x int) int { return x * x },     // Cuadrado
		func(x int) int { return x * x * x }, // Cubo
	}

	numero := 3
	fmt.Printf("  Operaciones con %d:\n", numero)
	nombres := []string{"Doble", "Cuadrado", "Cubo"}
	for i, op := range operaciones {
		resultado := op(numero)
		fmt.Printf("    %s: %d\n", nombres[i], resultado)
	}

	// Función anónima como argumento
	aplicarOperacion := func(nums []int, op func(int) int) []int {
		result := make([]int, len(nums))
		for i, num := range nums {
			result[i] = op(num)
		}
		return result
	}

	numeros := []int{1, 2, 3, 4, 5}
	cuadrados := aplicarOperacion(numeros, func(x int) int { return x * x })
	fmt.Printf("  Números: %v\n", numeros)
	fmt.Printf("  Cuadrados: %v\n", cuadrados)
}

// ============= CLOSURES =============

// Función que retorna una función (closure)
func contador() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Closure con parámetros
func multiplicadorPor(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Closure que modifica variable externa
func acumulador(inicial int) func(int) int {
	suma := inicial
	return func(valor int) int {
		suma += valor
		return suma
	}
}

// Generador de funciones de validación
func validadorRango(min, max int) func(int) bool {
	return func(valor int) bool {
		return valor >= min && valor <= max
	}
}

func ejemplosClosures() {
	// Contador simple
	miContador := contador()
	fmt.Printf("  Contador: %d\n", miContador())
	fmt.Printf("  Contador: %d\n", miContador())
	fmt.Printf("  Contador: %d\n", miContador())

	// Múltiples contadores independientes
	contador1 := contador()
	contador2 := contador()
	fmt.Printf("  Contador1: %d, Contador2: %d\n", contador1(), contador2())
	fmt.Printf("  Contador1: %d, Contador2: %d\n", contador1(), contador2())

	// Multiplicadores
	doblar := multiplicadorPor(2)
	triplicar := multiplicadorPor(3)

	numero := 7
	fmt.Printf("  %d × 2 = %d\n", numero, doblar(numero))
	fmt.Printf("  %d × 3 = %d\n", numero, triplicar(numero))

	// Acumulador
	miAcumulador := acumulador(10)
	fmt.Printf("  Acumulador inicial: 10\n")
	fmt.Printf("  Agregar 5: %d\n", miAcumulador(5))   // 15
	fmt.Printf("  Agregar 3: %d\n", miAcumulador(3))   // 18
	fmt.Printf("  Agregar -2: %d\n", miAcumulador(-2)) // 16

	// Validadores
	validarEdad := validadorRango(18, 65)
	validarNota := validadorRango(0, 10)

	edades := []int{15, 25, 70, 35}
	fmt.Println("  Validación de edades (18-65):")
	for _, edad := range edades {
		valida := validarEdad(edad)
		fmt.Printf("    Edad %d: %t\n", edad, valida)
	}

	notas := []int{8, 15, -1, 7}
	fmt.Println("  Validación de notas (0-10):")
	for _, nota := range notas {
		valida := validarNota(nota)
		fmt.Printf("    Nota %d: %t\n", nota, valida)
	}
}

// ============= FUNCIONES COMO FIRST-CLASS CITIZENS =============

// Tipo de función personalizado
type Operacion func(int, int) int

// Funciones que implementan el tipo
func suma(a, b int) int     { return a + b }
func resta(a, b int) int    { return a - b }
func producto(a, b int) int { return a * b }

// Función que recibe otra función como parámetro
func aplicarOperacion(a, b int, op Operacion) int {
	return op(a, b)
}

// Función que retorna diferentes funciones según condición
func obtenerOperacion(simbolo string) Operacion {
	switch simbolo {
	case "+":
		return suma
	case "-":
		return resta
	case "*":
		return producto
	default:
		return func(a, b int) int { return 0 }
	}
}

// Map de funciones
var operaciones = map[string]Operacion{
	"sumar":       suma,
	"restar":      resta,
	"multiplicar": producto,
}

func ejemplosFirstClass() {
	a, b := 8, 3

	// Usar funciones como argumentos
	fmt.Printf("  %d + %d = %d\n", a, b, aplicarOperacion(a, b, suma))
	fmt.Printf("  %d - %d = %d\n", a, b, aplicarOperacion(a, b, resta))
	fmt.Printf("  %d × %d = %d\n", a, b, aplicarOperacion(a, b, producto))

	// Obtener función dinámicamente
	simbolos := []string{"+", "-", "*", "/"}
	for _, simbolo := range simbolos {
		op := obtenerOperacion(simbolo)
		resultado := op(a, b)
		fmt.Printf("  %d %s %d = %d\n", a, simbolo, b, resultado)
	}

	// Usar map de funciones
	fmt.Println("  Usando map de funciones:")
	for nombre, op := range operaciones {
		resultado := op(a, b)
		fmt.Printf("    %s(%d, %d) = %d\n", nombre, a, b, resultado)
	}

	// Slice de funciones
	funciones := []Operacion{suma, resta, producto}
	nombres := []string{"suma", "resta", "producto"}

	fmt.Println("  Usando slice de funciones:")
	for i, fn := range funciones {
		resultado := fn(a, b)
		fmt.Printf("    %s: %d\n", nombres[i], resultado)
	}
}

// ============= FUNCIONES VARIÁDICAS =============

// Función variádica básica
func sumarTodos(numeros ...int) int {
	suma := 0
	for _, num := range numeros {
		suma += num
	}
	return suma
}

// Función variádica con parámetros fijos
func saludarGrupo(saludo string, nombres ...string) {
	for _, nombre := range nombres {
		fmt.Printf("  %s, %s!\n", saludo, nombre)
	}
}

// Función variádica que retorna múltiples valores
func estadisticas(numeros ...float64) (float64, float64, float64) {
	if len(numeros) == 0 {
		return 0, 0, 0
	}

	suma := 0.0
	min := numeros[0]
	max := numeros[0]

	for _, num := range numeros {
		suma += num
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	promedio := suma / float64(len(numeros))
	return min, max, promedio
}

// Función variádica con tipos interface{}
func imprimir(valores ...interface{}) {
	for i, valor := range valores {
		fmt.Printf("  Valor %d: %v (tipo: %T)\n", i+1, valor, valor)
	}
}

func ejemplosVariadicas() {
	// Suma con diferentes cantidades de argumentos
	fmt.Printf("  Suma(): %d\n", sumarTodos())
	fmt.Printf("  Suma(5): %d\n", sumarTodos(5))
	fmt.Printf("  Suma(1,2,3): %d\n", sumarTodos(1, 2, 3))
	fmt.Printf("  Suma(1,2,3,4,5): %d\n", sumarTodos(1, 2, 3, 4, 5))

	// Expandir slice
	numeros := []int{10, 20, 30, 40}
	fmt.Printf("  Suma de slice: %d\n", sumarTodos(numeros...))

	// Saludar grupo
	saludarGrupo("Hola", "Ana", "Bruno", "Carmen")
	saludarGrupo("Buenos días", "David")

	// Estadísticas
	min, max, prom := estadisticas(1.5, 2.8, 3.1, 4.7, 2.2)
	fmt.Printf("  Estadísticas: Min=%.1f, Max=%.1f, Promedio=%.2f\n", min, max, prom)

	// Función con interface{}
	imprimir(42, "texto", true, 3.14, []int{1, 2, 3})
}

// ============= DEFER, PANIC Y RECOVER =============

func ejemploDeferBasico() {
	fmt.Println("    Inicio de función")
	defer fmt.Println("    Este mensaje se ejecuta al final (defer)")
	fmt.Println("    Fin de función")
}

func ejemploDeferMultiple() {
	fmt.Println("    Función con múltiples defer")
	defer fmt.Println("    Defer 1")
	defer fmt.Println("    Defer 2")
	defer fmt.Println("    Defer 3")
	fmt.Println("    Código normal")
}

func ejemploDeferConPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("    Recuperado de panic: %v\n", r)
		}
	}()

	fmt.Println("    Antes del panic")
	panic("¡Algo salió mal!")
	// Esta línea nunca se ejecuta debido al panic
}

func ejemploDeferPanicRecover() {
	fmt.Println("  Ejemplo defer básico:")
	ejemploDeferBasico()

	fmt.Println("\n  Ejemplo defer múltiple (LIFO):")
	ejemploDeferMultiple()

	fmt.Println("\n  Ejemplo defer con panic/recover:")
	ejemploDeferConPanic()
	fmt.Println("    Continúa la ejecución después del recover")

	// Defer con variables
	fmt.Println("\n  Defer con variables:")
	x := 10
	defer func(val int) {
		fmt.Printf("    Valor capturado en defer: %d\n", val)
	}(x)
	x = 20
	fmt.Printf("    Valor actual de x: %d\n", x)
}

func ejemploDeferRecursos() {
	fmt.Println("    Simulando manejo de archivo:")
	defer func() {
		fmt.Println("    Cerrando archivo (defer)")
	}()
	fmt.Println("    Abriendo archivo")
	fmt.Println("    Trabajando con archivo")
}

func ejemplosDeferPanicRecover() {
	fmt.Println("  Defer básico:")
	ejemploDeferBasico()

	fmt.Println("\n  Múltiples defer (orden LIFO):")
	ejemploDeferMultiple()

	fmt.Println("\n  Defer con panic y recover:")
	ejemploDeferConPanic()

	fmt.Println("\n  Defer con recursos:")
	ejemploDeferRecursos()
}

// ============= MÉTODOS Y RECEIVERS =============

// Struct para demostrar métodos
type Rectangulo struct {
	Ancho, Alto float64
}

// Método con value receiver
func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

// Método con value receiver que retorna string
func (r Rectangulo) String() string {
	return fmt.Sprintf("Rectángulo(%.1f x %.1f)", r.Ancho, r.Alto)
}

// Método con pointer receiver (puede modificar el struct)
func (r *Rectangulo) Escalar(factor float64) {
	r.Ancho *= factor
	r.Alto *= factor
}

// Método con pointer receiver que no modifica
func (r *Rectangulo) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

// Struct para cuenta bancaria
type CuentaBancaria struct {
	numero  string
	titular string
	balance float64
}

// Constructor (función que retorna struct)
func NuevaCuentaBancaria(numero, titular string, balanceInicial float64) *CuentaBancaria {
	return &CuentaBancaria{
		numero:  numero,
		titular: titular,
		balance: balanceInicial,
	}
}

// Métodos getter
func (c *CuentaBancaria) Numero() string   { return c.numero }
func (c *CuentaBancaria) Titular() string  { return c.titular }
func (c *CuentaBancaria) Balance() float64 { return c.balance }

// Métodos de negocio
func (c *CuentaBancaria) Depositar(cantidad float64) error {
	if cantidad <= 0 {
		return fmt.Errorf("la cantidad debe ser positiva")
	}
	c.balance += cantidad
	return nil
}

func (c *CuentaBancaria) Retirar(cantidad float64) error {
	if cantidad <= 0 {
		return fmt.Errorf("la cantidad debe ser positiva")
	}
	if cantidad > c.balance {
		return fmt.Errorf("fondos insuficientes")
	}
	c.balance -= cantidad
	return nil
}

func (c *CuentaBancaria) String() string {
	return fmt.Sprintf("Cuenta %s (%s): $%.2f", c.numero, c.titular, c.balance)
}

func ejemplosMetodos() {
	// Métodos con structs geométricos
	rect := Rectangulo{Ancho: 5, Alto: 3}
	fmt.Printf("  %s\n", rect.String())
	fmt.Printf("  Área: %.1f\n", rect.Area())
	fmt.Printf("  Perímetro: %.1f\n", rect.Perimetro())

	// Modificar con pointer receiver
	fmt.Println("  Escalando por 2:")
	rect.Escalar(2)
	fmt.Printf("  %s\n", rect.String())
	fmt.Printf("  Nueva área: %.1f\n", rect.Area())

	// Cuenta bancaria
	cuenta := NuevaCuentaBancaria("001", "Ana García", 1000.0)
	fmt.Printf("  %s\n", cuenta.String())

	// Operaciones
	if err := cuenta.Depositar(500); err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  Después del depósito: %s\n", cuenta.String())
	}

	if err := cuenta.Retirar(200); err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  Después del retiro: %s\n", cuenta.String())
	}

	// Intento de retiro con fondos insuficientes
	if err := cuenta.Retirar(2000); err != nil {
		fmt.Printf("  Error: %v\n", err)
	}
}
