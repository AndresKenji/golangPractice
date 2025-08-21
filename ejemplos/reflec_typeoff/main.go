package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Estructura de ejemplo para demostrar reflection
type Persona struct {
	Nombre  string `json:"nombre" db:"name"`
	Edad    int    `json:"edad" db:"age"`
	Email   string `json:"email" db:"email"`
	Activo  bool   `json:"activo" db:"active"`
	privado string // campo privado
}

// Método de la estructura
func (p Persona) Saludar() string {
	return fmt.Sprintf("Hola, soy %s y tengo %d años", p.Nombre, p.Edad)
}

// Método con puntero
func (p *Persona) CambiarEdad(nuevaEdad int) {
	p.Edad = nuevaEdad
}

// Interfaz de ejemplo
type Hablador interface {
	Hablar() string
}

// Implementación de la interfaz
func (p Persona) Hablar() string {
	return "Estoy hablando"
}

func main() {
	fmt.Println("=== EJEMPLOS DE REFLECT Y TYPEOF EN GO ===")
	fmt.Println()

	// 1. Básicos de reflect.TypeOf y reflect.ValueOf
	fmt.Println("1. Básicos de reflect.TypeOf y reflect.ValueOf:")
	ejemplosBasicos()
	fmt.Println()

	// 2. Análisis de tipos primitivos
	fmt.Println("2. Análisis de tipos primitivos:")
	analizarTiposPrimitivos()
	fmt.Println()

	// 3. Análisis de estructuras
	fmt.Println("3. Análisis de estructuras:")
	analizarEstructuras()
	fmt.Println()

	// 4. Trabajando con campos de estructuras
	fmt.Println("4. Trabajando con campos de estructuras:")
	trabajarConCampos()
	fmt.Println()

	// 5. Análisis de métodos
	fmt.Println("5. Análisis de métodos:")
	analizarMetodos()
	fmt.Println()

	// 6. Modificando valores con reflection
	fmt.Println("6. Modificando valores con reflection:")
	modificarValores()
	fmt.Println()

	// 7. Trabajando con slices y arrays
	fmt.Println("7. Trabajando con slices y arrays:")
	trabajarConSlices()
	fmt.Println()

	// 8. Trabajando con maps
	fmt.Println("8. Trabajando con maps:")
	trabajarConMaps()
	fmt.Println()

	// 9. Interfaces y reflection
	fmt.Println("9. Interfaces y reflection:")
	trabajarConInterfaces()
	fmt.Println()

	// 10. Llamadas dinámicas a métodos
	fmt.Println("10. Llamadas dinámicas a métodos:")
	llamadasDinamicas()
	fmt.Println()

	// 11. Creación dinámica de tipos
	fmt.Println("11. Creación dinámica de tipos:")
	creacionDinamica()
	fmt.Println()

	fmt.Println("=== FIN DE LOS EJEMPLOS ===")
}

func ejemplosBasicos() {
	// Variables de diferentes tipos
	var numero int = 42
	var texto string = "Hola Go"
	var booleano bool = true
	var flotante float64 = 3.14159

	// reflect.TypeOf - obtiene el tipo
	fmt.Printf("Tipo de numero: %v\n", reflect.TypeOf(numero))
	fmt.Printf("Tipo de texto: %v\n", reflect.TypeOf(texto))
	fmt.Printf("Tipo de booleano: %v\n", reflect.TypeOf(booleano))
	fmt.Printf("Tipo de flotante: %v\n", reflect.TypeOf(flotante))

	// reflect.ValueOf - obtiene el valor reflectado
	valorNumero := reflect.ValueOf(numero)
	valorTexto := reflect.ValueOf(texto)

	fmt.Printf("Valor de numero: %v (Kind: %v)\n", valorNumero, valorNumero.Kind())
	fmt.Printf("Valor de texto: %v (Kind: %v)\n", valorTexto, valorTexto.Kind())

	// Verificar si es válido
	fmt.Printf("¿Es válido valorNumero?: %v\n", valorNumero.IsValid())
	fmt.Printf("¿Se puede establecer valorNumero?: %v\n", valorNumero.CanSet())
}

