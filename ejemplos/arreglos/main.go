package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("=== EJEMPLOS DE ARREGLOS Y SLICES EN GO ===")
	fmt.Println()

	// 1. Arreglos (Arrays) - Tamaño fijo
	fmt.Println("1. ARREGLOS (ARRAYS) - Tamaño fijo:")
	ejemplosArrays()
	fmt.Println()

	// 2. Slices - Tamaño dinámico
	fmt.Println("2. SLICES - Tamaño dinámico:")
	ejemplosSlices()
	fmt.Println()

	// 3. Creación de slices
	fmt.Println("3. CREACIÓN DE SLICES:")
	creacionSlices()
	fmt.Println()

	// 4. Operaciones con slices
	fmt.Println("4. OPERACIONES CON SLICES:")
	operacionesSlices()
	fmt.Println()

	// 5. Slicing (subslices)
	fmt.Println("5. SLICING (SUBSLICES):")
	ejemplosSlicing()
	fmt.Println()

	// 6. Append y capacidad
	fmt.Println("6. APPEND Y CAPACIDAD:")
	ejemplosAppend()
	fmt.Println()

	// 7. Copy de slices
	fmt.Println("7. COPY DE SLICES:")
	ejemplosCopy()
	fmt.Println()

	// 8. Slices multidimensionales
	fmt.Println("8. SLICES MULTIDIMENSIONALES:")
	slicesMultidimensionales()
	fmt.Println()

	// 9. Iteración
	fmt.Println("9. ITERACIÓN:")
	ejemplosIteracion()
	fmt.Println()

	// 10. Funciones útiles
	fmt.Println("10. FUNCIONES ÚTILES:")
	funcionesUtiles()
	fmt.Println()

	// 11. Casos de uso prácticos
	fmt.Println("11. CASOS DE USO PRÁCTICOS:")
	casosDeUsoPracticos()
	fmt.Println()

	// 12. Performance y mejores prácticas
	fmt.Println("12. PERFORMANCE Y MEJORES PRÁCTICAS:")
	mejoresPracticas()
	fmt.Println()

	fmt.Println("=== FIN DE LOS EJEMPLOS ===")
}

func ejemplosArrays() {
	// Declaración básica de arrays
	var numeros [5]int
	fmt.Printf("Array declarado (valores por defecto): %v\n", numeros)

	// Inicialización con valores
	var frutas [3]string = [3]string{"manzana", "banana", "naranja"}
	fmt.Printf("Array de frutas: %v\n", frutas)

	// Inicialización con inferencia de tamaño
	colores := [...]string{"rojo", "verde", "azul", "amarillo"}
	fmt.Printf("Array de colores (tamaño inferido): %v\n", colores)

	// Inicialización con índices específicos
	diasSemana := [7]string{
		0: "domingo",
		1: "lunes",
		6: "sábado",
	}
	fmt.Printf("Días de la semana (algunos índices): %v\n", diasSemana)

	// Acceso a elementos
	fmt.Printf("Primera fruta: %s\n", frutas[0])
	fmt.Printf("Último color: %s\n", colores[len(colores)-1])

	// Modificación de elementos
	numeros[0] = 10
	numeros[4] = 50
	fmt.Printf("Array modificado: %v\n", numeros)

	// Propiedades del array
	fmt.Printf("Tamaño del array colores: %d\n", len(colores))
	fmt.Printf("Tipo del array: %T\n", colores)

	// Arrays son valores, no referencias
	arrayOriginal := [3]int{1, 2, 3}
	arrayCopia := arrayOriginal
	arrayCopia[0] = 100
	fmt.Printf("Array original: %v\n", arrayOriginal)
	fmt.Printf("Array copia: %v\n", arrayCopia)
}

func ejemplosSlices() {
	// Declaración básica de slices
	var numeros []int
	fmt.Printf("Slice declarado (nil): %v, len: %d, cap: %d\n", numeros, len(numeros), cap(numeros))

	// Inicialización con make
	enteros := make([]int, 5) // longitud 5, capacidad 5
	fmt.Printf("Slice con make: %v, len: %d, cap: %d\n", enteros, len(enteros), cap(enteros))

	// Inicialización con make y capacidad específica
	flotantes := make([]float64, 3, 10) // longitud 3, capacidad 10
	fmt.Printf("Slice con capacidad: %v, len: %d, cap: %d\n", flotantes, len(flotantes), cap(flotantes))

	// Inicialización literal
	frutas := []string{"manzana", "banana", "cereza"}
	fmt.Printf("Slice literal: %v, len: %d, cap: %d\n", frutas, len(frutas), cap(frutas))

	// Slice de slice (slicing)
	array := [6]int{1, 2, 3, 4, 5, 6}
	slice1 := array[1:4] // elementos del índice 1 al 3
	slice2 := array[:3]  // elementos del índice 0 al 2
	slice3 := array[2:]  // elementos del índice 2 al final

	fmt.Printf("Array original: %v\n", array)
	fmt.Printf("Slice [1:4]: %v\n", slice1)
	fmt.Printf("Slice [:3]: %v\n", slice2)
	fmt.Printf("Slice [2:]: %v\n", slice3)

	// Los slices son referencias
	slice1[0] = 100
	fmt.Printf("Después de modificar slice1[0]: array = %v, slice1 = %v\n", array, slice1)
}

