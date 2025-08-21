package main

import (
	"fmt"
	"strings"
	"time"
)

// ============= PATRÓN STRATEGY =============

// Interfaz para estrategias de cálculo de descuentos
type EstrategiaDescuento interface {
	CalcularDescuento(precio float64) float64
	Descripcion() string
}

// Estrategia: Sin descuento
type SinDescuento struct{}

func (sd SinDescuento) CalcularDescuento(precio float64) float64 {
	return 0
}

func (sd SinDescuento) Descripcion() string {
	return "Sin descuento"
}

// Estrategia: Descuento por porcentaje
type DescuentoPorcentaje struct {
	Porcentaje float64
}

func (dp DescuentoPorcentaje) CalcularDescuento(precio float64) float64 {
	return precio * (dp.Porcentaje / 100)
}

func (dp DescuentoPorcentaje) Descripcion() string {
	return fmt.Sprintf("Descuento del %.1f%%", dp.Porcentaje)
}

// Estrategia: Descuento por cantidad fija
type DescuentoFijo struct {
	MontoFijo float64
}

func (df DescuentoFijo) CalcularDescuento(precio float64) float64 {
	descuento := df.MontoFijo
	if descuento > precio {
		descuento = precio // No puede ser mayor al precio
	}
	return descuento
}

func (df DescuentoFijo) Descripcion() string {
	return fmt.Sprintf("Descuento fijo de $%.2f", df.MontoFijo)
}

// Estrategia: Descuento por escalas
type DescuentoPorEscala struct {
	Escalas []EscalaDescuento
}

type EscalaDescuento struct {
	MontoMinimo float64
	Porcentaje  float64
}

func (dpe DescuentoPorEscala) CalcularDescuento(precio float64) float64 {
	var mejorPorcentaje float64

	for _, escala := range dpe.Escalas {
		if precio >= escala.MontoMinimo && escala.Porcentaje > mejorPorcentaje {
			mejorPorcentaje = escala.Porcentaje
		}
	}

	return precio * (mejorPorcentaje / 100)
}

func (dpe DescuentoPorEscala) Descripcion() string {
	var descripcion []string
	for _, escala := range dpe.Escalas {
		descripcion = append(descripcion,
			fmt.Sprintf("$%.0f+: %.1f%%", escala.MontoMinimo, escala.Porcentaje))
	}
	return fmt.Sprintf("Descuento por escalas: %s", strings.Join(descripcion, ", "))
}

// Contexto que usa estrategias
type CalculadoraPrecio struct {
	estrategia EstrategiaDescuento
}

func NuevaCalculadoraPrecio(estrategia EstrategiaDescuento) *CalculadoraPrecio {
	return &CalculadoraPrecio{estrategia: estrategia}
}

func (cp *CalculadoraPrecio) EstablecerEstrategia(estrategia EstrategiaDescuento) {
	cp.estrategia = estrategia
}

func (cp *CalculadoraPrecio) CalcularPrecioFinal(precio float64) (float64, float64, string) {
	descuento := cp.estrategia.CalcularDescuento(precio)
	precioFinal := precio - descuento
	descripcion := cp.estrategia.Descripcion()
	return precioFinal, descuento, descripcion
}

// ============= PATRÓN STRATEGY PARA ALGORITMOS DE ORDENAMIENTO =============

// Interfaz para estrategias de ordenamiento
type EstrategiaOrdenamiento interface {
	Ordenar(datos []int) []int
	Nombre() string
}

// Estrategia: Bubble Sort
type BubbleSort struct{}

func (bs BubbleSort) Ordenar(datos []int) []int {
	copia := make([]int, len(datos))
	copy(copia, datos)

	n := len(copia)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if copia[j] > copia[j+1] {
				copia[j], copia[j+1] = copia[j+1], copia[j]
			}
		}
	}
	return copia
}

func (bs BubbleSort) Nombre() string {
	return "Bubble Sort"
}

// Estrategia: Selection Sort
type SelectionSort struct{}

func (ss SelectionSort) Ordenar(datos []int) []int {
	copia := make([]int, len(datos))
	copy(copia, datos)

	n := len(copia)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if copia[j] < copia[minIdx] {
				minIdx = j
			}
		}
		copia[i], copia[minIdx] = copia[minIdx], copia[i]
	}
	return copia
}

