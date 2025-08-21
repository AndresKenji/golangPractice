package main

import "fmt"

// Declaración de constantes a nivel de paquete
const PI float64 = 3.14159
const AppName string = "Mi Aplicación Go"

// Declaración múltiple de constantes
const (
	StatusActive   = "ACTIVE"
	StatusInactive = "INACTIVE"
	MaxRetries     = 3
	TimeoutSeconds = 30
)

// Constantes con iota (enumeraciones)
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// Variables a nivel de paquete
var globalCounter int
var isRunning bool

func main() {
	fmt.Println("=== EJEMPLOS DE VARIABLES Y CONSTANTES EN GO ===")
	fmt.Println()

	// 1. Declaración básica de variables
	fmt.Println("1. Declaración básica de variables:")
	var nombre string
	var edad int
	var activo bool
	fmt.Printf("Valores por defecto - nombre: '%s', edad: %d, activo: %t\n\n", nombre, edad, activo)

	// 2. Declaración con inicialización
	fmt.Println("2. Declaración con inicialización:")
	var nombre2 string = "Juan"
	var edad2 int = 25
	var activo2 bool = true
	fmt.Printf("nombre2: %s, edad2: %d, activo2: %t\n\n", nombre2, edad2, activo2)

	// 3. Declaración con inferencia de tipo
	fmt.Println("3. Declaración con inferencia de tipo:")
	var nombre3 = "María"
	var edad3 = 30
	var altura3 = 1.65
	fmt.Printf("nombre3: %s (tipo: %T), edad3: %d (tipo: %T), altura3: %.2f (tipo: %T)\n\n",
		nombre3, nombre3, edad3, edad3, altura3, altura3)

	// 4. Declaración corta con :=
	fmt.Println("4. Declaración corta con :=:")
	nombre4 := "Carlos"
	edad4 := 28
	esEstudiante := false
	fmt.Printf("nombre4: %s, edad4: %d, esEstudiante: %t\n\n", nombre4, edad4, esEstudiante)

	// 5. Declaración múltiple
	fmt.Println("5. Declaración múltiple:")
	var (
		ciudad     string = "Madrid"
		poblacion  int    = 3200000
		esCapital  bool   = true
		superficie float64
	)
	superficie = 604.3
	fmt.Printf("ciudad: %s, poblacion: %d, esCapital: %t, superficie: %.1f km²\n\n",
		ciudad, poblacion, esCapital, superficie)

	// 6. Asignación múltiple
	fmt.Println("6. Asignación múltiple:")
	x, y, z := 10, 20, 30
	fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z)

	// Intercambio de variables
	x, y = y, x
	fmt.Printf("Después del intercambio - x: %d, y: %d\n\n", x, y)

	// 7. Usando constantes
	fmt.Println("7. Usando constantes:")
	radio := 5.0
	area := PI * radio * radio
	fmt.Printf("Radio: %.1f, Área del círculo: %.2f\n", radio, area)
	fmt.Printf("Aplicación: %s\n", AppName)
	fmt.Printf("Estado: %s, Reintentos máximos: %d\n\n", StatusActive, MaxRetries)

	// 8. Constantes con iota
	fmt.Println("8. Constantes con iota (días de la semana):")
	fmt.Printf("Domingo: %d, Lunes: %d, Martes: %d\n", Sunday, Monday, Tuesday)
	fmt.Printf("Miércoles: %d, Jueves: %d, Viernes: %d, Sábado: %d\n\n", Wednesday, Thursday, Friday, Saturday)

	// 9. Variables con diferentes tipos
	fmt.Println("9. Variables con diferentes tipos:")
	var entero8 int8 = 127
	var entero16 int16 = 32767
	var entero32 int32 = 2147483647
	var entero64 int64 = 9223372036854775807

	var flotante32 float32 = 3.14
	var flotante64 float64 = 3.141592653589793

	var complejo64 complex64 = 1 + 2i
	var complejo128 complex128 = 2 + 3i

	fmt.Printf("int8: %d, int16: %d\n", entero8, entero16)
	fmt.Printf("int32: %d, int64: %d\n", entero32, entero64)
	fmt.Printf("float32: %.2f, float64: %.15f\n", flotante32, flotante64)
	fmt.Printf("complex64: %v, complex128: %v\n\n", complejo64, complejo128)

	// 10. Strings y runes
	fmt.Println("10. Strings y runes:")
	mensaje := "Hola, 世界"
	fmt.Printf("String: %s, Longitud en bytes: %d\n", mensaje, len(mensaje))

	runa := 'A'
	runaUnicode := '世'
	fmt.Printf("Runa A: %c (valor: %d), Runa 世: %c (valor: %d)\n\n", runa, runa, runaUnicode, runaUnicode)

	// 11. Punteros
	fmt.Println("11. Punteros:")
	numero := 42
	puntero := &numero
	fmt.Printf("Valor de numero: %d, Dirección: %p\n", numero, &numero)
	fmt.Printf("Puntero apunta a: %p, Valor al que apunta: %d\n", puntero, *puntero)

	*puntero = 100 // Modificar valor a través del puntero
	fmt.Printf("Después de modificar por puntero - numero: %d\n\n", numero)

	// 12. Arrays y Slices
	fmt.Println("12. Arrays y Slices:")
	var array [3]int = [3]int{1, 2, 3}
	slice := []int{4, 5, 6, 7, 8}

	fmt.Printf("Array: %v, Longitud: %d\n", array, len(array))
	fmt.Printf("Slice: %v, Longitud: %d, Capacidad: %d\n\n", slice, len(slice), cap(slice))

	// 13. Maps
	fmt.Println("13. Maps:")
	edades := map[string]int{
		"Ana":    25,
		"Bruno":  30,
		"Carmen": 28,
	}
	fmt.Printf("Edades: %v\n", edades)
	fmt.Printf("Edad de Ana: %d\n\n", edades["Ana"])

	// 14. Variables globales
	fmt.Println("14. Variables globales:")
	globalCounter = 5
	isRunning = true
	fmt.Printf("Contador global: %d, En ejecución: %t\n\n", globalCounter, isRunning)

	// 15. Scope de variables
	fmt.Println("15. Scope de variables:")
	variableExterna := "Externa"

	if true {
		variableInterna := "Interna"
		fmt.Printf("Dentro del bloque - Externa: %s, Interna: %s\n", variableExterna, variableInterna)
	}
	// variableInterna no está disponible aquí
	fmt.Printf("Fuera del bloque - Externa: %s\n\n", variableExterna)

	// 16. Constantes tipadas vs no tipadas
	fmt.Println("16. Constantes tipadas vs no tipadas:")
	const constanteNoTipada = 42
	const constanteTipada int = 42

	var flotante float64 = constanteNoTipada // Funciona
	var entero int = constanteTipada         // Funciona

	fmt.Printf("Usando constante no tipada como float64: %.1f\n", flotante)
	fmt.Printf("Usando constante tipada como int: %d\n", entero)

	fmt.Println("\n=== FIN DE LOS EJEMPLOS ===")
}
