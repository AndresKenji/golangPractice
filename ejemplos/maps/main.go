package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("=== EJEMPLOS DE MAPS EN GO ===")
	fmt.Println()

	// 1. Declaración e inicialización básica
	fmt.Println("1. DECLARACIÓN E INICIALIZACIÓN BÁSICA:")
	declaracionBasica()
	fmt.Println()

	// 2. Operaciones fundamentales
	fmt.Println("2. OPERACIONES FUNDAMENTALES:")
	operacionesFundamentales()
	fmt.Println()

	// 3. Verificación de existencia de llaves
	fmt.Println("3. VERIFICACIÓN DE EXISTENCIA DE LLAVES:")
	verificacionLlaves()
	fmt.Println()

	// 4. Iteración sobre maps
	fmt.Println("4. ITERACIÓN SOBRE MAPS:")
	iteracionMaps()
	fmt.Println()

	// 5. Maps anidados
	fmt.Println("5. MAPS ANIDADOS:")
	mapsAnidados()
	fmt.Println()

	// 6. Maps con diferentes tipos de datos
	fmt.Println("6. MAPS CON DIFERENTES TIPOS DE DATOS:")
	tiposDeDatos()
	fmt.Println()

	// 7. Maps con structs
	fmt.Println("7. MAPS CON STRUCTS:")
	mapsConStructs()
	fmt.Println()

	// 8. Maps como sets (conjuntos)
	fmt.Println("8. MAPS COMO SETS (CONJUNTOS):")
	mapsComoSets()
	fmt.Println()

	// 9. Contadores y frecuencias
	fmt.Println("9. CONTADORES Y FRECUENCIAS:")
	contadoresYFrecuencias()
	fmt.Println()

	// 10. Agrupación de datos
	fmt.Println("10. AGRUPACIÓN DE DATOS:")
	agrupacionDatos()
	fmt.Println()

	// 11. Maps ordenados
	fmt.Println("11. MAPS ORDENADOS:")
	mapsOrdenados()
	fmt.Println()

	// 12. Operaciones avanzadas
	fmt.Println("12. OPERACIONES AVANZADAS:")
	operacionesAvanzadas()
	fmt.Println()

	// 13. Casos de uso prácticos
	fmt.Println("13. CASOS DE USO PRÁCTICOS:")
	casosDeUsoPracticos()
	fmt.Println()

	// 14. Performance y mejores prácticas
	fmt.Println("14. PERFORMANCE Y MEJORES PRÁCTICAS:")
	mejoresPracticas()
	fmt.Println()

	fmt.Println("=== FIN DE LOS EJEMPLOS ===")
}

func declaracionBasica() {
	// Declaración con make
	edades := make(map[string]int)
	fmt.Printf("Map vacío con make: %v\n", edades)

	// Inicialización literal
	capitales := map[string]string{
		"España":   "Madrid",
		"Francia":  "París",
		"Italia":   "Roma",
		"Alemania": "Berlín",
		"Portugal": "Lisboa",
	}
	fmt.Printf("Map de capitales: %v\n", capitales)

	// Map con diferentes tipos
	info := map[string]interface{}{
		"nombre":  "Juan",
		"edad":    30,
		"activo":  true,
		"salario": 50000.50,
	}
	fmt.Printf("Map con interface{}: %v\n", info)

	// Map con llaves numéricas
	cuadrados := map[int]int{
		1: 1,
		2: 4,
		3: 9,
		4: 16,
		5: 25,
	}
	fmt.Printf("Map de cuadrados: %v\n", cuadrados)

	// Map declarado pero no inicializado (nil)
	var mapNil map[string]int
	fmt.Printf("Map nil: %v, es nil: %t\n", mapNil, mapNil == nil)

	// Inicializar map nil
	mapNil = make(map[string]int)
	mapNil["clave"] = 42
	fmt.Printf("Map después de inicializar: %v\n", mapNil)
}

func operacionesFundamentales() {
	personas := make(map[string]int)

	// Agregar elementos
	personas["Ana"] = 25
	personas["Bruno"] = 30
	personas["Carmen"] = 28
	fmt.Printf("Después de agregar: %v\n", personas)

	// Acceder a elementos
	edadAna := personas["Ana"]
	fmt.Printf("Edad de Ana: %d\n", edadAna)

	// Modificar elementos
	personas["Ana"] = 26
	fmt.Printf("Después de modificar edad de Ana: %v\n", personas)

	// Eliminar elementos
	delete(personas, "Bruno")
	fmt.Printf("Después de eliminar Bruno: %v\n", personas)

	// Longitud del map
	fmt.Printf("Número de personas: %d\n", len(personas))

	// Limpiar todo el map
	for k := range personas {
		delete(personas, k)
	}
	fmt.Printf("Map después de limpiar: %v\n", personas)

	// Verificar si está vacío
	fmt.Printf("¿Está vacío?: %t\n", len(personas) == 0)
}