func (ss SelectionSort) Nombre() string {
	return "Selection Sort"
}

// Estrategia: Quick Sort (simplificado)
type QuickSort struct{}

func (qs QuickSort) Ordenar(datos []int) []int {
	copia := make([]int, len(datos))
	copy(copia, datos)
	qs.quickSort(copia, 0, len(copia)-1)
	return copia
}

func (qs QuickSort) quickSort(arr []int, low, high int) {
	if low < high {
		pi := qs.partition(arr, low, high)
		qs.quickSort(arr, low, pi-1)
		qs.quickSort(arr, pi+1, high)
	}
}

func (qs QuickSort) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func (qs QuickSort) Nombre() string {
	return "Quick Sort"
}

// Contexto para ordenamiento
type OrdenadorDatos struct {
	estrategia EstrategiaOrdenamiento
}

func NuevoOrdenadorDatos(estrategia EstrategiaOrdenamiento) *OrdenadorDatos {
	return &OrdenadorDatos{estrategia: estrategia}
}

func (od *OrdenadorDatos) CambiarEstrategia(estrategia EstrategiaOrdenamiento) {
	od.estrategia = estrategia
}

func (od *OrdenadorDatos) OrdenarYMedir(datos []int) ([]int, time.Duration) {
	inicio := time.Now()
	resultado := od.estrategia.Ordenar(datos)
	duracion := time.Since(inicio)
	return resultado, duracion
}

func (od *OrdenadorDatos) NombreEstrategia() string {
	return od.estrategia.Nombre()
}

// ============= PATRÓN STRATEGY PARA VALIDACIONES =============

// Interfaz para estrategias de validación
type ValidadorStrategy interface {
	Validar(valor string) (bool, string)
	TipoValidacion() string
}

// Validador de email
type ValidadorEmail struct{}

func (ve ValidadorEmail) Validar(valor string) (bool, string) {
	if !strings.Contains(valor, "@") {
		return false, "debe contener @"
	}
	if !strings.Contains(valor, ".") {
		return false, "debe contener un dominio válido"
	}
	if len(valor) < 5 {
		return false, "muy corto para ser un email válido"
	}
	return true, "email válido"
}

func (ve ValidadorEmail) TipoValidacion() string {
	return "Email"
}

// Validador de contraseña
type ValidadorPassword struct {
	LongitudMinima   int
	RequiereMayus    bool
	RequiereMinus    bool
	RequiereNumero   bool
	RequiereEspecial bool
}

func (vp ValidadorPassword) Validar(valor string) (bool, string) {
	if len(valor) < vp.LongitudMinima {
		return false, fmt.Sprintf("debe tener al menos %d caracteres", vp.LongitudMinima)
	}

	if vp.RequiereMayus && !strings.ContainsAny(valor, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false, "debe contener al menos una mayúscula"
	}

	if vp.RequiereMinus && !strings.ContainsAny(valor, "abcdefghijklmnopqrstuvwxyz") {
		return false, "debe contener al menos una minúscula"
	}

	if vp.RequiereNumero && !strings.ContainsAny(valor, "0123456789") {
		return false, "debe contener al menos un número"
	}

	if vp.RequiereEspecial && !strings.ContainsAny(valor, "!@#$%^&*()_+-=[]{}|;:,.<>?") {
		return false, "debe contener al menos un carácter especial"
	}

	return true, "contraseña válida"
}

func (vp ValidadorPassword) TipoValidacion() string {
	return "Contraseña"
}

// Validador de teléfono
type ValidadorTelefono struct{}

func (vt ValidadorTelefono) Validar(valor string) (bool, string) {
	// Remover espacios y guiones
	telefono := strings.ReplaceAll(strings.ReplaceAll(valor, " ", ""), "-", "")

	if len(telefono) < 10 {
		return false, "debe tener al menos 10 dígitos"
	}

	// Verificar que solo contenga números y +
	for _, char := range telefono {
		if !strings.ContainsRune("0123456789+", char) {
			return false, "solo puede contener números, + y guiones"
		}
	}

	return true, "número de teléfono válido"
}

func (vt ValidadorTelefono) TipoValidacion() string {
	return "Teléfono"
}

// Contexto para validación
type ValidadorFormulario struct {
	validadores map[string]ValidadorStrategy
}