func creacionSlices() {
	// 1. Con make()
	slice1 := make([]int, 5)
	slice2 := make([]int, 3, 8)
	fmt.Printf("make([]int, 5): %v\n", slice1)
	fmt.Printf("make([]int, 3, 8): %v, cap: %d\n", slice2, cap(slice2))

	// 2. Literal
	slice3 := []string{"Go", "Python", "JavaScript"}
	fmt.Printf("Slice literal: %v\n", slice3)

	// 3. De un array
	array := [5]int{10, 20, 30, 40, 50}
	slice4 := array[1:4]
	fmt.Printf("De array[1:4]: %v\n", slice4)

	// 4. De otro slice
	slice5 := slice3[1:]
	fmt.Printf("De otro slice[1:]: %v\n", slice5)

	// 5. Slice vacío vs nil
	var sliceNil []int
	sliceVacio := []int{}
	sliceVacioMake := make([]int, 0)

	fmt.Printf("Slice nil: %v, es nil: %t\n", sliceNil, sliceNil == nil)
	fmt.Printf("Slice vacío: %v, es nil: %t\n", sliceVacio, sliceVacio == nil)
	fmt.Printf("Slice vacío con make: %v, es nil: %t\n", sliceVacioMake, sliceVacioMake == nil)

	// 6. Diferentes tipos de slices
	sliceBytes := []byte("Hola")
	sliceRunes := []rune("Mundo")
	sliceBool := []bool{true, false, true}

	fmt.Printf("Slice de bytes: %v\n", sliceBytes)
	fmt.Printf("Slice de runes: %v\n", sliceRunes)
	fmt.Printf("Slice de bool: %v\n", sliceBool)
}

func operacionesSlices() {
	// Slice inicial
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice inicial: %v\n", numeros)

	// Acceso a elementos
	fmt.Printf("Primer elemento: %d\n", numeros[0])
	fmt.Printf("Último elemento: %d\n", numeros[len(numeros)-1])

	// Modificación de elementos
	numeros[2] = 100
	fmt.Printf("Después de modificar índice 2: %v\n", numeros)

	// Verificar si slice está vacío
	var vacio []int
	fmt.Printf("¿Slice vacío?: %t\n", len(vacio) == 0)

	// Comparación de slices (no se puede usar ==)
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	sonIguales := compararSlices(slice1, slice2)
	fmt.Printf("¿Son iguales %v y %v?: %t\n", slice1, slice2, sonIguales)

	// Búsqueda en slice
	frutas := []string{"manzana", "banana", "cereza", "durazno"}
	indice := buscarEnSlice(frutas, "cereza")
	fmt.Printf("Índice de 'cereza' en %v: %d\n", frutas, indice)

	// Verificar si contiene elemento
	contiene := contieneElemento(frutas, "banana")
	fmt.Printf("¿Contiene 'banana'?: %t\n", contiene)
}

func ejemplosSlicing() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Slice original: %v\n", slice)

	// Sintaxis básica: slice[inicio:fin]
	fmt.Printf("slice[2:5]: %v\n", slice[2:5]) // elementos 2, 3, 4
	fmt.Printf("slice[:4]: %v\n", slice[:4])   // elementos 0, 1, 2, 3
	fmt.Printf("slice[6:]: %v\n", slice[6:])   // elementos 6, 7, 8, 9
	fmt.Printf("slice[:]: %v\n", slice[:])     // todos los elementos

	// Sintaxis extendida: slice[inicio:fin:capacidad]
	sub1 := slice[2:5:7] // elementos 2,3,4 con capacidad máxima hasta índice 7
	fmt.Printf("slice[2:5:7]: %v, len: %d, cap: %d\n", sub1, len(sub1), cap(sub1))

	// Slicing de slices
	subSlice := slice[3:7]
	fmt.Printf("Sub-slice [3:7]: %v\n", subSlice)

	subSubSlice := subSlice[1:3]
	fmt.Printf("Sub-sub-slice [1:3]: %v\n", subSubSlice)

	// Modificación afecta al slice original
	subSlice[0] = 999
	fmt.Printf("Después de modificar sub-slice: original = %v\n", slice)

	// Slicing con strings
	texto := "Hola Mundo"
	fmt.Printf("Texto: %s\n", texto)
	fmt.Printf("texto[0:4]: %s\n", texto[0:4])
	fmt.Printf("texto[5:]: %s\n", texto[5:])
}