func verificacionLlaves() {
	puntuaciones := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}

	// Verificación básica de existencia
	nombre := "Alice"
	if puntuacion, existe := puntuaciones[nombre]; existe {
		fmt.Printf("%s tiene puntuación: %d\n", nombre, puntuacion)
	} else {
		fmt.Printf("%s no encontrado\n", nombre)
	}

	// Verificación con clave que no existe
	nombre = "David"
	if puntuacion, existe := puntuaciones[nombre]; existe {
		fmt.Printf("%s tiene puntuación: %d\n", nombre, puntuacion)
	} else {
		fmt.Printf("%s no encontrado\n", nombre)
	}

	// Acceso sin verificación (retorna valor cero)
	puntuacionDavid := puntuaciones["David"]
	fmt.Printf("Puntuación de David (sin verificar): %d\n", puntuacionDavid)

	// Función auxiliar para verificar existencia
	nombres := []string{"Alice", "Bob", "Eve", "Charlie"}
	for _, nombre := range nombres {
		if existe := claveExiste(puntuaciones, nombre); existe {
			fmt.Printf("✓ %s está en el map\n", nombre)
		} else {
			fmt.Printf("✗ %s NO está en el map\n", nombre)
		}
	}

	// Obtener valor con valor por defecto
	valorPorDefecto := 0
	puntuacionEve := obtenerConDefault(puntuaciones, "Eve", valorPorDefecto)
	fmt.Printf("Puntuación de Eve (con default): %d\n", puntuacionEve)
}

func iteracionMaps() {
	notas := map[string]float64{
		"Matemáticas": 8.5,
		"Física":      9.0,
		"Química":     7.8,
		"Biología":    8.2,
		"Historia":    9.5,
	}

	// Iteración básica (clave y valor)
	fmt.Println("Iteración básica:")
	for materia, nota := range notas {
		fmt.Printf("  %s: %.1f\n", materia, nota)
	}

	// Solo claves
	fmt.Println("\nSolo claves:")
	for materia := range notas {
		fmt.Printf("  %s\n", materia)
	}

	// Solo valores (usando blank identifier)
	fmt.Println("\nSolo valores:")
	for _, nota := range notas {
		fmt.Printf("  %.1f\n", nota)
	}

	// Iteración con índice manual
	fmt.Println("\nCon índice manual:")
	i := 0
	for materia, nota := range notas {
		fmt.Printf("  %d: %s = %.1f\n", i, materia, nota)
		i++
	}

	// Recopilar claves y valores en slices
	var materias []string
	var calificaciones []float64

	for materia, nota := range notas {
		materias = append(materias, materia)
		calificaciones = append(calificaciones, nota)
	}

	fmt.Printf("\nMaterias: %v\n", materias)
	fmt.Printf("Calificaciones: %v\n", calificaciones)

	// Iteración con condiciones
	fmt.Println("\nMaterias con nota >= 9.0:")
	for materia, nota := range notas {
		if nota >= 9.0 {
			fmt.Printf("  %s: %.1f ⭐\n", materia, nota)
		}
	}
}

func mapsAnidados() {
	// Map de maps para representar información por regiones
	ventas := map[string]map[string]int{
		"Norte": {
			"Enero":   1000,
			"Febrero": 1200,
			"Marzo":   1100,
		},
		"Sur": {
			"Enero":   800,
			"Febrero": 900,
			"Marzo":   950,
		},
		"Este": {
			"Enero":   1300,
			"Febrero": 1400,
			"Marzo":   1350,
		},
	}

	fmt.Printf("Ventas por región: %v\n", ventas)

	// Acceder a datos anidados
	ventasNorteEnero := ventas["Norte"]["Enero"]
	fmt.Printf("Ventas Norte en Enero: %d\n", ventasNorteEnero)

	// Agregar nueva región
	ventas["Oeste"] = map[string]int{
		"Enero":   700,
		"Febrero": 750,
		"Marzo":   800,
	}

	// Agregar nuevo mes a región existente
	ventas["Norte"]["Abril"] = 1250

	// Iteración anidada
	fmt.Println("\nReporte completo de ventas:")
	for region, meses := range ventas {
		fmt.Printf("Región %s:\n", region)
		for mes, venta := range meses {
			fmt.Printf("  %s: %d\n", mes, venta)
		}
		fmt.Println()
	}

	// Calcular totales por región
	fmt.Println("Totales por región:")
	for region, meses := range ventas {
		total := 0
		for _, venta := range meses {
			total += venta
		}
		fmt.Printf("  %s: %d\n", region, total)
	}

	// Map tridimensional (región -> año -> mes -> ventas)
	ventasTriples := map[string]map[int]map[string]int{
		"Norte": {
			2023: {
				"Enero":   1000,
				"Febrero": 1100,
			},
			2024: {
				"Enero":   1200,
				"Febrero": 1300,
			},
		},
	}

	fmt.Printf("\nVentas Norte 2024 Enero: %d\n", ventasTriples["Norte"][2024]["Enero"])
}

