package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// ============= INTERFACES VACÍAS Y TYPE ASSERTIONS =============

// Función que acepta cualquier tipo usando interface{}
func describir(valor interface{}) {
	fmt.Printf("    Valor: %v, Tipo: %T\n", valor, valor)
}

// Función que procesa diferentes tipos usando type assertion
func procesarValor(valor interface{}) {
	switch v := valor.(type) {
	case int:
		fmt.Printf("    Entero: %d (doble: %d)\n", v, v*2)
	case float64:
		fmt.Printf("    Float: %.2f (cuadrado: %.2f)\n", v, v*v)
	case string:
		fmt.Printf("    String: '%s' (longitud: %d)\n", v, len(v))
	case bool:
		fmt.Printf("    Boolean: %t (negado: %t)\n", v, !v)
	case []int:
		suma := 0
		for _, num := range v {
			suma += num
		}
		fmt.Printf("    Slice de enteros: %v (suma: %d)\n", v, suma)
	case map[string]int:
		fmt.Printf("    Map string->int: %v (claves: %d)\n", v, len(v))
	case *Punto:
		fmt.Printf("    Puntero a Punto: %s (distancia al origen: %.2f)\n",
			v, v.Distancia(Punto{0, 0}))
	case Punto:
		fmt.Printf("    Punto: %s (distancia al origen: %.2f)\n",
			v, v.Distancia(Punto{0, 0}))
	default:
		fmt.Printf("    Tipo desconocido: %T con valor %v\n", v, v)
	}
}

// Función que intenta convertir tipos
func convertirAString(valor interface{}) (string, bool) {
	switch v := valor.(type) {
	case string:
		return v, true
	case int:
		return strconv.Itoa(v), true
	case float64:
		return strconv.FormatFloat(v, 'f', 2, 64), true
	case bool:
		return strconv.FormatBool(v), true
	default:
		return "", false
	}
}

// Type assertion más segura
func extraerEntero(valor interface{}) (int, error) {
	if entero, ok := valor.(int); ok {
		return entero, nil
	}
	return 0, fmt.Errorf("el valor %v no es un entero", valor)
}

// Función que suma cualquier tipo numérico
func sumarNumericos(valores ...interface{}) (float64, error) {
	var suma float64

	for i, valor := range valores {
		switch v := valor.(type) {
		case int:
			suma += float64(v)
		case int32:
			suma += float64(v)
		case int64:
			suma += float64(v)
		case float32:
			suma += float64(v)
		case float64:
			suma += v
		default:
			return 0, fmt.Errorf("el valor en posición %d (%v) no es numérico", i, valor)
		}
	}

	return suma, nil
}

// Estructura para almacenar diferentes tipos
type Contenedor struct {
	elementos []interface{}
}

func NuevoContenedor() *Contenedor {
	return &Contenedor{elementos: make([]interface{}, 0)}
}

func (c *Contenedor) Agregar(elemento interface{}) {
	c.elementos = append(c.elementos, elemento)
}

func (c *Contenedor) Obtener(indice int) (interface{}, bool) {
	if indice < 0 || indice >= len(c.elementos) {
		return nil, false
	}
	return c.elementos[indice], true
}

func (c *Contenedor) ObtenerComoInt(indice int) (int, error) {
	elemento, ok := c.Obtener(indice)
	if !ok {
		return 0, fmt.Errorf("índice %d fuera de rango", indice)
	}

	if entero, ok := elemento.(int); ok {
		return entero, nil
	}

	return 0, fmt.Errorf("el elemento en índice %d no es un entero", indice)
}

func (c *Contenedor) ObtenerComoString(indice int) (string, error) {
	elemento, ok := c.Obtener(indice)
	if !ok {
		return "", fmt.Errorf("índice %d fuera de rango", indice)
	}

	if cadena, ok := elemento.(string); ok {
		return cadena, nil
	}

	return "", fmt.Errorf("el elemento en índice %d no es una cadena", indice)
}

func (c *Contenedor) Tamaño() int {
	return len(c.elementos)
}

func (c *Contenedor) ListarTipos() {
	fmt.Println("    Tipos en el contenedor:")
	for i, elemento := range c.elementos {
		fmt.Printf("      [%d] %T: %v\n", i, elemento, elemento)
	}
}