func ejemplosAppend() {
	// Append básico
	var numeros []int
	fmt.Printf("Slice inicial: %v, len: %d, cap: %d\n", numeros, len(numeros), cap(numeros))

	numeros = append(numeros, 1)
	fmt.Printf("Después de append(1): %v, len: %d, cap: %d\n", numeros, len(numeros), cap(numeros))

	numeros = append(numeros, 2, 3, 4)
	fmt.Printf("Después de append(2,3,4): %v, len: %d, cap: %d\n", numeros, len(numeros), cap(numeros))

	// Append de otro slice
	masNumeros := []int{5, 6, 7}
	numeros = append(numeros, masNumeros...)
	fmt.Printf("Después de append slice: %v, len: %d, cap: %d\n", numeros, len(numeros), cap(numeros))

	// Observar crecimiento de capacidad
	fmt.Println("\nCrecimiento de capacidad:")
	slice := make([]int, 0, 1)

	for i := 0; i < 10; i++ {
		prevCap := cap(slice)
		slice = append(slice, i)
		if cap(slice) != prevCap {
			fmt.Printf("Capacidad cambió de %d a %d al agregar elemento %d\n", prevCap, cap(slice), i)
		}
	}
	fmt.Printf("Slice final: %v\n", slice)

	// Append con capacidad preallocada
	fmt.Println("\nCon capacidad preallocada:")
	slicePrealloc := make([]int, 0, 10)
	fmt.Printf("Inicial: len: %d, cap: %d\n", len(slicePrealloc), cap(slicePrealloc))

	for i := 0; i < 5; i++ {
		slicePrealloc = append(slicePrealloc, i)
	}
	fmt.Printf("Después de 5 append: %v, len: %d, cap: %d\n", slicePrealloc, len(slicePrealloc), cap(slicePrealloc))

	// Prepend (agregar al inicio)
	original := []int{3, 4, 5}
	prepend := []int{1, 2}
	resultado := append(prepend, original...)
	fmt.Printf("Prepend: %v + %v = %v\n", prepend, original, resultado)
}

func ejemplosCopy() {
	// Copy básico
	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source))

	n := copy(destination, source)
	fmt.Printf("Copiados %d elementos\n", n)
	fmt.Printf("Source: %v\n", source)
	fmt.Printf("Destination: %v\n", destination)

	// Modificar el destino no afecta el origen
	destination[0] = 999
	fmt.Printf("Después de modificar destination: source = %v, dest = %v\n", source, destination)

	// Copy con diferentes tamaños
	fmt.Println("\nCopy con diferentes tamaños:")

	// Destino más pequeño
	small := make([]int, 3)
	n1 := copy(small, source)
	fmt.Printf("Copy a slice pequeño: copiados %d elementos, resultado: %v\n", n1, small)

	// Destino más grande
	large := make([]int, 8)
	n2 := copy(large, source)
	fmt.Printf("Copy a slice grande: copiados %d elementos, resultado: %v\n", n2, large)

	// Copy overlapping
	fmt.Println("\nCopy overlapping:")
	overlap := []int{1, 2, 3, 4, 5}
	copy(overlap[2:], overlap[:3]) // Copia los primeros 3 elementos a partir del índice 2
	fmt.Printf("Después de copy overlapping: %v\n", overlap)

	// Copy strings a bytes
	texto := "Hola"
	bytes := make([]byte, len(texto))
	copy(bytes, texto)
	fmt.Printf("String a bytes: '%s' -> %v\n", texto, bytes)
}