func tiposDeDatos() {
	// Map con diferentes tipos de claves
	fmt.Println("Maps con diferentes tipos de claves:")

	// Claves int
	fibonacci := map[int]int{
		0: 0, 1: 1, 2: 1, 3: 2, 4: 3, 5: 5, 6: 8, 7: 13,
	}
	fmt.Printf("Fibonacci: %v\n", fibonacci)

	// Claves float64
	temperaturas := map[float64]string{
		0.0:     "Congelación",
		100.0:   "Ebullición",
		37.0:    "Corporal",
		-273.15: "Cero absoluto",
	}
	fmt.Printf("Temperaturas: %v\n", temperaturas)

	// Claves bool
	estados := map[bool]string{
		true:  "Activo",
		false: "Inactivo",
	}
	fmt.Printf("Estados: %v\n", estados)

	// Map con valores de diferentes tipos usando interface{}
	mixto := map[string]interface{}{
		"entero":   42,
		"flotante": 3.14,
		"cadena":   "texto",
		"booleano": true,
		"slice":    []int{1, 2, 3},
		"map":      map[string]int{"a": 1, "b": 2},
	}

	fmt.Println("\nMap mixto:")
	for clave, valor := range mixto {
		fmt.Printf("  %s: %v (tipo: %T)\n", clave, valor, valor)
	}

	// Type assertions para extraer valores
	if entero, ok := mixto["entero"].(int); ok {
		fmt.Printf("Entero extraído: %d\n", entero)
	}

	if slice, ok := mixto["slice"].([]int); ok {
		fmt.Printf("Slice extraído: %v\n", slice)
	}

	// Map de funciones
	operaciones := map[string]func(int, int) int{
		"suma":        func(a, b int) int { return a + b },
		"resta":       func(a, b int) int { return a - b },
		"multiplicar": func(a, b int) int { return a * b },
		"dividir":     func(a, b int) int { return a / b },
	}

	fmt.Println("\nMap de funciones:")
	a, b := 10, 3
	for nombre, operacion := range operaciones {
		if nombre == "dividir" && b == 0 {
			continue // Evitar división por cero
		}
		resultado := operacion(a, b)
		fmt.Printf("  %s(%d, %d) = %d\n", nombre, a, b, resultado)
	}
}

func mapsConStructs() {
	// Definir struct local
	type Persona struct {
		Nombre  string
		Edad    int
		Email   string
		Salario float64
		Activo  bool
	}

	// Map con structs como valores
	empleados := map[int]Persona{
		1001: {"Ana García", 28, "ana@empresa.com", 50000, true},
		1002: {"Bruno López", 32, "bruno@empresa.com", 60000, true},
		1003: {"Carmen Ruiz", 29, "carmen@empresa.com", 55000, false},
		1004: {"David Chen", 35, "david@empresa.com", 70000, true},
	}

	fmt.Println("Empleados:")
	for id, empleado := range empleados {
		estado := "Activo"
		if !empleado.Activo {
			estado = "Inactivo"
		}
		fmt.Printf("  ID %d: %s (%d años) - %s - $%.2f\n",
			id, empleado.Nombre, empleado.Edad, estado, empleado.Salario)
	}

	// Modificar struct en map
	empleado1001 := empleados[1001]
	empleado1001.Salario = 52000
	empleados[1001] = empleado1001 // Necesario reasignar
	fmt.Printf("\nSalario actualizado de Ana: $%.2f\n", empleados[1001].Salario)

	// Map con punteros a structs (más eficiente para modificaciones)
	empleadosPtr := map[int]*Persona{
		2001: {"Elena Vega", 26, "elena@empresa.com", 48000, true},
		2002: {"Fernando Soto", 31, "fernando@empresa.com", 58000, true},
	}

	// Modificar a través del puntero
	empleadosPtr[2001].Salario = 50000
	fmt.Printf("Salario actualizado de Elena: $%.2f\n", empleadosPtr[2001].Salario)

	// Struct como clave (debe ser comparable)
	type Coordenada struct {
		X, Y int
	}

	ciudades := map[Coordenada]string{
		{0, 0}:    "Origen",
		{10, 20}:  "Ciudad A",
		{-5, 15}:  "Ciudad B",
		{30, -10}: "Ciudad C",
	}

	fmt.Println("\nCiudades por coordenadas:")
	for coord, ciudad := range ciudades {
		fmt.Printf("  (%d, %d): %s\n", coord.X, coord.Y, ciudad)
	}

	// Buscar ciudad por coordenada
	buscar := Coordenada{10, 20}
	if ciudad, existe := ciudades[buscar]; existe {
		fmt.Printf("En coordenada %v se encuentra: %s\n", buscar, ciudad)
	}

	// Funciones auxiliares para trabajar con empleados
	fmt.Println("\nEmpleados activos:")
	for id, empleado := range empleados {
		if empleado.Activo {
			fmt.Printf("  ID %d: %s\n", id, empleado.Nombre)
		}
	}

	// Calcular promedio de salario
	total := 0.0
	count := 0
	for _, empleado := range empleados {
		total += empleado.Salario
		count++
	}
	promedioSalario := total / float64(count)
	fmt.Printf("\nSalario promedio: $%.2f\n", promedioSalario)
}

