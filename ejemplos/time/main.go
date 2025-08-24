package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== EJEMPLOS DE MANEJO DE TIEMPO EN GO ===")
	fmt.Println()

	// 1. Obtener la fecha y hora actual
	ejemploTiempoActual()

	// 2. Formatear fechas
	ejemploFormateoFechas()

	// 3. Parsear strings a fechas
	ejemploParsearFechas()

	// 4. Agregar y quitar tiempo
	ejemploAgregarQuitarTiempo()

	// 5. Comparar fechas
	ejemploCompararFechas()

	// 6. Trabajar con zonas horarias
	ejemploZonasHorarias()

	// 7. Duración y medición de tiempo
	ejemploDuracion()

	// 8. Timestamps Unix
	ejemploTimestamps()

	// 9. Fechas específicas
	ejemploFechasEspecificas()

	// 10. Trabajar con diferentes formatos
	ejemploDiferentesFormatos()
}

// 1. Obtener la fecha y hora actual
func ejemploTiempoActual() {
	fmt.Println("1. TIEMPO ACTUAL:")

	now := time.Now()
	fmt.Printf("Fecha y hora actual: %v\n", now)
	fmt.Printf("Solo fecha: %v\n", now.Format("2006-01-02"))
	fmt.Printf("Solo hora: %v\n", now.Format("15:04:05"))
	fmt.Printf("Año: %d\n", now.Year())
	fmt.Printf("Mes: %v (%d)\n", now.Month(), int(now.Month()))
	fmt.Printf("Día: %d\n", now.Day())
	fmt.Printf("Día de la semana: %v\n", now.Weekday())
	fmt.Printf("Día del año: %d\n", now.YearDay())
	fmt.Printf("Hora: %d\n", now.Hour())
	fmt.Printf("Minuto: %d\n", now.Minute())
	fmt.Printf("Segundo: %d\n", now.Second())
	fmt.Printf("Nanosegundo: %d\n", now.Nanosecond())
	fmt.Println()
}

// 2. Formatear fechas
func ejemploFormateoFechas() {
	fmt.Println("2. FORMATEO DE FECHAS:")

	now := time.Now()

	// El layout de referencia en Go es: Mon Jan 2 15:04:05 MST 2006
	// que corresponde a: Unix timestamp 1136239445
	fmt.Printf("RFC3339: %v\n", now.Format(time.RFC3339))
	fmt.Printf("RFC822: %v\n", now.Format(time.RFC822))
	fmt.Printf("Formato personalizado 1: %v\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Formato personalizado 2: %v\n", now.Format("02/01/2006 3:04 PM"))
	fmt.Printf("Formato personalizado 3: %v\n", now.Format("January 2, 2006"))
	fmt.Printf("Formato personalizado 4: %v\n", now.Format("Mon, 02 Jan 2006 15:04:05 MST"))
	fmt.Printf("Solo fecha española: %v\n", now.Format("02/01/2006"))
	fmt.Printf("Formato ISO 8601: %v\n", now.Format("2006-01-02T15:04:05Z07:00"))
	fmt.Println()
}

// 3. Parsear strings a fechas
func ejemploParsearFechas() {
	fmt.Println("3. PARSEAR STRINGS A FECHAS:")

	// Diferentes formatos de fecha
	fechas := []string{
		"2023-12-25 14:30:00",
		"25/12/2023",
		"December 25, 2023",
		"2023-12-25T14:30:00Z",
	}

	layouts := []string{
		"2006-01-02 15:04:05",
		"02/01/2006",
		"January 2, 2006",
		time.RFC3339,
	}

	for i, fechaStr := range fechas {
		parsed, err := time.Parse(layouts[i], fechaStr)
		if err != nil {
			fmt.Printf("Error parseando '%s': %v\n", fechaStr, err)
		} else {
			fmt.Printf("Parseado '%s' -> %v\n", fechaStr, parsed)
		}
	}
	fmt.Println()
}