func analizarTiposPrimitivos() {
	valores := []interface{}{
		42,
		"cadena",
		3.14,
		true,
		'A',
		complex(1, 2),
	}

	for i, valor := range valores {
		tipo := reflect.TypeOf(valor)
		valorRef := reflect.ValueOf(valor)

		fmt.Printf("Elemento %d:\n", i+1)
		fmt.Printf("  Valor: %v\n", valor)
		fmt.Printf("  Tipo: %v\n", tipo)
		fmt.Printf("  Kind: %v\n", tipo.Kind())
		fmt.Printf("  Nombre: %v\n", tipo.Name())
		fmt.Printf("  Tamaño: %d bytes\n", tipo.Size())
		fmt.Printf("  Es comparable: %v\n", tipo.Comparable())
		fmt.Printf("  String del valor: %v\n", valorRef.String())
		fmt.Println()
	}
}

func analizarEstructuras() {
	persona := Persona{
		Nombre:  "Juan",
		Edad:    30,
		Email:   "juan@example.com",
		Activo:  true,
		privado: "secreto",
	}

	tipo := reflect.TypeOf(persona)
	valor := reflect.ValueOf(persona)

	fmt.Printf("Tipo de la estructura: %v\n", tipo)
	fmt.Printf("Kind: %v\n", tipo.Kind())
	fmt.Printf("Nombre del tipo: %v\n", tipo.Name())
	fmt.Printf("Paquete: %v\n", tipo.PkgPath())
	fmt.Printf("Número de campos: %d\n", tipo.NumField())
	fmt.Printf("Número de métodos: %d\n", tipo.NumMethod())

	// Información sobre cada campo
	fmt.Println("\nCampos de la estructura:")
	for i := 0; i < tipo.NumField(); i++ {
		campo := tipo.Field(i)
		valorCampo := valor.Field(i)

		fmt.Printf("  Campo %d:\n", i+1)
		fmt.Printf("    Nombre: %v\n", campo.Name)
		fmt.Printf("    Tipo: %v\n", campo.Type)
		fmt.Printf("    Tag: %v\n", campo.Tag)
		fmt.Printf("    Es exportado: %v\n", campo.IsExported())

		if valorCampo.CanInterface() {
			fmt.Printf("    Valor: %v\n", valorCampo.Interface())
		} else {
			fmt.Printf("    Valor: (no accesible)\n")
		}
		fmt.Println()
	}
}

func trabajarConCampos() {
	persona := Persona{
		Nombre: "María",
		Edad:   25,
		Email:  "maria@example.com",
		Activo: true,
	}

	valor := reflect.ValueOf(&persona).Elem() // Necesitamos el puntero para modificar

	// Acceder a campos por nombre
	campoNombre := valor.FieldByName("Nombre")
	campoEdad := valor.FieldByName("Edad")

	fmt.Printf("Nombre original: %v\n", campoNombre.Interface())
	fmt.Printf("Edad original: %v\n", campoEdad.Interface())

	// Modificar campos si es posible
	if campoNombre.CanSet() {
		campoNombre.SetString("María José")
		fmt.Printf("Nombre modificado: %v\n", campoNombre.Interface())
	}

	if campoEdad.CanSet() {
		campoEdad.SetInt(26)
		fmt.Printf("Edad modificada: %v\n", campoEdad.Interface())
	}

	// Trabajar con tags
	tipo := reflect.TypeOf(persona)
	campoEmail, _ := tipo.FieldByName("Email")
	jsonTag := campoEmail.Tag.Get("json")
	dbTag := campoEmail.Tag.Get("db")

	fmt.Printf("Tag JSON del campo Email: %v\n", jsonTag)
	fmt.Printf("Tag DB del campo Email: %v\n", dbTag)
}

func analizarMetodos() {
	persona := Persona{Nombre: "Carlos", Edad: 35}

	tipo := reflect.TypeOf(persona)
	valor := reflect.ValueOf(persona)

	fmt.Printf("Métodos de %v:\n", tipo)

	for i := 0; i < tipo.NumMethod(); i++ {
		metodo := tipo.Method(i)
		fmt.Printf("  Método %d:\n", i+1)
		fmt.Printf("    Nombre: %v\n", metodo.Name)
		fmt.Printf("    Tipo: %v\n", metodo.Type)
		fmt.Printf("    Es exportado: %v\n", metodo.IsExported())

		// Llamar al método si no tiene parámetros (excepto el receptor)
		if metodo.Type.NumIn() == 1 { // Solo el receptor
			resultado := valor.Method(i).Call(nil)
			if len(resultado) > 0 {
				fmt.Printf("    Resultado: %v\n", resultado[0].Interface())
			}
		}
		fmt.Println()
	}

	// También verificar métodos con puntero
	tipoPtr := reflect.TypeOf(&persona)
	fmt.Printf("\nMétodos adicionales con puntero (*%v):\n", tipo)
	for i := 0; i < tipoPtr.NumMethod(); i++ {
		metodo := tipoPtr.Method(i)
		fmt.Printf("  %v\n", metodo.Name)
	}
}