func mapsComoSets() {
	// Simular set usando map[tipo]bool
	fmt.Println("Usando maps como sets:")

	// Set de strings
	frutas := map[string]bool{
		"manzana": true,
		"banana":  true,
		"naranja": true,
	}

	// Agregar elemento
	frutas["kiwi"] = true
	fmt.Printf("Set de frutas: %v\n", obtenerClaves(frutas))

	// Verificar pertenencia
	if frutas["manzana"] {
		fmt.Println("manzana está en el set")
	}

	// Eliminar elemento
	delete(frutas, "banana")
	fmt.Printf("Después de eliminar banana: %v\n", obtenerClaves(frutas))

	// Set usando map[tipo]struct{} (más eficiente en memoria)
	numeros := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		5: {},
		8: {},
	}

	fmt.Printf("Set de números: %v\n", obtenerClavesInt(numeros))

	// Operaciones de conjuntos
	set1 := map[string]bool{
		"a": true, "b": true, "c": true,
	}
	set2 := map[string]bool{
		"b": true, "c": true, "d": true,
	}

	fmt.Printf("Set1: %v\n", obtenerClaves(set1))
	fmt.Printf("Set2: %v\n", obtenerClaves(set2))

	// Unión
	union := unionSets(set1, set2)
	fmt.Printf("Unión: %v\n", obtenerClaves(union))

	// Intersección
	interseccion := interseccionSets(set1, set2)
	fmt.Printf("Intersección: %v\n", obtenerClaves(interseccion))

	// Diferencia
	diferencia := diferenciaSets(set1, set2)
	fmt.Printf("Diferencia (set1 - set2): %v\n", obtenerClaves(diferencia))

	// Set de elementos únicos de un slice
	elementos := []string{"a", "b", "c", "b", "d", "a", "e", "c"}
	unicos := crearSetDeSlice(elementos)
	fmt.Printf("Elementos originales: %v\n", elementos)
	fmt.Printf("Elementos únicos: %v\n", obtenerClaves(unicos))
}

func contadoresYFrecuencias() {
	// Contar frecuencia de palabras
	texto := "el gato subió al tejado el gato bajó del tejado"
	palabras := strings.Fields(texto)

	frecuencias := make(map[string]int)
	for _, palabra := range palabras {
		frecuencias[palabra]++
	}

	fmt.Printf("Texto: \"%s\"\n", texto)
	fmt.Println("Frecuencia de palabras:")
	for palabra, count := range frecuencias {
		fmt.Printf("  %s: %d\n", palabra, count)
	}

	// Contar frecuencia de caracteres
	frase := "programación"
	frecuenciaChars := make(map[rune]int)
	for _, char := range frase {
		frecuenciaChars[char]++
	}

	fmt.Printf("\nFrase: \"%s\"\n", frase)
	fmt.Println("Frecuencia de caracteres:")
	for char, count := range frecuenciaChars {
		fmt.Printf("  '%c': %d\n", char, count)
	}

	// Contador de ocurrencias en slice
	numeros := []int{1, 2, 3, 2, 1, 4, 2, 5, 1, 3, 2}
	conteoNumeros := make(map[int]int)
	for _, num := range numeros {
		conteoNumeros[num]++
	}

	fmt.Printf("\nNúmeros: %v\n", numeros)
	fmt.Println("Conteo de números:")
	for num, count := range conteoNumeros {
		fmt.Printf("  %d: %d veces\n", num, count)
	}

	// Encontrar el elemento más frecuente
	maxCount := 0
	elementoMasFrecuente := 0
	for num, count := range conteoNumeros {
		if count > maxCount {
			maxCount = count
			elementoMasFrecuente = num
		}
	}
	fmt.Printf("Elemento más frecuente: %d (aparece %d veces)\n", elementoMasFrecuente, maxCount)

	// Histograma de calificaciones
	calificaciones := []int{85, 92, 78, 96, 87, 91, 83, 89, 94, 86, 90, 88}
	rangos := map[string]int{
		"90-100": 0,
		"80-89":  0,
		"70-79":  0,
		"60-69":  0,
		"<60":    0,
	}

	for _, calif := range calificaciones {
		switch {
		case calif >= 90:
			rangos["90-100"]++
		case calif >= 80:
			rangos["80-89"]++
		case calif >= 70:
			rangos["70-79"]++
		case calif >= 60:
			rangos["60-69"]++
		default:
			rangos["<60"]++
		}
	}

	fmt.Printf("\nCalificaciones: %v\n", calificaciones)
	fmt.Println("Histograma por rangos:")
	for rango, count := range rangos {
		fmt.Printf("  %s: %d estudiantes\n", rango, count)
	}
}

