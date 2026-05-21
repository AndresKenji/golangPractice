package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== EJEMPLOS DE CICLOS EN GO ===\n")

	// 1. FOR básico (tradicional)
	fmt.Println("1. FOR básico (tradicional):")
	fmt.Print("   Conteo del 1 al 5: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println()

	// 2. FOR con múltiples variables
	fmt.Println("2. FOR con múltiples variables:")
	fmt.Print("   Conteo simultáneo: ")
	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("(%d,%d) ", i, j)
	}
	fmt.Println()
	fmt.Println()

	// 3. FOR como WHILE
	fmt.Println("3. FOR como WHILE:")
	contador := 1
	fmt.Print("   Potencias de 2: ")
	for contador <= 16 {
		fmt.Printf("%d ", contador)
		contador *= 2
	}
	fmt.Println()
	fmt.Println()

	// 4. FOR infinito con break
	fmt.Println("4. FOR infinito con break:")
	fmt.Print("   Números aleatorios hasta encontrar uno > 8: ")
	rand.Seed(time.Now().UnixNano())
	for {
		numero := rand.Intn(10)
		fmt.Printf("%d ", numero)
		if numero > 8 {
			fmt.Printf("(¡Encontrado!)")
			break
		}
	}
	fmt.Println()
	fmt.Println()

	// 5. FOR con continue
	fmt.Println("5. FOR con continue (solo números pares):")
	fmt.Print("   Números pares del 1 al 10: ")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // Salta los números impares
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Println()

	// 6. FOR con range - Arrays
	fmt.Println("6. FOR con range - Arrays:")
	colores := [4]string{"rojo", "verde", "azul", "amarillo"}

	fmt.Println("   Índices y valores:")
	for indice, valor := range colores {
		fmt.Printf("   [%d]: %s\n", indice, valor)
	}

	fmt.Print("   Solo valores: ")
	for _, valor := range colores {
		fmt.Printf("%s ", valor)
	}
	fmt.Println()

	fmt.Print("   Solo índices: ")
	for indice := range colores {
		fmt.Printf("%d ", indice)
	}
	fmt.Println()
	fmt.Println()

	// 7. FOR con range - Slices
	fmt.Println("7. FOR con range - Slices:")
	numeros := []int{10, 20, 30, 40, 50}

	fmt.Print("   Suma de elementos: ")
	suma := 0
	for _, numero := range numeros {
		suma += numero
		fmt.Printf("%d ", numero)
	}
	fmt.Printf("= %d\n", suma)
	fmt.Println()

	// 8. FOR con range - Maps
	fmt.Println("8. FOR con range - Maps:")
	edades := map[string]int{
		"Ana":    25,
		"Carlos": 30,
		"Elena":  22,
		"Diego":  28,
	}

	fmt.Println("   Personas y edades:")
	for nombre, edad := range edades {
		fmt.Printf("   %s tiene %d años\n", nombre, edad)
	}
	fmt.Println()

	// 9. FOR con range - Strings (caracteres)
	fmt.Println("9. FOR con range - Strings:")
	palabra := "Golang"

	fmt.Println("   Caracteres por posición:")
	for indice, caracter := range palabra {
		fmt.Printf("   [%d]: %c (Unicode: %d)\n", indice, caracter, caracter)
	}
	fmt.Println()

	// 10. FOR anidados - Tabla de multiplicar
	fmt.Println("10. FOR anidados - Tabla de multiplicar (3x3):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("   %d x %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}

	// 11. FOR anidados - Patrón de asteriscos
	fmt.Println("11. FOR anidados - Patrón de asteriscos:")
	for i := 1; i <= 5; i++ {
		fmt.Print("   ")
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()

	// 12. Labels y break/continue con etiquetas
	fmt.Println("12. Labels y break/continue con etiquetas:")
	fmt.Println("   Búsqueda en matriz 3x3:")
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	objetivo := 5
	encontrado := false

buscar:
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("   Revisando posición [%d][%d] = %d\n", i, j, matriz[i][j])
			if matriz[i][j] == objetivo {
				fmt.Printf("   ¡Encontrado %d en posición [%d][%d]!\n", objetivo, i, j)
				encontrado = true
				break buscar // Sale de ambos loops
			}
		}
	}

	if !encontrado {
		fmt.Printf("   No se encontró %d\n", objetivo)
	}
	fmt.Println()

	// 13. FOR con channels (ejemplo básico)
	fmt.Println("13. FOR con channels:")
	canal := make(chan int, 5)

	// Enviar datos al canal
	go func() {
		for i := 1; i <= 5; i++ {
			canal <- i * i // Enviar cuadrados
		}
		close(canal)
	}()

	fmt.Print("   Cuadrados recibidos del canal: ")
	for valor := range canal {
		fmt.Printf("%d ", valor)
	}
	fmt.Println()
	fmt.Println()

	// 14. Diferentes formas de iterar un slice
	fmt.Println("14. Diferentes formas de iterar un slice:")
	frutas := []string{"manzana", "banana", "naranja", "uva"}

	fmt.Println("   Método 1 - FOR tradicional:")
	for i := 0; i < len(frutas); i++ {
		fmt.Printf("   [%d]: %s\n", i, frutas[i])
	}

	fmt.Println("   Método 2 - FOR con range:")
	for i, fruta := range frutas {
		fmt.Printf("   [%d]: %s\n", i, fruta)
	}

	fmt.Println("   Método 3 - Solo valores:")
	for _, fruta := range frutas {
		fmt.Printf("   %s ", fruta)
	}
	fmt.Println()
	fmt.Println()

	// 15. Iteración inversa
	fmt.Println("15. Iteración inversa:")
	fmt.Print("   Números del 10 al 1: ")
	for i := 10; i >= 1; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Print("   Array en orden inverso: ")
	for i := len(frutas) - 1; i >= 0; i-- {
		fmt.Printf("%s ", frutas[i])
	}
	fmt.Println()
	fmt.Println()

	// 16. Ejemplo práctico: Filtrar y procesar datos
	fmt.Println("16. Ejemplo práctico: Filtrar y procesar datos:")
	estudiantes := []Estudiante{
		{"Ana", 85},
		{"Carlos", 92},
		{"Elena", 78},
		{"Diego", 88},
		{"Sofia", 95},
	}

	fmt.Println("   Estudiantes con calificación >= 85:")
	aprobados := 0
	for _, estudiante := range estudiantes {
		if estudiante.Calificacion >= 85 {
			fmt.Printf("   %s: %d\n", estudiante.Nombre, estudiante.Calificacion)
			aprobados++
		}
	}
	fmt.Printf("   Total aprobados: %d de %d\n", aprobados, len(estudiantes))
	fmt.Println()

	// 17. Ejemplo práctico: Búsqueda y modificación
	fmt.Println("17. Ejemplo práctico: Búsqueda y modificación:")
	inventario := map[string]int{
		"laptops":   10,
		"ratones":   25,
		"teclados":  15,
		"monitores": 8,
	}

	fmt.Println("   Inventario original:")
	for producto, cantidad := range inventario {
		fmt.Printf("   %s: %d unidades\n", producto, cantidad)
	}

	// Actualizar inventario
	fmt.Println("   Aplicando descuentos (restando 2 unidades):")
	for producto := range inventario {
		if inventario[producto] > 2 {
			inventario[producto] -= 2
		}
	}

	for producto, cantidad := range inventario {
		fmt.Printf("   %s: %d unidades\n", producto, cantidad)
	}
	fmt.Println()

	// 18. Medición de rendimiento de ciclos
	fmt.Println("18. Medición de rendimiento:")
	medirRendimiento()

	fmt.Println("=== FIN DE EJEMPLOS DE CICLOS ===")
}

// Estructura para el ejemplo de estudiantes
type Estudiante struct {
	Nombre       string
	Calificacion int
}

// Función para medir rendimiento de diferentes tipos de ciclos
func medirRendimiento() {
	slice := make([]int, 1000000)
	for i := range slice {
		slice[i] = i
	}

	// Método 1: FOR tradicional
	inicio := time.Now()
	suma1 := 0
	for i := 0; i < len(slice); i++ {
		suma1 += slice[i]
	}
	tiempo1 := time.Since(inicio)

	// Método 2: FOR con range
	inicio = time.Now()
	suma2 := 0
	for _, valor := range slice {
		suma2 += valor
	}
	tiempo2 := time.Since(inicio)

	fmt.Printf("   FOR tradicional: %v (suma: %d)\n", tiempo1, suma1)
	fmt.Printf("   FOR con range: %v (suma: %d)\n", tiempo2, suma2)
	fmt.Println()
}