func modificarValores() {
	// Modificar variables básicas
	numero := 42
	texto := "original"

	valorNumero := reflect.ValueOf(&numero).Elem()
	valorTexto := reflect.ValueOf(&texto).Elem()

	fmt.Printf("Valores originales - numero: %d, texto: %s\n", numero, texto)

	if valorNumero.CanSet() {
		valorNumero.SetInt(100)
	}

	if valorTexto.CanSet() {
		valorTexto.SetString("modificado")
	}

	fmt.Printf("Valores modificados - numero: %d, texto: %s\n", numero, texto)

	// Modificar slice
	slice := []int{1, 2, 3}
	valorSlice := reflect.ValueOf(&slice).Elem()

	fmt.Printf("Slice original: %v\n", slice)

	// Agregar elemento al slice
	nuevoElemento := reflect.ValueOf(4)
	valorSlice.Set(reflect.Append(valorSlice, nuevoElemento))

	fmt.Printf("Slice modificado: %v\n", slice)
}

func trabajarConSlices() {
	numeros := []int{1, 2, 3, 4, 5}
	valor := reflect.ValueOf(numeros)

	fmt.Printf("Slice: %v\n", numeros)
	fmt.Printf("Tipo: %v\n", valor.Type())
	fmt.Printf("Kind: %v\n", valor.Kind())
	fmt.Printf("Longitud: %d\n", valor.Len())
	fmt.Printf("Capacidad: %d\n", valor.Cap())

	// Iterar sobre elementos
	fmt.Println("Elementos:")
	for i := 0; i < valor.Len(); i++ {
		elemento := valor.Index(i)
		fmt.Printf("  [%d]: %v (tipo: %v)\n", i, elemento.Interface(), elemento.Type())
	}

	// Crear slice dinámicamente
	tipoInt := reflect.TypeOf(0)
	nuevoSlice := reflect.MakeSlice(reflect.SliceOf(tipoInt), 0, 3)

	// Agregar elementos
	nuevoSlice = reflect.Append(nuevoSlice, reflect.ValueOf(10))
	nuevoSlice = reflect.Append(nuevoSlice, reflect.ValueOf(20))
	nuevoSlice = reflect.Append(nuevoSlice, reflect.ValueOf(30))

	fmt.Printf("Nuevo slice creado dinámicamente: %v\n", nuevoSlice.Interface())
}

func trabajarConMaps() {
	mapa := map[string]int{
		"uno":  1,
		"dos":  2,
		"tres": 3,
	}

	valor := reflect.ValueOf(mapa)

	fmt.Printf("Map: %v\n", mapa)
	fmt.Printf("Tipo: %v\n", valor.Type())
	fmt.Printf("Kind: %v\n", valor.Kind())
	fmt.Printf("Longitud: %d\n", valor.Len())

	// Iterar sobre el map
	fmt.Println("Elementos del map:")
	for _, clave := range valor.MapKeys() {
		valorElemento := valor.MapIndex(clave)
		fmt.Printf("  %v: %v\n", clave.Interface(), valorElemento.Interface())
	}

	// Crear map dinámicamente
	tipoMap := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	nuevoMap := reflect.MakeMap(tipoMap)

	// Agregar elementos
	nuevoMap.SetMapIndex(reflect.ValueOf("cuatro"), reflect.ValueOf(4))
	nuevoMap.SetMapIndex(reflect.ValueOf("cinco"), reflect.ValueOf(5))

	fmt.Printf("Nuevo map creado dinámicamente: %v\n", nuevoMap.Interface())
}