func agrupacionDatos() {
	// Agrupar estudiantes por grado
	type Estudiante struct {
		Nombre string
		Grado  int
		Nota   float64
	}

	estudiantes := []Estudiante{
		{"Ana", 10, 8.5},
		{"Bruno", 11, 7.8},
		{"Carmen", 10, 9.2},
		{"David", 11, 8.1},
		{"Elena", 10, 9.0},
		{"Fernando", 12, 8.8},
		{"Gabriela", 11, 9.5},
		{"Hugo", 12, 7.9},
	}

	// Agrupar por grado
	porGrado := make(map[int][]Estudiante)
	for _, estudiante := range estudiantes {
		porGrado[estudiante.Grado] = append(porGrado[estudiante.Grado], estudiante)
	}

	fmt.Println("Estudiantes agrupados por grado:")
	for grado, lista := range porGrado {
		fmt.Printf("Grado %d:\n", grado)
		for _, estudiante := range lista {
			fmt.Printf("  %s (%.1f)\n", estudiante.Nombre, estudiante.Nota)
		}
		fmt.Println()
	}

	// Agrupar por rango de notas
	porRango := map[string][]Estudiante{
		"Excelente (9.0+)":       {},
		"Bueno (8.0-8.9)":        {},
		"Regular (7.0-7.9)":      {},
		"Necesita mejora (<7.0)": {},
	}

	for _, estudiante := range estudiantes {
		switch {
		case estudiante.Nota >= 9.0:
			porRango["Excelente (9.0+)"] = append(porRango["Excelente (9.0+)"], estudiante)
		case estudiante.Nota >= 8.0:
			porRango["Bueno (8.0-8.9)"] = append(porRango["Bueno (8.0-8.9)"], estudiante)
		case estudiante.Nota >= 7.0:
			porRango["Regular (7.0-7.9)"] = append(porRango["Regular (7.0-7.9)"], estudiante)
		default:
			porRango["Necesita mejora (<7.0)"] = append(porRango["Necesita mejora (<7.0)"], estudiante)
		}
	}

	fmt.Println("Estudiantes agrupados por rendimiento:")
	for rango, lista := range porRango {
		if len(lista) > 0 {
			fmt.Printf("%s:\n", rango)
			for _, estudiante := range lista {
				fmt.Printf("  %s (Grado %d, %.1f)\n", estudiante.Nombre, estudiante.Grado, estudiante.Nota)
			}
			fmt.Println()
		}
	}

	// Índice inverso: agrupar por primera letra del nombre
	porLetra := make(map[rune][]string)
	for _, estudiante := range estudiantes {
		primeraLetra := rune(estudiante.Nombre[0])
		porLetra[primeraLetra] = append(porLetra[primeraLetra], estudiante.Nombre)
	}

	fmt.Println("Estudiantes agrupados por primera letra:")
	for letra, nombres := range porLetra {
		fmt.Printf("Letra '%c': %v\n", letra, nombres)
	}
}

