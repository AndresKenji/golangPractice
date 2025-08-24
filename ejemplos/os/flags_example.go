package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func runFlagsExample() {
	// Definir flags personalizados
	var (
		name    = flag.String("name", "Usuario", "Nombre del usuario")
		age     = flag.Int("age", 0, "Edad del usuario")
		verbose = flag.Bool("verbose", false, "Mostrar información detallada")
		emails  = flag.String("emails", "", "Lista de emails separados por comas")
		output  = flag.String("output", "output.txt", "Archivo de salida")
		help    = flag.Bool("help", false, "Mostrar ayuda")
	)

	// Personalizar el mensaje de uso
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Uso del programa: %s [opciones]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Opciones disponibles:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nEjemplos de uso:\n")
		fmt.Fprintf(os.Stderr, "  %s -name=\"Juan Pérez\" -age=30 -verbose\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -emails=\"juan@email.com,maria@email.com\" -output=\"resultados.txt\"\n", os.Args[0])
	}

	// Parsear los flags
	flag.Parse()

	// Mostrar ayuda si se solicita
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Mostrar información según los flags
	fmt.Println("=== EJEMPLO DE FLAGS AVANZADO ===")
	fmt.Printf("Nombre: %s\n", *name)

	if *age > 0 {
		fmt.Printf("Edad: %d años\n", *age)
		if *age >= 18 {
			fmt.Println("Estatus: Mayor de edad")
		} else {
			fmt.Println("Estatus: Menor de edad")
		}
	}

	if *emails != "" {
		emailList := strings.Split(*emails, ",")
		fmt.Printf("Emails (%d):\n", len(emailList))
		for i, email := range emailList {
			fmt.Printf("  %d. %s\n", i+1, strings.TrimSpace(email))
		}
	}

	fmt.Printf("Archivo de salida: %s\n", *output)

	if *verbose {
		fmt.Println("\n=== INFORMACIÓN DETALLADA ===")
		fmt.Printf("Programa ejecutado: %s\n", os.Args[0])
		fmt.Printf("Argumentos totales: %d\n", len(os.Args))
		fmt.Printf("Flags parseados: %d\n", flag.NFlag())

		// Mostrar variables de entorno relevantes
		fmt.Println("Variables de entorno importantes:")
		envVars := []string{"USER", "USERNAME", "HOME", "USERPROFILE", "PATH"}
		for _, envVar := range envVars {
			if value := os.Getenv(envVar); value != "" {
				if envVar == "PATH" {
					fmt.Printf("  %s: %s...\n", envVar, value[:50])
				} else {
					fmt.Printf("  %s: %s\n", envVar, value)
				}
			}
		}
	}

	// Argumentos no parseados
	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) > 0 {
		fmt.Printf("\nArgumentos adicionales: %v\n", nonFlagArgs)
	}

	// Crear archivo de salida si no existe
	if _, err := os.Stat(*output); os.IsNotExist(err) {
		content := fmt.Sprintf("Archivo creado para: %s\nFecha: %s\n", *name, "2025-08-23")
		err := os.WriteFile(*output, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error creando archivo de salida: %v\n", err)
		} else {
			fmt.Printf("Archivo '%s' creado exitosamente\n", *output)
		}
	}

	fmt.Println("\nPara ver todas las opciones, ejecuta: go run flags_example.go -help")
}