func trabajarConInterfaces() {
	var interfaz interface{} = "Hola mundo"

	valor := reflect.ValueOf(interfaz)
	tipo := reflect.TypeOf(interfaz)

	fmt.Printf("Interface contiene: %v\n", interfaz)
	fmt.Printf("Tipo real: %v\n", tipo)
	fmt.Printf("Kind: %v\n", valor.Kind())

	// Verificar si implementa una interfaz específica
	tipoHablador := reflect.TypeOf((*Hablador)(nil)).Elem()
	persona := Persona{Nombre: "Test"}

	tipoPersona := reflect.TypeOf(persona)
	fmt.Printf("¿Persona implementa Hablador?: %v\n", tipoPersona.Implements(tipoHablador))

	// Trabajar con interface{} que contiene diferentes tipos
	valores := []interface{}{42, "texto", 3.14, true, Persona{Nombre: "Juan"}}

	for i, v := range valores {
		valorRef := reflect.ValueOf(v)
		tipoRef := reflect.TypeOf(v)

		fmt.Printf("Elemento %d: %v (tipo: %v, kind: %v)\n",
			i+1, v, tipoRef, valorRef.Kind())

		// Hacer type assertion dinámico
		switch valorRef.Kind() {
		case reflect.Int:
			fmt.Printf("  Es un entero: %d\n", valorRef.Int())
		case reflect.String:
			fmt.Printf("  Es un string: %s\n", valorRef.String())
		case reflect.Float64:
			fmt.Printf("  Es un float: %.2f\n", valorRef.Float())
		case reflect.Bool:
			fmt.Printf("  Es un booleano: %t\n", valorRef.Bool())
		case reflect.Struct:
			fmt.Printf("  Es una estructura con %d campos\n", valorRef.NumField())
		}
	}
}

func llamadasDinamicas() {
	persona := Persona{Nombre: "Ana", Edad: 28}

	// Obtener método por nombre
	valor := reflect.ValueOf(persona)
	metodoSaludar := valor.MethodByName("Saludar")

	if metodoSaludar.IsValid() {
		// Llamar al método
		resultado := metodoSaludar.Call(nil)
		fmt.Printf("Resultado de Saludar(): %v\n", resultado[0].Interface())
	}

	// Llamar método con parámetros (necesitamos puntero)
	valorPtr := reflect.ValueOf(&persona)
	metodoCambiarEdad := valorPtr.MethodByName("CambiarEdad")

	if metodoCambiarEdad.IsValid() {
		// Preparar argumentos
		args := []reflect.Value{reflect.ValueOf(30)}
		metodoCambiarEdad.Call(args)
		fmt.Printf("Después de CambiarEdad(30): %s\n", persona.Saludar())
	}

	// Llamar función global dinámicamente
	funcionStrconv := reflect.ValueOf(strconv.Itoa)
	argumentos := []reflect.Value{reflect.ValueOf(123)}
	resultadoStrconv := funcionStrconv.Call(argumentos)

	fmt.Printf("strconv.Itoa(123) = %v\n", resultadoStrconv[0].Interface())
}

func creacionDinamica() {
	// Crear instancia de tipo por nombre
	tipoPersona := reflect.TypeOf(Persona{})
	nuevaPersona := reflect.New(tipoPersona).Elem()

	fmt.Printf("Nueva instancia creada: %v\n", nuevaPersona.Interface())

	// Establecer valores en los campos
	campoNombre := nuevaPersona.FieldByName("Nombre")
	campoEdad := nuevaPersona.FieldByName("Edad")

	if campoNombre.CanSet() {
		campoNombre.SetString("Persona Dinámica")
	}
	if campoEdad.CanSet() {
		campoEdad.SetInt(25)
	}

	fmt.Printf("Después de establecer valores: %v\n", nuevaPersona.Interface())

	// Crear slice dinámicamente
	tipoSlice := reflect.SliceOf(reflect.TypeOf(0))
	nuevoSlice := reflect.MakeSlice(tipoSlice, 3, 5)

	// Establecer valores
	nuevoSlice.Index(0).SetInt(10)
	nuevoSlice.Index(1).SetInt(20)
	nuevoSlice.Index(2).SetInt(30)

	fmt.Printf("Slice creado dinámicamente: %v\n", nuevoSlice.Interface())

	// Crear puntero dinámicamente
	puntero := reflect.New(reflect.TypeOf(""))
	puntero.Elem().SetString("Valor desde puntero")

	fmt.Printf("Valor del puntero: %v\n", puntero.Elem().Interface())
	fmt.Printf("Dirección del puntero: %v\n", puntero.Interface())
}