func mapsOrdenados() {
	// Los maps no mantienen orden, pero podemos ordenar las claves
	poblaciones := map[string]int{
		"Madrid":    3200000,
		"Barcelona": 1600000,
		"Valencia":  790000,
		"Sevilla":   690000,
		"Zaragoza":  670000,
		"Málaga":    570000,
		"Murcia":    450000,
		"Palma":     410000,
	}

	fmt.Println("Map original (orden aleatorio):")
	for ciudad, poblacion := range poblaciones {
		fmt.Printf("  %s: %d\n", ciudad, poblacion)
	}

	// Ordenar por claves (alfabéticamente)
	var ciudades []string
	for ciudad := range poblaciones {
		ciudades = append(ciudades, ciudad)
	}
	sort.Strings(ciudades)

	fmt.Println("\nOrdenado por ciudad (alfabéticamente):")
	for _, ciudad := range ciudades {
		fmt.Printf("  %s: %d\n", ciudad, poblaciones[ciudad])
	}

	// Ordenar por valores (población)
	type CiudadPoblacion struct {
		Ciudad    string
		Poblacion int
	}

	var ciudadesPoblacion []CiudadPoblacion
	for ciudad, poblacion := range poblaciones {
		ciudadesPoblacion = append(ciudadesPoblacion, CiudadPoblacion{ciudad, poblacion})
	}

	// Ordenar por población (descendente)
	sort.Slice(ciudadesPoblacion, func(i, j int) bool {
		return ciudadesPoblacion[i].Poblacion > ciudadesPoblacion[j].Poblacion
	})

	fmt.Println("\nOrdenado por población (mayor a menor):")
	for i, cp := range ciudadesPoblacion {
		fmt.Printf("  %d. %s: %d\n", i+1, cp.Ciudad, cp.Poblacion)
	}

	// Top 3 ciudades más pobladas
	fmt.Println("\nTop 3 ciudades más pobladas:")
	for i := 0; i < 3 && i < len(ciudadesPoblacion); i++ {
		cp := ciudadesPoblacion[i]
		fmt.Printf("  %d. %s: %d\n", i+1, cp.Ciudad, cp.Poblacion)
	}

	// Filtrar ciudades con más de 600,000 habitantes
	fmt.Println("\nCiudades con más de 600,000 habitantes:")
	for _, cp := range ciudadesPoblacion {
		if cp.Poblacion > 600000 {
			fmt.Printf("  %s: %d\n", cp.Ciudad, cp.Poblacion)
		}
	}
}

func operacionesAvanzadas() {
	// Merge de maps
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	map2 := map[string]int{"c": 30, "d": 4, "e": 5}

	fmt.Printf("Map1: %v\n", map1)
	fmt.Printf("Map2: %v\n", map2)

	merged := mergeMaps(map1, map2)
	fmt.Printf("Merged: %v\n", merged)

	// Invertir map (valores como claves)
	colores := map[string]string{
		"rojo":     "#FF0000",
		"verde":    "#00FF00",
		"azul":     "#0000FF",
		"amarillo": "#FFFF00",
	}

	invertido := invertirMap(colores)
	fmt.Printf("\nColores originales: %v\n", colores)
	fmt.Printf("Map invertido: %v\n", invertido)

	// Filtrar map
	numeros := map[string]int{
		"uno": 1, "dos": 2, "tres": 3, "cuatro": 4, "cinco": 5,
		"seis": 6, "siete": 7, "ocho": 8, "nueve": 9, "diez": 10,
	}

	pares := filtrarMap(numeros, func(k string, v int) bool {
		return v%2 == 0
	})

	fmt.Printf("\nNúmeros originales: %v\n", numeros)
	fmt.Printf("Solo pares: %v\n", pares)

	// Transformar valores del map
	cuadrados := transformarValores(numeros, func(v int) int {
		return v * v
	})

	fmt.Printf("Cuadrados: %v\n", cuadrados)

	// Verificar si dos maps son iguales
	mapA := map[string]int{"x": 1, "y": 2}
	mapB := map[string]int{"y": 2, "x": 1}
	mapC := map[string]int{"x": 1, "y": 3}

	fmt.Printf("\nMapA: %v\n", mapA)
	fmt.Printf("MapB: %v\n", mapB)
	fmt.Printf("MapC: %v\n", mapC)
	fmt.Printf("¿MapA == MapB?: %t\n", mapsIguales(mapA, mapB))
	fmt.Printf("¿MapA == MapC?: %t\n", mapsIguales(mapA, mapC))

	// Clonar map
	original := map[string]int{"a": 1, "b": 2, "c": 3}
	clon := clonarMap(original)

	fmt.Printf("\nOriginal: %v\n", original)
	fmt.Printf("Clon: %v\n", clon)

	// Modificar clon no afecta original
	clon["d"] = 4
	clon["a"] = 10

	fmt.Printf("Después de modificar clon:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Clon: %v\n", clon)
}