// 4. Agregar y quitar tiempo
func ejemploAgregarQuitarTiempo() {
	fmt.Println("4. AGREGAR Y QUITAR TIEMPO:")

	now := time.Now()
	fmt.Printf("Fecha actual: %v\n", now.Format("2006-01-02 15:04:05"))

	// Agregar tiempo
	fmt.Printf("+ 1 día: %v\n", now.AddDate(0, 0, 1).Format("2006-01-02 15:04:05"))
	fmt.Printf("+ 1 semana: %v\n", now.AddDate(0, 0, 7).Format("2006-01-02 15:04:05"))
	fmt.Printf("+ 1 mes: %v\n", now.AddDate(0, 1, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("+ 1 año: %v\n", now.AddDate(1, 0, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("+ 2 horas: %v\n", now.Add(2*time.Hour).Format("2006-01-02 15:04:05"))
	fmt.Printf("+ 30 minutos: %v\n", now.Add(30*time.Minute).Format("2006-01-02 15:04:05"))
	fmt.Printf("+ 45 segundos: %v\n", now.Add(45*time.Second).Format("2006-01-02 15:04:05"))

	// Quitar tiempo
	fmt.Printf("- 1 día: %v\n", now.AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
	fmt.Printf("- 1 mes: %v\n", now.AddDate(0, -1, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("- 1 año: %v\n", now.AddDate(-1, 0, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("- 3 horas: %v\n", now.Add(-3*time.Hour).Format("2006-01-02 15:04:05"))

	// Fechas futuras específicas
	fmt.Printf("En 15 días: %v\n", now.AddDate(0, 0, 15).Format("2006-01-02"))
	fmt.Printf("En 3 meses: %v\n", now.AddDate(0, 3, 0).Format("2006-01-02"))
	fmt.Printf("El próximo año: %v\n", now.AddDate(1, 0, 0).Format("2006-01-02"))

	fmt.Println()
}

// 5. Comparar fechas
func ejemploCompararFechas() {
	fmt.Println("5. COMPARAR FECHAS:")

	now := time.Now()
	ayer := now.AddDate(0, 0, -1)
	mañana := now.AddDate(0, 0, 1)

	fmt.Printf("Hoy: %v\n", now.Format("2006-01-02"))
	fmt.Printf("Ayer: %v\n", ayer.Format("2006-01-02"))
	fmt.Printf("Mañana: %v\n", mañana.Format("2006-01-02"))

	fmt.Printf("¿Ayer es antes que hoy? %v\n", ayer.Before(now))
	fmt.Printf("¿Mañana es después que hoy? %v\n", mañana.After(now))
	fmt.Printf("¿Hoy es igual a hoy? %v\n", now.Equal(now))

	// Diferencia entre fechas
	diferencia := mañana.Sub(now)
	fmt.Printf("Diferencia entre mañana y hoy: %v\n", diferencia)
	fmt.Printf("Diferencia en horas: %.0f\n", diferencia.Hours())

	// Verificar si una fecha está en un rango
	inicio := now.AddDate(0, 0, -5)
	fin := now.AddDate(0, 0, 5)
	fechaTest := now.AddDate(0, 0, 2)

	fmt.Printf("¿La fecha %v está entre %v y %v? %v\n",
		fechaTest.Format("2006-01-02"),
		inicio.Format("2006-01-02"),
		fin.Format("2006-01-02"),
		fechaTest.After(inicio) && fechaTest.Before(fin))

	fmt.Println()
}

// 6. Trabajar con zonas horarias
func ejemploZonasHorarias() {
	fmt.Println("6. ZONAS HORARIAS:")

	now := time.Now()
	fmt.Printf("Hora local: %v\n", now.Format("2006-01-02 15:04:05 MST"))

	// UTC
	utc := now.UTC()
	fmt.Printf("Hora UTC: %v\n", utc.Format("2006-01-02 15:04:05 MST"))

	// Cargar diferentes zonas horarias
	locations := []string{
		"America/New_York",
		"America/Los_Angeles",
		"Europe/London",
		"Europe/Madrid",
		"Asia/Tokyo",
		"Australia/Sydney",
	}

	for _, locName := range locations {
		loc, err := time.LoadLocation(locName)
		if err != nil {
			fmt.Printf("Error cargando zona horaria %s: %v\n", locName, err)
			continue
		}

		timeInLoc := now.In(loc)
		fmt.Printf("Hora en %-20s: %v\n", locName, timeInLoc.Format("2006-01-02 15:04:05 MST"))
	}

	fmt.Println()
}

// 7. Duración y medición de tiempo
func ejemploDuracion() {
	fmt.Println("7. DURACIÓN Y MEDICIÓN DE TIEMPO:")

	// Medir tiempo de ejecución
	inicio := time.Now()

	// Simular algún trabajo
	time.Sleep(100 * time.Millisecond)

	duracion := time.Since(inicio)
	fmt.Printf("Tiempo transcurrido: %v\n", duracion)
	fmt.Printf("Tiempo en milisegundos: %.2f ms\n", float64(duracion.Nanoseconds())/1e6)

	// Diferentes unidades de duración
	fmt.Printf("1 segundo = %v nanosegundos\n", time.Second)
	fmt.Printf("1 minuto = %v segundos\n", time.Minute/time.Second)
	fmt.Printf("1 hora = %v minutos\n", time.Hour/time.Minute)
	fmt.Printf("24 horas = %v horas\n", 24*time.Hour/time.Hour)

	// Crear duraciones personalizadas
	duracionPersonalizada := 2*time.Hour + 30*time.Minute + 45*time.Second
	fmt.Printf("Duración personalizada: %v\n", duracionPersonalizada)
	fmt.Printf("Total en segundos: %.0f\n", duracionPersonalizada.Seconds())

	fmt.Println()
}

// 8. Timestamps Unix
func ejemploTimestamps() {
	fmt.Println("8. TIMESTAMPS UNIX:")

	now := time.Now()

	// Timestamp Unix (segundos desde epoch)
	timestamp := now.Unix()
	fmt.Printf("Timestamp Unix: %d\n", timestamp)

	// Timestamp Unix en nanosegundos
	timestampNano := now.UnixNano()
	fmt.Printf("Timestamp Unix (nano): %d\n", timestampNano)

	// Timestamp Unix en milisegundos
	timestampMilli := now.UnixMilli()
	fmt.Printf("Timestamp Unix (milli): %d\n", timestampMilli)

	// Convertir timestamp de vuelta a time.Time
	fechaDesdeTimestamp := time.Unix(timestamp, 0)
	fmt.Printf("Fecha desde timestamp: %v\n", fechaDesdeTimestamp.Format("2006-01-02 15:04:05"))

	// Timestamp específico (1 de enero de 2024)
	fechaEspecifica := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("1 de enero 2024 timestamp: %d\n", fechaEspecifica.Unix())

	fmt.Println()
}

// 9. Fechas específicas
func ejemploFechasEspecificas() {
	fmt.Println("9. CREAR FECHAS ESPECÍFICAS:")

	// Crear fecha específica
	navidad2024 := time.Date(2024, time.December, 25, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Navidad 2024: %v\n", navidad2024.Format("Monday, January 2, 2006"))

	añoNuevo2025 := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Año Nuevo 2025: %v\n", añoNuevo2025.Format("Monday, January 2, 2006"))

	// Calcular días hasta una fecha específica
	now := time.Now()
	diasHastaNavidad := navidad2024.Sub(now).Hours() / 24
	if diasHastaNavidad > 0 {
		fmt.Printf("Días hasta Navidad 2024: %.0f días\n", diasHastaNavidad)
	} else {
		fmt.Printf("Navidad 2024 ya pasó hace %.0f días\n", -diasHastaNavidad)
	}

	// Primera fecha del mes actual
	primerDiaDelMes := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	fmt.Printf("Primer día de este mes: %v\n", primerDiaDelMes.Format("2006-01-02"))

	// Último día del mes actual
	ultimoDiaDelMes := primerDiaDelMes.AddDate(0, 1, -1)
	fmt.Printf("Último día de este mes: %v\n", ultimoDiaDelMes.Format("2006-01-02"))

	// Inicio y fin de semana
	diasHastaLunes := int(now.Weekday()) - 1
	if diasHastaLunes < 0 {
		diasHastaLunes = 6 // Domingo
	}
	inicioSemana := now.AddDate(0, 0, -diasHastaLunes)
	finSemana := inicioSemana.AddDate(0, 0, 6)

	fmt.Printf("Inicio de esta semana (lunes): %v\n", inicioSemana.Format("2006-01-02"))
	fmt.Printf("Fin de esta semana (domingo): %v\n", finSemana.Format("2006-01-02"))

	fmt.Println()
}

// 10. Trabajar con diferentes formatos
func ejemploDiferentesFormatos() {
	fmt.Println("10. DIFERENTES FORMATOS DE FECHA:")

	now := time.Now()

	// Formatos comunes
	formatos := map[string]string{
		"ISO 8601":          "2006-01-02T15:04:05Z07:00",
		"RFC 3339":          time.RFC3339,
		"RFC 822":           time.RFC822,
		"RFC 850":           time.RFC850,
		"SQL DateTime":      "2006-01-02 15:04:05",
		"Fecha española":    "02/01/2006",
		"Fecha americana":   "01/02/2006",
		"Fecha con texto":   "January 2, 2006",
		"Fecha compacta":    "20060102",
		"Hora 12h":          "3:04:05 PM",
		"Hora militar":      "15:04:05",
		"Mes y año":         "January 2006",
		"Solo día y mes":    "02 Jan",
		"Timestamp legible": "Mon Jan 2 15:04:05 2006",
	}

	fmt.Printf("Fecha actual en diferentes formatos:\n")
	for nombre, formato := range formatos {
		fmt.Printf("%-20s: %s\n", nombre, now.Format(formato))
	}

	fmt.Println()

	// Ejemplo de validación de fecha
	fmt.Println("VALIDACIÓN DE FECHAS:")
	fechasParaValidar := []string{
		"2024-02-29", // Año bisiesto válido
		"2023-02-29", // Año no bisiesto inválido
		"2024-13-01", // Mes inválido
		"2024-12-32", // Día inválido
		"2024-12-25", // Fecha válida
	}

	for _, fecha := range fechasParaValidar {
		_, err := time.Parse("2006-01-02", fecha)
		if err != nil {
			fmt.Printf("Fecha '%s': INVÁLIDA ❌\n", fecha)
		} else {
			fmt.Printf("Fecha '%s': VÁLIDA ✅\n", fecha)
		}
	}

	fmt.Println("\n=== FIN DE EJEMPLOS ===")
}