func NuevoValidadorFormulario() *ValidadorFormulario {
	return &ValidadorFormulario{
		validadores: make(map[string]ValidadorStrategy),
	}
}

func (vf *ValidadorFormulario) AgregarValidador(campo string, validador ValidadorStrategy) {
	vf.validadores[campo] = validador
}

func (vf *ValidadorFormulario) ValidarCampo(campo, valor string) (bool, string) {
	if validador, existe := vf.validadores[campo]; existe {
		return validador.Validar(valor)
	}
	return false, "no hay validador configurado para este campo"
}

func (vf *ValidadorFormulario) ValidarFormulario(datos map[string]string) map[string]string {
	errores := make(map[string]string)

	for campo, valor := range datos {
		if valido, mensaje := vf.ValidarCampo(campo, valor); !valido {
			errores[campo] = mensaje
		}
	}

	return errores
}

func ejemploPatronStrategy() {
	fmt.Println("  === STRATEGY: CÁLCULO DE DESCUENTOS ===")

	precios := []float64{50, 150, 500, 1000}

	// Crear diferentes estrategias
	estrategias := []EstrategiaDescuento{
		SinDescuento{},
		DescuentoPorcentaje{Porcentaje: 10},
		DescuentoFijo{MontoFijo: 25},
		DescuentoPorEscala{
			Escalas: []EscalaDescuento{
				{MontoMinimo: 100, Porcentaje: 5},
				{MontoMinimo: 300, Porcentaje: 10},
				{MontoMinimo: 500, Porcentaje: 15},
				{MontoMinimo: 1000, Porcentaje: 20},
			},
		},
	}

	calculadora := NuevaCalculadoraPrecio(SinDescuento{})

	for _, estrategia := range estrategias {
		fmt.Printf("    %s:\n", estrategia.Descripcion())
		calculadora.EstablecerEstrategia(estrategia)

		for _, precio := range precios {
			precioFinal, descuento, _ := calculadora.CalcularPrecioFinal(precio)
			fmt.Printf("      Precio: $%.2f -> Descuento: $%.2f -> Final: $%.2f\n",
				precio, descuento, precioFinal)
		}
		fmt.Println()
	}

	fmt.Println("  === STRATEGY: ALGORITMOS DE ORDENAMIENTO ===")

	datos := []int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50, 42}
	fmt.Printf("    Datos originales: %v\n\n", datos)

	algoritmos := []EstrategiaOrdenamiento{
		BubbleSort{},
		SelectionSort{},
		QuickSort{},
	}

	ordenador := NuevoOrdenadorDatos(BubbleSort{})

	for _, algoritmo := range algoritmos {
		ordenador.CambiarEstrategia(algoritmo)
		resultado, duracion := ordenador.OrdenarYMedir(datos)

		fmt.Printf("    %s:\n", ordenador.NombreEstrategia())
		fmt.Printf("      Resultado: %v\n", resultado)
		fmt.Printf("      Tiempo: %v\n\n", duracion)
	}

	fmt.Println("  === STRATEGY: VALIDACIÓN DE FORMULARIOS ===")

	// Configurar validadores
	validador := NuevoValidadorFormulario()
	validador.AgregarValidador("email", ValidadorEmail{})
	validador.AgregarValidador("password", ValidadorPassword{
		LongitudMinima:   8,
		RequiereMayus:    true,
		RequiereMinus:    true,
		RequiereNumero:   true,
		RequiereEspecial: true,
	})
	validador.AgregarValidador("telefono", ValidadorTelefono{})

	// Datos de prueba
	formularios := []map[string]string{
		{
			"email":    "usuario@ejemplo.com",
			"password": "MiPassword123!",
			"telefono": "+54 11 1234-5678",
		},
		{
			"email":    "email_invalido",
			"password": "123",
			"telefono": "abc",
		},
		{
			"email":    "otro@test.org",
			"password": "SoloMayusculas",
			"telefono": "1234567890",
		},
	}

	for i, datos := range formularios {
		fmt.Printf("    Formulario %d:\n", i+1)
		errores := validador.ValidarFormulario(datos)

		if len(errores) == 0 {
			fmt.Println("      ✓ Todos los campos son válidos")
		} else {
			fmt.Println("      ✗ Errores encontrados:")
			for campo, error := range errores {
				fmt.Printf("        %s: %s\n", campo, error)
			}
		}
		fmt.Println()
	}
}