func casosDeUsoPracticos() {
	// 1. Cache/Memoización
	fmt.Println("1. Cache/Memoización:")
	cache := make(map[int]int)

	var fibonacci func(int) int
	fibonacci = func(n int) int {
		if val, existe := cache[n]; existe {
			fmt.Printf("  Cache hit para fibonacci(%d)\n", n)
			return val
		}

		var result int
		if n <= 1 {
			result = n
		} else {
			result = fibonacci(n-1) + fibonacci(n-2)
		}

		cache[n] = result
		fmt.Printf("  Calculado fibonacci(%d) = %d\n", n, result)
		return result
	}

	fmt.Printf("fibonacci(10) = %d\n", fibonacci(10))
	fmt.Printf("fibonacci(8) = %d\n", fibonacci(8)) // Debería usar cache

	// 2. Configuración de aplicación
	fmt.Println("\n2. Configuración de aplicación:")
	config := map[string]interface{}{
		"database_host":   "localhost",
		"database_port":   5432,
		"max_connections": 100,
		"enable_logging":  true,
		"log_level":       "INFO",
		"timeout_seconds": 30,
		"allowed_origins": []string{"localhost", "127.0.0.1"},
	}

	// Función para obtener configuración con tipo específico
	getConfig := func(key string, defaultValue interface{}) interface{} {
		if val, existe := config[key]; existe {
			return val
		}
		return defaultValue
	}

	dbHost := getConfig("database_host", "127.0.0.1").(string)
	dbPort := getConfig("database_port", 3306).(int)
	enableLogging := getConfig("enable_logging", false).(bool)

	fmt.Printf("DB Host: %s\n", dbHost)
	fmt.Printf("DB Port: %d\n", dbPort)
	fmt.Printf("Logging: %t\n", enableLogging)

	// 3. Router HTTP simple
	fmt.Println("\n3. Router HTTP simple:")
	type Handler func(string) string

	routes := map[string]Handler{
		"GET /users":    func(path string) string { return "Lista de usuarios" },
		"POST /users":   func(path string) string { return "Crear usuario" },
		"GET /products": func(path string) string { return "Lista de productos" },
		"DELETE /users": func(path string) string { return "Eliminar usuario" },
	}

	// Simular peticiones
	requests := []string{"GET /users", "POST /users", "GET /orders", "DELETE /users"}

	for _, request := range requests {
		if handler, existe := routes[request]; existe {
			response := handler(request)
			fmt.Printf("  %s -> %s\n", request, response)
		} else {
			fmt.Printf("  %s -> 404 Not Found\n", request)
		}
	}

	// 4. Sistema de permisos
	fmt.Println("\n4. Sistema de permisos:")
	permisos := map[string]map[string]bool{
		"admin": {
			"read":   true,
			"write":  true,
			"delete": true,
		},
		"editor": {
			"read":   true,
			"write":  true,
			"delete": false,
		},
		"viewer": {
			"read":   true,
			"write":  false,
			"delete": false,
		},
	}

	verificarPermiso := func(rol, accion string) bool {
		if rolPermisos, existe := permisos[rol]; existe {
			return rolPermisos[accion]
		}
		return false
	}

	usuarios := []string{"admin", "editor", "viewer", "guest"}
	acciones := []string{"read", "write", "delete"}

	fmt.Println("  Matriz de permisos:")
	fmt.Printf("  %-10s", "Rol\\Acción")
	for _, accion := range acciones {
		fmt.Printf("%-8s", accion)
	}
	fmt.Println()

	for _, usuario := range usuarios {
		fmt.Printf("  %-10s", usuario)
		for _, accion := range acciones {
			if verificarPermiso(usuario, accion) {
				fmt.Printf("%-8s", "✓")
			} else {
				fmt.Printf("%-8s", "✗")
			}
		}
		fmt.Println()
	}
}

