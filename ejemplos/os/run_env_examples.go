package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	ejemploVariablesEntornoAvanzado()
}

// Función para demostrar manejo avanzado de variables de entorno
func ejemploVariablesEntornoAvanzado() {
	fmt.Println("=== MANEJO AVANZADO DE VARIABLES DE ENTORNO ===")

	// 1. Listar todas las variables de entorno
	listarVariablesEntorno()

	// 2. Trabajar con PATH
	analizarPATH()

	// 3. Crear configuración desde variables de entorno
	crearConfiguracionDesdeEnv()

	// 4. Variables de entorno temporales
	variablesTemporales()
}

func listarVariablesEntorno() {
	fmt.Println("\n1. LISTADO DE VARIABLES DE ENTORNO:")

	env := os.Environ()
	fmt.Printf("Total de variables de entorno: %d\n", len(env))

	// Agrupar por prefijo
	grupos := make(map[string][]string)
	for _, envVar := range env {
		parts := strings.SplitN(envVar, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			// Agrupar por prefijo común
			if strings.HasPrefix(key, "GO") {
				grupos["GO"] = append(grupos["GO"], envVar)
			} else if strings.HasPrefix(key, "PATH") || strings.HasPrefix(key, "HOME") || strings.HasPrefix(key, "USER") {
				grupos["SISTEMA"] = append(grupos["SISTEMA"], envVar)
			} else if len(grupos["OTRAS"]) < 5 { // Solo mostrar las primeras 5 otras
				grupos["OTRAS"] = append(grupos["OTRAS"], envVar)
			}
		}
	}

	// Mostrar grupos
	for grupo, vars := range grupos {
		fmt.Printf("\n%s:\n", grupo)
		sort.Strings(vars)
		for _, v := range vars {
			parts := strings.SplitN(v, "=", 2)
			if len(parts) == 2 {
				if len(parts[1]) > 60 {
					fmt.Printf("  %s=%s...\n", parts[0], parts[1][:60])
				} else {
					fmt.Printf("  %s=%s\n", parts[0], parts[1])
				}
			}
		}
	}
}

func analizarPATH() {
	fmt.Println("\n2. ANÁLISIS DE LA VARIABLE PATH:")

	pathVar := os.Getenv("PATH")
	if pathVar == "" {
		fmt.Println("PATH no está definido")
		return
	}

	// Dividir PATH por el separador del sistema
	var separator string
	if strings.Contains(strings.ToLower(os.Getenv("OS")), "windows") {
		separator = ";"
	} else {
		separator = ":"
	}

	paths := strings.Split(pathVar, separator)
	fmt.Printf("Directorios en PATH (%d):\n", len(paths))

	for i, path := range paths {
		if i >= 10 { // Solo mostrar los primeros 10
			fmt.Printf("  ... y %d más\n", len(paths)-10)
			break
		}

		// Verificar si el directorio existe
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			fmt.Printf("  ✅ %s\n", path)
		} else {
			fmt.Printf("  ❌ %s (no existe)\n", path)
		}
	}
}

func crearConfiguracionDesdeEnv() {
	fmt.Println("\n3. CONFIGURACIÓN DESDE VARIABLES DE ENTORNO:")

	// Definir configuración por defecto
	config := map[string]string{
		"APP_NAME":     "MiAplicacion",
		"APP_VERSION":  "1.0.0",
		"DEBUG":        "false",
		"PORT":         "8080",
		"DATABASE_URL": "localhost:5432",
	}

	// Establecer variables de entorno de ejemplo
	os.Setenv("APP_NAME", "EjemploOS")
	os.Setenv("DEBUG", "true")
	os.Setenv("PORT", "3000")

	fmt.Println("Configuración final:")
	for key, defaultValue := range config {
		// Usar variable de entorno si existe, sino usar valor por defecto
		if envValue := os.Getenv(key); envValue != "" {
			fmt.Printf("  %s=%s (desde env)\n", key, envValue)
		} else {
			fmt.Printf("  %s=%s (por defecto)\n", key, defaultValue)
		}
	}

	// Limpiar variables de ejemplo
	os.Unsetenv("APP_NAME")
	os.Unsetenv("DEBUG")
	os.Unsetenv("PORT")
}

func variablesTemporales() {
	fmt.Println("\n4. VARIABLES TEMPORALES:")

	// Crear variables temporales
	tempVars := map[string]string{
		"TEMP_VAR_1":  "valor_temporal_1",
		"TEMP_VAR_2":  "valor_temporal_2",
		"TEMP_CONFIG": "configuracion_temporal",
	}

	fmt.Println("Estableciendo variables temporales:")
	for key, value := range tempVars {
		os.Setenv(key, value)
		fmt.Printf("  %s=%s\n", key, value)
	}

	// Verificar que existen
	fmt.Println("\nVerificando variables:")
	for key := range tempVars {
		if value, exists := os.LookupEnv(key); exists {
			fmt.Printf("  ✅ %s=%s\n", key, value)
		} else {
			fmt.Printf("  ❌ %s no encontrado\n", key)
		}
	}

	// Limpiar variables temporales
	fmt.Println("\nLimpiando variables temporales:")
	for key := range tempVars {
		os.Unsetenv(key)
		fmt.Printf("  🗑️  %s eliminado\n", key)
	}

	// Verificar que fueron eliminadas
	fmt.Println("\nVerificando eliminación:")
	for key := range tempVars {
		if _, exists := os.LookupEnv(key); !exists {
			fmt.Printf("  ✅ %s eliminado correctamente\n", key)
		} else {
			fmt.Printf("  ❌ %s aún existe\n", key)
		}
	}
}
