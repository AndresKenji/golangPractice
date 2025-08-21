package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== EJEMPLOS DE CONDICIONALES EN GO ===\n")

	// 1. IF básico
	fmt.Println("1. IF básico:")
	edad := 18
	if edad >= 18 {
		fmt.Println("   Es mayor de edad")
	}
	fmt.Println()

	// 2. IF-ELSE
	fmt.Println("2. IF-ELSE:")
	temperatura := 25
	if temperatura > 30 {
		fmt.Println("   Hace calor")
	} else {
		fmt.Println("   El clima está agradable")
	}
	fmt.Println()

	// 3. IF-ELSE IF-ELSE
	fmt.Println("3. IF-ELSE IF-ELSE:")
	nota := 85
	if nota >= 90 {
		fmt.Println("   Calificación: A")
	} else if nota >= 80 {
		fmt.Println("   Calificación: B")
	} else if nota >= 70 {
		fmt.Println("   Calificación: C")
	} else if nota >= 60 {
		fmt.Println("   Calificación: D")
	} else {
		fmt.Println("   Calificación: F")
	}
	fmt.Println()

	// 4. IF con declaración de variable
	fmt.Println("4. IF con declaración de variable:")
	if numero := 42; numero%2 == 0 {
		fmt.Printf("   %d es par\n", numero)
	} else {
		fmt.Printf("   %d es impar\n", numero)
	}
	fmt.Println()

	// 5. IF con múltiples condiciones (AND)
	fmt.Println("5. IF con múltiples condiciones (AND):")
	usuario := "admin"
	password := "123456"
	if usuario == "admin" && password == "123456" {
		fmt.Println("   Acceso concedido")
	} else {
		fmt.Println("   Acceso denegado")
	}
	fmt.Println()

	// 6. IF con múltiples condiciones (OR)
	fmt.Println("6. IF con múltiples condiciones (OR):")
	dia := "sábado"
	if dia == "sábado" || dia == "domingo" {
		fmt.Println("   Es fin de semana!")
	} else {
		fmt.Println("   Es día laboral")
	}
	fmt.Println()

	// 7. IF con negación
	fmt.Println("7. IF con negación:")
	estaLloviendo := false
	if !estaLloviendo {
		fmt.Println("   Puedes salir sin paraguas")
	}
	fmt.Println()

	// 8. SWITCH básico
	fmt.Println("8. SWITCH básico:")
	diaSemana := 3
	switch diaSemana {
	case 1:
		fmt.Println("   Lunes")
	case 2:
		fmt.Println("   Martes")
	case 3:
		fmt.Println("   Miércoles")
	case 4:
		fmt.Println("   Jueves")
	case 5:
		fmt.Println("   Viernes")
	case 6:
		fmt.Println("   Sábado")
	case 7:
		fmt.Println("   Domingo")
	default:
		fmt.Println("   Día inválido")
	}
	fmt.Println()

	// 9. SWITCH con múltiples valores
	fmt.Println("9. SWITCH con múltiples valores:")
	mes := "enero"
	switch mes {
	case "diciembre", "enero", "febrero":
		fmt.Println("   Invierno")
	case "marzo", "abril", "mayo":
		fmt.Println("   Primavera")
	case "junio", "julio", "agosto":
		fmt.Println("   Verano")
	case "septiembre", "octubre", "noviembre":
		fmt.Println("   Otoño")
	default:
		fmt.Println("   Mes inválido")
	}
	fmt.Println()

	// 10. SWITCH con condiciones
	fmt.Println("10. SWITCH con condiciones:")
	puntuacion := 95
	switch {
	case puntuacion >= 90:
		fmt.Println("   Excelente!")
	case puntuacion >= 80:
		fmt.Println("   Muy bien!")
	case puntuacion >= 70:
		fmt.Println("   Bien")
	case puntuacion >= 60:
		fmt.Println("   Suficiente")
	default:
		fmt.Println("   Necesitas mejorar")
	}
	fmt.Println()

	// 11. SWITCH con inicialización
	fmt.Println("11. SWITCH con inicialización:")
	switch hora := time.Now().Hour(); {
	case hora < 12:
		fmt.Println("   Buenos días!")
	case hora < 18:
		fmt.Println("   Buenas tardes!")
	default:
		fmt.Println("   Buenas noches!")
	}
	fmt.Println()

	// 12. SWITCH con fallthrough
	fmt.Println("12. SWITCH con fallthrough:")
	numero2 := 2
	switch numero2 {
	case 1:
		fmt.Println("   Uno")
		fallthrough
	case 2:
		fmt.Println("   Dos")
		fallthrough
	case 3:
		fmt.Println("   Tres o menos")
	default:
		fmt.Println("   Número mayor a 3")
	}
	fmt.Println()

	// 13. SWITCH con tipos (type switch)
	fmt.Println("13. SWITCH con tipos:")
	var valor interface{} = "Hola mundo"
	switch v := valor.(type) {
	case int:
		fmt.Printf("   Es un entero: %d\n", v)
	case string:
		fmt.Printf("   Es una cadena: %s\n", v)
	case bool:
		fmt.Printf("   Es un booleano: %t\n", v)
	default:
		fmt.Printf("   Tipo desconocido: %T\n", v)
	}
	fmt.Println()

	// 14. Condicionales anidados
	fmt.Println("14. Condicionales anidados:")
	edad2 := 25
	tieneTrabaajo := true
	if edad2 >= 18 {
		if tieneTrabaajo {
			fmt.Println("   Puede solicitar un crédito")
		} else {
			fmt.Println("   Necesita tener trabajo para el crédito")
		}
	} else {
		fmt.Println("   Debe ser mayor de edad")
	}
	fmt.Println()

	// 15. Operador ternario simulado con función
	fmt.Println("15. Operador ternario simulado:")
	resultado := ternario(edad >= 18, "Mayor de edad", "Menor de edad")
	fmt.Printf("   %s\n", resultado)
	fmt.Println()

	// 16. Validación de rangos
	fmt.Println("16. Validación de rangos:")
	validarRango(5)
	validarRango(15)
	validarRango(25)
	fmt.Println()

	// 17. Validación con múltiples criterios
	fmt.Println("17. Validación con múltiples criterios:")
	validarUsuario("juan", 25, true)
	validarUsuario("ana", 16, false)
	fmt.Println()

	fmt.Println("=== FIN DE EJEMPLOS ===")
}

// Función auxiliar para simular operador ternario
func ternario(condicion bool, verdadero, falso string) string {
	if condicion {
		return verdadero
	}
	return falso
}

// Función para validar rangos
func validarRango(numero int) {
	switch {
	case numero < 0:
		fmt.Printf("   %d es negativo\n", numero)
	case numero >= 0 && numero <= 10:
		fmt.Printf("   %d está en el rango 0-10\n", numero)
	case numero > 10 && numero <= 20:
		fmt.Printf("   %d está en el rango 11-20\n", numero)
	default:
		fmt.Printf("   %d es mayor a 20\n", numero)
	}
}

// Función para validar usuario con múltiples criterios
func validarUsuario(nombre string, edad int, activo bool) {
	if nombre != "" && edad >= 18 && activo {
		fmt.Printf("   Usuario %s: VÁLIDO (edad: %d, activo: %t)\n", nombre, edad, activo)
	} else {
		fmt.Printf("   Usuario %s: INVÁLIDO", nombre)
		if nombre == "" {
			fmt.Print(" (sin nombre)")
		}
		if edad < 18 {
			fmt.Print(" (menor de edad)")
		}
		if !activo {
			fmt.Print(" (inactivo)")
		}
		fmt.Println()
	}
}