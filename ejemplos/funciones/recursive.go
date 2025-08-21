package main

import (
	"fmt"
	"strings"
)

// ============= EJEMPLOS DE RECURSIVIDAD =============

func runRecursiveExamples() {
	fmt.Println("=== EJEMPLOS DE RECURSIVIDAD EN GO ===")
	fmt.Println()

	// 1. Recursividad básica
	fmt.Println("1. FACTORIAL:")
	ejemplosFactorial()
	fmt.Println()

	// 2. Fibonacci
	fmt.Println("2. FIBONACCI:")
	ejemplosFibonacci()
	fmt.Println()

	// 3. Suma de dígitos
	fmt.Println("3. SUMA DE DÍGITOS:")
	ejemplosSumaDigitos()
	fmt.Println()

	// 4. Potencia
	fmt.Println("4. POTENCIA:")
	ejemplosPotencia()
	fmt.Println()

	// 5. Máximo común divisor
	fmt.Println("5. MÁXIMO COMÚN DIVISOR (MCD):")
	ejemplosMCD()
	fmt.Println()

	// 6. Torres de Hanoi
	fmt.Println("6. TORRES DE HANOI:")
	ejemplosTorresHanoi()
	fmt.Println()

	// 7. Recursividad con slices
	fmt.Println("7. RECURSIVIDAD CON SLICES:")
	ejemplosSlices()
	fmt.Println()

	// 8. Recursividad con strings
	fmt.Println("8. RECURSIVIDAD CON STRINGS:")
	ejemplosStrings()
	fmt.Println()

	// 9. Búsqueda binaria
	fmt.Println("9. BÚSQUEDA BINARIA:")
	ejemplosBusquedaBinaria()
	fmt.Println()

	// 10. Estructuras de datos recursivas
	fmt.Println("10. ESTRUCTURAS DE DATOS RECURSIVAS:")
	ejemplosEstructurasRecursivas()
	fmt.Println()

	fmt.Println("=== FIN DE LOS EJEMPLOS DE RECURSIVIDAD ===")
}

// ============= FACTORIAL =============

// Factorial básico
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Factorial con validación
func factorialSeguro(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("el factorial no está definido para números negativos")
	}
	if n > 20 {
		return 0, fmt.Errorf("factorial demasiado grande (overflow)")
	}

	if n <= 1 {
		return 1, nil
	}

	result, err := factorialSeguro(n - 1)
	if err != nil {
		return 0, err
	}

	return n * result, nil
}

func ejemplosFactorial() {
	numeros := []int{0, 1, 5, 7, 10}

	for _, n := range numeros {
		result := factorial(n)
		fmt.Printf("  %d! = %d\n", n, result)
	}

	fmt.Println("\n  Factorial con validación:")
	numerosTest := []int{-1, 5, 25}
	for _, n := range numerosTest {
		if result, err := factorialSeguro(n); err != nil {
			fmt.Printf("  %d!: Error - %v\n", n, err)
		} else {
			fmt.Printf("  %d! = %d\n", n, result)
		}
	}
}

// ============= FIBONACCI =============

// Fibonacci básico (ineficiente)
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Fibonacci con memoización
func fibonacciMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	result := fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	memo[n] = result
	return result
}

// Wrapper para fibonacci con memoización
func fibonacciConMemo(n int) int {
	memo := make(map[int]int)
	return fibonacciMemo(n, memo)
}

func ejemplosFibonacci() {
	fmt.Println("  Fibonacci básico:")
	for i := 0; i <= 10; i++ {
		result := fibonacci(i)
		fmt.Printf("  F(%d) = %d\n", i, result)
	}

	fmt.Println("\n  Fibonacci con memoización (más eficiente):")
	for i := 0; i <= 15; i++ {
		result := fibonacciConMemo(i)
		fmt.Printf("  F(%d) = %d\n", i, result)
	}

	// Comparación de rendimiento implícita
	fmt.Println("\n  (La memoización permite calcular números más grandes eficientemente)")
}

// ============= SUMA DE DÍGITOS =============

// Suma de dígitos de un número
func sumaDigitos(n int) int {
	if n < 10 {
		return n
	}
	return (n % 10) + sumaDigitos(n/10)
}

// Suma de dígitos hasta llegar a un solo dígito
func sumaDigitosUnico(n int) int {
	suma := sumaDigitos(n)
	if suma < 10 {
		return suma
	}
	return sumaDigitosUnico(suma)
}

func ejemplosSumaDigitos() {
	numeros := []int{123, 456, 789, 9876, 12345}

	for _, num := range numeros {
		suma := sumaDigitos(num)
		fmt.Printf("  Suma dígitos de %d = %d\n", num, suma)
	}

	fmt.Println("\n  Suma hasta dígito único:")
	for _, num := range numeros {
		unico := sumaDigitosUnico(num)
		fmt.Printf("  %d -> %d\n", num, unico)
	}
}

// ============= POTENCIA =============