func slicesMultidimensionales() {
	// Slice de slices (matriz)
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Matriz 3x3: %v\n", matriz)

	// Acceso a elementos
	fmt.Printf("Elemento [1][2]: %d\n", matriz[1][2])

	// Crear matriz dinámicamente
	filas, columnas := 3, 4
	matrizDinamica := make([][]int, filas)

	for i := range matrizDinamica {
		matrizDinamica[i] = make([]int, columnas)
		for j := range matrizDinamica[i] {
			matrizDinamica[i][j] = i*columnas + j + 1
		}
	}

	fmt.Printf("Matriz dinámica %dx%d:\n", filas, columnas)
	for _, fila := range matrizDinamica {
		fmt.Printf("%v\n", fila)
	}

	// Slice irregular (jagged array)
	irregular := [][]string{
		{"a"},
		{"b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i", "j"},
	}

	fmt.Println("Slice irregular:")
	for i, fila := range irregular {
		fmt.Printf("Fila %d: %v\n", i, fila)
	}

	// Slice 3D
	cubo := make([][][]int, 2)
	for i := range cubo {
		cubo[i] = make([][]int, 2)
		for j := range cubo[i] {
			cubo[i][j] = make([]int, 2)
			for k := range cubo[i][j] {
				cubo[i][j][k] = i*100 + j*10 + k
			}
		}
	}

	fmt.Printf("Cubo 2x2x2: %v\n", cubo)
}

func ejemplosIteracion() {
	numeros := []int{10, 20, 30, 40, 50}

	// Iteración con for tradicional
	fmt.Println("For tradicional:")
	for i := 0; i < len(numeros); i++ {
		fmt.Printf("numeros[%d] = %d\n", i, numeros[i])
	}

	// Iteración con range (índice y valor)
	fmt.Println("\nCon range (índice y valor):")
	for i, v := range numeros {
		fmt.Printf("índice: %d, valor: %d\n", i, v)
	}

	// Iteración con range (solo valor)
	fmt.Println("\nCon range (solo valor):")
	for _, v := range numeros {
		fmt.Printf("valor: %d\n", v)
	}

	// Iteración con range (solo índice)
	fmt.Println("\nCon range (solo índice):")
	for i := range numeros {
		fmt.Printf("índice: %d\n", i)
	}

	// Iteración reversa
	fmt.Println("\nIteración reversa:")
	for i := len(numeros) - 1; i >= 0; i-- {
		fmt.Printf("numeros[%d] = %d\n", i, numeros[i])
	}

	// Iteración con condición
	fmt.Println("\nCon condición (números > 25):")
	for _, v := range numeros {
		if v > 25 {
			fmt.Printf("valor: %d\n", v)
		}
	}

	// Iteración anidada para slice 2D
	matriz := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("\nIteración 2D:")
	for i, fila := range matriz {
		for j, valor := range fila {
			fmt.Printf("matriz[%d][%d] = %d\n", i, j, valor)
		}
	}
}