func mejoresPracticas() {
	// 1. Inicialización con capacidad conocida
	fmt.Println("1. Pre-allocar capacidad cuando sea posible:")

	// Si conocemos el tamaño aproximado, mejor usar make con capacidad
	elementos := 1000
	mapGrande := make(map[int]string, elementos)
	for i := 0; i < elementos; i++ {
		mapGrande[i] = fmt.Sprintf("valor_%d", i)
	}
	fmt.Printf("Map con %d elementos creado eficientemente\n", len(mapGrande))

	// 2. Verificar existencia antes de usar
	fmt.Println("\n2. Siempre verificar existencia de claves:")

	datos := map[string]int{"a": 1, "b": 2}

	// Malo: no verificar existencia
	//valor := datos["c"] // Retorna 0, puede ser confuso

	// Bueno: verificar existencia
	if valor, existe := datos["c"]; existe {
		fmt.Printf("Valor de 'c': %d\n", valor)
	} else {
		fmt.Println("Clave 'c' no existe")
	}

	// 3. Usar zero value vs nil maps
	fmt.Println("\n3. Diferencia entre nil map y map vacío:")

	var mapNil map[string]int
	mapVacio := make(map[string]int)

	fmt.Printf("Map nil: %v, es nil: %t\n", mapNil, mapNil == nil)
	fmt.Printf("Map vacío: %v, es nil: %t\n", mapVacio, mapVacio == nil)

	// mapNil["key"] = 1 // Esto causaría panic!
	mapVacio["key"] = 1 // Esto es seguro

	// 4. Maps no son thread-safe
	fmt.Println("\n4. Maps no son thread-safe (usar sync.Map para concurrencia)")
	fmt.Println("Para acceso concurrente, considerar sync.Map o mutex")

	// 5. Memory leaks con maps grandes
	fmt.Println("\n5. Cuidado con memory leaks:")
	fmt.Println("- Los maps no reducen su memoria automáticamente")
	fmt.Println("- Si un map crece mucho y luego se reduce, la memoria no se libera")
	fmt.Println("- Considerar recrear el map si se reduce significativamente")

	// Ejemplo: recrear map después de eliminar muchos elementos
	mapGrande2 := make(map[int]string)
	for i := 0; i < 10000; i++ {
		mapGrande2[i] = fmt.Sprintf("valor_%d", i)
	}

	// Eliminar muchos elementos
	for i := 0; i < 9000; i++ {
		delete(mapGrande2, i)
	}

	// Si el map se reduce mucho, mejor recrearlo
	if len(mapGrande2) < 1000 { // Threshold arbitrario
		nuevoMap := make(map[int]string)
		for k, v := range mapGrande2 {
			nuevoMap[k] = v
		}
		mapGrande2 = nuevoMap
	}

	fmt.Printf("Map reducido tiene %d elementos\n", len(mapGrande2))

	// 6. Comparar maps
	fmt.Println("\n6. Maps no son comparables directamente:")
	fmt.Println("- Usar funciones auxiliares para comparar")
	fmt.Println("- Solo se puede comparar con nil")

	// 7. Iteración no es determinista
	fmt.Println("\n7. El orden de iteración no es determinista:")
	fmt.Println("- Si necesitas orden específico, usa slices ordenados para las claves")

	testMap := map[string]int{"c": 3, "a": 1, "b": 2}
	fmt.Print("Iteración 1: ")
	for k := range testMap {
		fmt.Printf("%s ", k)
	}
	fmt.Println()

	fmt.Print("Iteración 2: ")
	for k := range testMap {
		fmt.Printf("%s ", k)
	}
	fmt.Println()
	fmt.Println("(El orden puede variar entre ejecuciones)")
}

// Funciones auxiliares
func claveExiste(m map[string]int, clave string) bool {
	_, existe := m[clave]
	return existe
}

func obtenerConDefault(m map[string]int, clave string, defaultValue int) int {
	if valor, existe := m[clave]; existe {
		return valor
	}
	return defaultValue
}

func obtenerClaves(m map[string]bool) []string {
	claves := make([]string, 0, len(m))
	for k := range m {
		claves = append(claves, k)
	}
	sort.Strings(claves)
	return claves
}

func obtenerClavesInt(m map[int]struct{}) []int {
	claves := make([]int, 0, len(m))
	for k := range m {
		claves = append(claves, k)
	}
	sort.Ints(claves)
	return claves
}

func unionSets(set1, set2 map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for k := range set1 {
		result[k] = true
	}
	for k := range set2 {
		result[k] = true
	}
	return result
}

func interseccionSets(set1, set2 map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for k := range set1 {
		if set2[k] {
			result[k] = true
		}
	}
	return result
}

func diferenciaSets(set1, set2 map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for k := range set1 {
		if !set2[k] {
			result[k] = true
		}
	}
	return result
}

func crearSetDeSlice(slice []string) map[string]bool {
	set := make(map[string]bool)
	for _, item := range slice {
		set[item] = true
	}
	return set
}

func mergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range map1 {
		result[k] = v
	}
	for k, v := range map2 {
		result[k] = v // Los valores de map2 sobrescriben map1
	}
	return result
}

func invertirMap(m map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		result[v] = k
	}
	return result
}

func filtrarMap(m map[string]int, predicado func(string, int) bool) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		if predicado(k, v) {
			result[k] = v
		}
	}
	return result
}

func transformarValores(m map[string]int, transform func(int) int) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		result[k] = transform(v)
	}
	return result
}

func mapsIguales(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v1 := range map1 {
		if v2, existe := map2[k]; !existe || v1 != v2 {
			return false
		}
	}
	return true
}

func clonarMap(original map[string]int) map[string]int {
	clon := make(map[string]int)
	for k, v := range original {
		clon[k] = v
	}
	return clon
}