// Potencia básica
func potencia(base, exponente int) int {
	if exponente == 0 {
		return 1
	}
	if exponente == 1 {
		return base
	}
	return base * potencia(base, exponente-1)
}

// Potencia optimizada (exponenciación rápida)
func potenciaRapida(base, exponente int) int {
	if exponente == 0 {
		return 1
	}
	if exponente == 1 {
		return base
	}

	if exponente%2 == 0 {
		mitad := potenciaRapida(base, exponente/2)
		return mitad * mitad
	} else {
		return base * potenciaRapida(base, exponente-1)
	}
}

func ejemplosPotencia() {
	casos := []struct{ base, exp int }{
		{2, 5}, {3, 4}, {5, 3}, {10, 3},
	}

	fmt.Println("  Potencia básica:")
	for _, caso := range casos {
		result := potencia(caso.base, caso.exp)
		fmt.Printf("  %d^%d = %d\n", caso.base, caso.exp, result)
	}

	fmt.Println("\n  Potencia rápida (optimizada):")
	for _, caso := range casos {
		result := potenciaRapida(caso.base, caso.exp)
		fmt.Printf("  %d^%d = %d\n", caso.base, caso.exp, result)
	}
}

// ============= MÁXIMO COMÚN DIVISOR =============

// MCD usando algoritmo de Euclides
func mcd(a, b int) int {
	if b == 0 {
		return a
	}
	return mcd(b, a%b)
}

// MCM usando MCD
func mcm(a, b int) int {
	return (a * b) / mcd(a, b)
}

func ejemplosMCD() {
	pares := []struct{ a, b int }{
		{48, 18}, {56, 42}, {17, 13}, {100, 25},
	}

	for _, par := range pares {
		mcdResult := mcd(par.a, par.b)
		mcmResult := mcm(par.a, par.b)
		fmt.Printf("  MCD(%d, %d) = %d, MCM(%d, %d) = %d\n",
			par.a, par.b, mcdResult, par.a, par.b, mcmResult)
	}
}

// ============= TORRES DE HANOI =============

// Resolver Torres de Hanoi
func torresHanoi(n int, origen, destino, auxiliar string) {
	if n == 1 {
		fmt.Printf("    Mover disco 1 de %s a %s\n", origen, destino)
		return
	}

	// Mover n-1 discos de origen a auxiliar
	torresHanoi(n-1, origen, auxiliar, destino)

	// Mover el disco más grande de origen a destino
	fmt.Printf("    Mover disco %d de %s a %s\n", n, origen, destino)

	// Mover n-1 discos de auxiliar a destino
	torresHanoi(n-1, auxiliar, destino, origen)
}

// Contar movimientos necesarios
func contarMovimientosTorres(n int) int {
	if n == 1 {
		return 1
	}
	return 2*contarMovimientosTorres(n-1) + 1
}

func ejemplosTorresHanoi() {
	discos := []int{1, 2, 3, 4}

	for _, n := range discos {
		movimientos := contarMovimientosTorres(n)
		fmt.Printf("  Torres de Hanoi con %d disco(s) - %d movimientos:\n", n, movimientos)
		if n <= 3 { // Solo mostrar secuencia para casos pequeños
			torresHanoi(n, "A", "C", "B")
		}
		fmt.Println()
	}
}

// ============= RECURSIVIDAD CON SLICES =============

// Suma de elementos en slice
func sumaSlice(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	if len(slice) == 1 {
		return slice[0]
	}
	return slice[0] + sumaSlice(slice[1:])
}

// Encontrar máximo en slice
func maximoSlice(slice []int) int {
	if len(slice) == 1 {
		return slice[0]
	}

	maxResto := maximoSlice(slice[1:])
	if slice[0] > maxResto {
		return slice[0]
	}
	return maxResto
}

// Invertir slice
func invertirSlice(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	return append(invertirSlice(slice[1:]), slice[0])
}

// Verificar si slice está ordenado
func estaOrdenado(slice []int) bool {
	if len(slice) <= 1 {
		return true
	}
	if slice[0] > slice[1] {
		return false
	}
	return estaOrdenado(slice[1:])
}

func ejemplosSlices() {
	numeros := []int{3, 7, 1, 9, 4, 6, 2}

	suma := sumaSlice(numeros)
	fmt.Printf("  Slice: %v\n", numeros)
	fmt.Printf("  Suma: %d\n", suma)

	max := maximoSlice(numeros)
	fmt.Printf("  Máximo: %d\n", max)

	invertido := invertirSlice(numeros)
	fmt.Printf("  Invertido: %v\n", invertido)

	ordenado := estaOrdenado(numeros)
	fmt.Printf("  ¿Está ordenado?: %t\n", ordenado)

	numerosOrdenados := []int{1, 2, 3, 4, 5}
	ordenado2 := estaOrdenado(numerosOrdenados)
	fmt.Printf("  %v ¿está ordenado?: %t\n", numerosOrdenados, ordenado2)
}

// ============= RECURSIVIDAD CON STRINGS =============

// Invertir string
func invertirString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return invertirString(s[1:]) + string(s[0])
}