func (c *Contenedor) FiltrarPorTipo(tipoDeseado reflect.Type) []interface{} {
	var filtrados []interface{}

	for _, elemento := range c.elementos {
		if reflect.TypeOf(elemento) == tipoDeseado {
			filtrados = append(filtrados, elemento)
		}
	}

	return filtrados
}

// Procesador genérico usando reflection
type ProcesadorGenerico struct {
	procesadores map[reflect.Type]func(interface{}) string
}

func NuevoProcesadorGenerico() *ProcesadorGenerico {
	pg := &ProcesadorGenerico{
		procesadores: make(map[reflect.Type]func(interface{}) string),
	}

	// Registrar procesadores por defecto
	pg.RegistrarProcesador(reflect.TypeOf(0), func(v interface{}) string {
		return fmt.Sprintf("Entero procesado: %d", v.(int))
	})

	pg.RegistrarProcesador(reflect.TypeOf(""), func(v interface{}) string {
		return fmt.Sprintf("String procesado: '%s' (mayúsculas)",
			fmt.Sprintf("%s", v.(string)))
	})

	pg.RegistrarProcesador(reflect.TypeOf(0.0), func(v interface{}) string {
		return fmt.Sprintf("Float procesado: %.2f", v.(float64))
	})

	return pg
}

func (pg *ProcesadorGenerico) RegistrarProcesador(tipo reflect.Type, procesador func(interface{}) string) {
	pg.procesadores[tipo] = procesador
}

func (pg *ProcesadorGenerico) Procesar(valor interface{}) string {
	tipo := reflect.TypeOf(valor)
	if procesador, existe := pg.procesadores[tipo]; existe {
		return procesador(valor)
	}

	return fmt.Sprintf("Sin procesador para tipo %T: %v", valor, valor)
}

// Función para demostrar type assertion con panic recovery
func typeAssertionSegura(valor interface{}, tipoEsperado string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("    ¡Panic recuperado! Error de type assertion: %v\n", r)
		}
	}()

	switch tipoEsperado {
	case "int":
		entero := valor.(int) // Puede causar panic si no es int
		fmt.Printf("    Type assertion exitosa: %d es un entero\n", entero)
	case "string":
		cadena := valor.(string) // Puede causar panic si no es string
		fmt.Printf("    Type assertion exitosa: '%s' es una cadena\n", cadena)
	case "[]int":
		slice := valor.([]int) // Puede causar panic si no es []int
		fmt.Printf("    Type assertion exitosa: %v es un slice de enteros\n", slice)
	default:
		fmt.Printf("    Tipo no soportado para assertion: %s\n", tipoEsperado)
	}
}