func funcionesUtiles() {
	numeros := []int{5, 2, 8, 1, 9, 3}
	fmt.Printf("Slice original: %v\n", numeros)

	// Ordenar slice
	sort.Ints(numeros)
	fmt.Printf("Después de sort.Ints: %v\n", numeros)

	// Buscar en slice ordenado
	target := 5
	index := sort.SearchInts(numeros, target)
	fmt.Printf("Índice de %d en slice ordenado: %d\n", target, index)

	// Strings
	palabras := []string{"zebra", "apple", "banana", "cherry"}
	fmt.Printf("Palabras originales: %v\n", palabras)

	sort.Strings(palabras)
	fmt.Printf("Palabras ordenadas: %v\n", palabras)

	// Join strings
	frase := strings.Join(palabras, " - ")
	fmt.Printf("Join con ' - ': %s\n", frase)

	// Split string
	texto := "Go,Python,JavaScript,Rust"
	lenguajes := strings.Split(texto, ",")
	fmt.Printf("Split de '%s': %v\n", texto, lenguajes)

	// Reversar slice
	reverseSlice(numeros)
	fmt.Printf("Números reversados: %v\n", numeros)

	// Filtrar slice
	pares := filtrarPares([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("Números pares: %v\n", pares)

	// Map función
	cuadrados := mapearCuadrados([]int{1, 2, 3, 4, 5})
	fmt.Printf("Cuadrados de [1,2,3,4,5]: %v\n", cuadrados)

	// Reducir slice
	suma := reducirSuma([]int{1, 2, 3, 4, 5})
	fmt.Printf("Suma de [1,2,3,4,5]: %d\n", suma)
}

func casosDeUsoPracticos() {
	fmt.Println("--- Stack (Pila) ---")
	stack := make([]int, 0)

	// Push
	stack = append(stack, 1, 2, 3)
	fmt.Printf("Después de push 1,2,3: %v\n", stack)

	// Pop
	if len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("Pop: %d, stack: %v\n", top, stack)
	}

	fmt.Println("\n--- Queue (Cola) ---")
	queue := make([]string, 0)

	// Enqueue
	queue = append(queue, "primero", "segundo", "tercero")
	fmt.Printf("Después de enqueue: %v\n", queue)

	// Dequeue
	if len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		fmt.Printf("Dequeue: %s, queue: %v\n", front, queue)
	}

	fmt.Println("\n--- Buffer circular ---")
	buffer := make([]int, 5)
	head, tail, size := 0, 0, 0

	// Agregar elementos
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7} {
		if size < len(buffer) {
			buffer[tail] = v
			tail = (tail + 1) % len(buffer)
			size++
		} else {
			buffer[tail] = v
			tail = (tail + 1) % len(buffer)
			head = (head + 1) % len(buffer)
		}
		fmt.Printf("Agregar %d: buffer=%v, head=%d, tail=%d\n", v, buffer, head, tail)
	}

	fmt.Println("\n--- Procesamiento de lotes ---")
	datos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	tamanoLote := 3

	for i := 0; i < len(datos); i += tamanoLote {
		fin := i + tamanoLote
		if fin > len(datos) {
			fin = len(datos)
		}
		lote := datos[i:fin]
		fmt.Printf("Lote %d: %v\n", i/tamanoLote+1, lote)
	}

	fmt.Println("\n--- Eliminar elementos ---")
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original: %v\n", numeros)

	// Eliminar por índice (elemento en posición 3)
	indiceEliminar := 3
	numeros = append(numeros[:indiceEliminar], numeros[indiceEliminar+1:]...)
	fmt.Printf("Después de eliminar índice 3: %v\n", numeros)

	// Eliminar por valor
	valorEliminar := 7
	for i, v := range numeros {
		if v == valorEliminar {
			numeros = append(numeros[:i], numeros[i+1:]...)
			break
		}
	}
	fmt.Printf("Después de eliminar valor 7: %v\n", numeros)
}

func mejoresPracticas() {
	fmt.Println("--- Pre-allocar capacidad ---")

	// Mal: sin pre-allocar
	var malo []int
	for i := 0; i < 1000; i++ {
		malo = append(malo, i)
	}

	// Bien: pre-allocar
	bueno := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		bueno = append(bueno, i)
	}
	fmt.Println("Pre-allocar capacidad mejora el rendimiento")

	fmt.Println("\n--- Evitar memory leaks ---")
	// Problema: mantener referencia a slice grande
	datosGrandes := make([]byte, 1000000)
	// ... llenar datos ...

	// Mal: mantiene referencia al slice grande
	//pequenoMalo := datosGrandes[0:10]

	// Bien: copiar solo lo necesario
	pequenoBueno := make([]byte, 10)
	copy(pequenoBueno, datosGrandes[0:10])
	fmt.Printf("Copiado slice pequeño: len=%d\n", len(pequenoBueno))

	fmt.Println("\n--- Slicing eficiente ---")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Para evitar memory leaks al hacer slicing
	subSlice := make([]int, 3)
	copy(subSlice, slice[2:5])
	fmt.Printf("Sub-slice copiado: %v\n", subSlice)

	fmt.Println("\n--- Verificar nil vs empty ---")
	var nilSlice []int
	emptySlice := []int{}

	fmt.Printf("nil slice: %v, len: %d, es nil: %t\n", nilSlice, len(nilSlice), nilSlice == nil)
	fmt.Printf("empty slice: %v, len: %d, es nil: %t\n", emptySlice, len(emptySlice), emptySlice == nil)

	// Ambos son seguros para iterar (ranging over nil slice is safe)
	for range nilSlice {
		// No se ejecuta porque el slice es nil
	}
	for range emptySlice {
		// No se ejecuta porque el slice está vacío
	}

	fmt.Println("Ambos slices son seguros para usar con range")
}

// Funciones auxiliares
func compararSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func buscarEnSlice(slice []string, target string) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func contieneElemento(slice []string, target string) bool {
	return buscarEnSlice(slice, target) != -1
}

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func filtrarPares(slice []int) []int {
	var resultado []int
	for _, v := range slice {
		if v%2 == 0 {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

func mapearCuadrados(slice []int) []int {
	resultado := make([]int, len(slice))
	for i, v := range slice {
		resultado[i] = v * v
	}
	return resultado
}

func reducirSuma(slice []int) int {
	suma := 0
	for _, v := range slice {
		suma += v
	}
	return suma
}