// Verificar si es palíndromo
func esPalindromo(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return esPalindromo(s[1 : len(s)-1])
}

// Contar vocales
func contarVocales(s string) int {
	if len(s) == 0 {
		return 0
	}

	vocales := "aeiouAEIOU"
	count := 0
	if strings.ContainsRune(vocales, rune(s[0])) {
		count = 1
	}

	return count + contarVocales(s[1:])
}

func ejemplosStrings() {
	textos := []string{
		"Hola",
		"reconocer",
		"A man a plan a canal Panama",
		"programming",
	}

	for _, texto := range textos {
		invertido := invertirString(texto)
		palindromo := esPalindromo(texto)
		vocales := contarVocales(texto)

		fmt.Printf("  Texto: \"%s\"\n", texto)
		fmt.Printf("    Invertido: \"%s\"\n", invertido)
		fmt.Printf("    ¿Palíndromo?: %t\n", palindromo)
		fmt.Printf("    Vocales: %d\n", vocales)
		fmt.Println()
	}
}

// ============= BÚSQUEDA BINARIA =============

// Búsqueda binaria recursiva
func busquedaBinaria(slice []int, target, inicio, fin int) int {
	if inicio > fin {
		return -1 // No encontrado
	}

	medio := (inicio + fin) / 2

	if slice[medio] == target {
		return medio
	}

	if target < slice[medio] {
		return busquedaBinaria(slice, target, inicio, medio-1)
	} else {
		return busquedaBinaria(slice, target, medio+1, fin)
	}
}

// Wrapper para búsqueda binaria
func buscar(slice []int, target int) int {
	return busquedaBinaria(slice, target, 0, len(slice)-1)
}

func ejemplosBusquedaBinaria() {
	numeros := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	targets := []int{1, 7, 13, 20, 5}

	fmt.Printf("  Array ordenado: %v\n", numeros)
	fmt.Println("  Búsquedas:")

	for _, target := range targets {
		indice := buscar(numeros, target)
		if indice != -1 {
			fmt.Printf("    %d encontrado en índice %d\n", target, indice)
		} else {
			fmt.Printf("    %d no encontrado\n", target)
		}
	}
}

// ============= ESTRUCTURAS DE DATOS RECURSIVAS =============

// Nodo de árbol binario
type Nodo struct {
	Valor     int
	Izquierdo *Nodo
	Derecho   *Nodo
}

// Insertar en árbol binario de búsqueda
func (n *Nodo) Insertar(valor int) *Nodo {
	if n == nil {
		return &Nodo{Valor: valor}
	}

	if valor < n.Valor {
		n.Izquierdo = n.Izquierdo.Insertar(valor)
	} else if valor > n.Valor {
		n.Derecho = n.Derecho.Insertar(valor)
	}

	return n
}

// Buscar en árbol
func (n *Nodo) Buscar(valor int) bool {
	if n == nil {
		return false
	}

	if valor == n.Valor {
		return true
	}

	if valor < n.Valor {
		return n.Izquierdo.Buscar(valor)
	} else {
		return n.Derecho.Buscar(valor)
	}
}

// Recorrido inorden (izquierdo, raíz, derecho)
func (n *Nodo) RecorridoInorden() []int {
	if n == nil {
		return []int{}
	}

	var resultado []int
	resultado = append(resultado, n.Izquierdo.RecorridoInorden()...)
	resultado = append(resultado, n.Valor)
	resultado = append(resultado, n.Derecho.RecorridoInorden()...)

	return resultado
}

// Altura del árbol
func (n *Nodo) Altura() int {
	if n == nil {
		return 0
	}

	alturaIzq := n.Izquierdo.Altura()
	alturaDer := n.Derecho.Altura()

	if alturaIzq > alturaDer {
		return alturaIzq + 1
	} else {
		return alturaDer + 1
	}
}

// Contar nodos
func (n *Nodo) ContarNodos() int {
	if n == nil {
		return 0
	}
	return 1 + n.Izquierdo.ContarNodos() + n.Derecho.ContarNodos()
}

func ejemplosEstructurasRecursivas() {
	// Crear árbol binario de búsqueda
	var raiz *Nodo
	valores := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}

	fmt.Printf("  Insertando valores: %v\n", valores)

	for _, valor := range valores {
		raiz = raiz.Insertar(valor)
	}

	// Recorrido inorden (muestra valores ordenados)
	inorden := raiz.RecorridoInorden()
	fmt.Printf("  Recorrido inorden: %v\n", inorden)

	// Propiedades del árbol
	altura := raiz.Altura()
	nodos := raiz.ContarNodos()
	fmt.Printf("  Altura del árbol: %d\n", altura)
	fmt.Printf("  Número de nodos: %d\n", nodos)

	// Búsquedas
	busquedas := []int{40, 55, 80, 90}
	fmt.Println("  Búsquedas:")
	for _, valor := range busquedas {
		encontrado := raiz.Buscar(valor)
		fmt.Printf("    Buscar %d: %t\n", valor, encontrado)
	}
}