func ejemploInterfazVacia() {
	fmt.Println("  === INTERFACE{} BÁSICA ===")

	// Diferentes tipos usando interface{}
	valores := []interface{}{
		42,
		3.14159,
		"Hola, mundo",
		true,
		[]int{1, 2, 3, 4, 5},
		map[string]int{"a": 1, "b": 2},
		Punto{X: 3, Y: 4},
		&Punto{X: 5, Y: 12},
	}

	fmt.Println("    Describiendo diferentes tipos:")
	for i, valor := range valores {
		fmt.Printf("  [%d] ", i)
		describir(valor)
	}

	fmt.Println("\n  === TYPE SWITCH ===")
	fmt.Println("    Procesando valores con type switch:")
	for i, valor := range valores {
		fmt.Printf("  [%d] ", i)
		procesarValor(valor)
	}

	fmt.Println("\n  === CONVERSIONES SEGURAS ===")
	valoresParaConvertir := []interface{}{42, 3.14, "texto", true, []int{1, 2}}

	for i, valor := range valoresParaConvertir {
		if cadena, ok := convertirAString(valor); ok {
			fmt.Printf("    [%d] %T -> string: '%s'\n", i, valor, cadena)
		} else {
			fmt.Printf("    [%d] %T no se puede convertir a string\n", i, valor)
		}
	}

	fmt.Println("\n  === TYPE ASSERTIONS SEGURAS ===")
	valoresEnteros := []interface{}{42, "no es entero", 99, 3.14, 100}

	for i, valor := range valoresEnteros {
		if entero, err := extraerEntero(valor); err == nil {
			fmt.Printf("    [%d] Entero extraído: %d\n", i, entero)
		} else {
			fmt.Printf("    [%d] Error: %v\n", i, err)
		}
	}

	fmt.Println("\n  === SUMA DE TIPOS NUMÉRICOS ===")
	numericos := []interface{}{10, 3.5, int32(7), float32(2.2), int64(15)}

	if suma, err := sumarNumericos(numericos...); err == nil {
		fmt.Printf("    Suma de valores numéricos: %.2f\n", suma)
	} else {
		fmt.Printf("    Error en suma: %v\n", err)
	}

	// Intentar con tipos no numéricos
	mixtos := []interface{}{10, 3.5, "no numérico", 7}
	if _, err := sumarNumericos(mixtos...); err != nil {
		fmt.Printf("    ✓ Error esperado con tipos mixtos: %v\n", err)
	}

	fmt.Println("\n  === CONTENEDOR GENÉRICO ===")
	contenedor := NuevoContenedor()

	// Agregar diferentes tipos
	contenedor.Agregar(100)
	contenedor.Agregar("texto ejemplo")
	contenedor.Agregar(3.14159)
	contenedor.Agregar([]int{10, 20, 30})
	contenedor.Agregar(true)

	fmt.Printf("    Contenedor con %d elementos:\n", contenedor.Tamaño())
	contenedor.ListarTipos()

	// Extraer valores específicos
	if entero, err := contenedor.ObtenerComoInt(0); err == nil {
		fmt.Printf("    Entero en posición 0: %d\n", entero)
	} else {
		fmt.Printf("    Error: %v\n", err)
	}

	if cadena, err := contenedor.ObtenerComoString(1); err == nil {
		fmt.Printf("    String en posición 1: '%s'\n", cadena)
	} else {
		fmt.Printf("    Error: %v\n", err)
	}

	// Intentar extraer tipo incorrecto
	if _, err := contenedor.ObtenerComoInt(1); err != nil {
		fmt.Printf("    ✓ Error esperado al extraer string como int: %v\n", err)
	}

	fmt.Println("\n  === FILTRADO POR TIPO ===")
	// Agregar más elementos para filtrar
	contenedor.Agregar(200)
	contenedor.Agregar("otro texto")
	contenedor.Agregar(300)

	enteros := contenedor.FiltrarPorTipo(reflect.TypeOf(0))
	fmt.Printf("    Enteros encontrados: %v\n", enteros)

	strings := contenedor.FiltrarPorTipo(reflect.TypeOf(""))
	fmt.Printf("    Strings encontrados: %v\n", strings)

	fmt.Println("\n  === PROCESADOR GENÉRICO ===")
	procesador := NuevoProcesadorGenerico()

	valoresProcesar := []interface{}{42, "hello world", 3.14159, []int{1, 2, 3}}

	for i, valor := range valoresProcesar {
		resultado := procesador.Procesar(valor)
		fmt.Printf("    [%d] %s\n", i, resultado)
	}

	// Registrar procesador personalizado para slice de enteros
	procesador.RegistrarProcesador(reflect.TypeOf([]int{}), func(v interface{}) string {
		slice := v.([]int)
		suma := 0
		for _, num := range slice {
			suma += num
		}
		return fmt.Sprintf("Slice de enteros procesado: %v (suma: %d)", slice, suma)
	})

	// Procesar de nuevo con el nuevo procesador
	resultado := procesador.Procesar([]int{1, 2, 3})
	fmt.Printf("    Con procesador personalizado: %s\n", resultado)

	fmt.Println("\n  === TYPE ASSERTIONS CON PANIC RECOVERY ===")
	valoresAssert := []interface{}{42, "texto", []int{1, 2, 3}}
	tiposEsperados := []string{"int", "int", "[]int", "string"}

	for i, valor := range valoresAssert {
		if i < len(tiposEsperados) {
			fmt.Printf("    Intentando assertion de %T como %s:\n", valor, tiposEsperados[i])
			typeAssertionSegura(valor, tiposEsperados[i])
		}
	}

	// Demostrar assertion que causará panic
	fmt.Println("    Intentando assertion incorrecta (causará panic):")
	typeAssertionSegura("esto es string", "int")
}
