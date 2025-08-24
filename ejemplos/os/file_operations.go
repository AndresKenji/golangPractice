package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Función para demostrar operaciones avanzadas con archivos
func ejemploArchivosAvanzado() {
	fmt.Println("=== OPERACIONES AVANZADAS CON ARCHIVOS ===")

	// Crear directorio temporal para las pruebas
	tempDir, err := os.MkdirTemp("", "archivos_avanzado_")
	if err != nil {
		fmt.Printf("Error creando directorio temporal: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir) // Limpiar al final

	fmt.Printf("Directorio temporal creado: %s\n", tempDir)

	// 1. Crear múltiples archivos
	crearArchivosEjemplo(tempDir)

	// 2. Buscar archivos por patrón
	buscarArchivosPorPatron(tempDir)

	// 3. Recorrer directorio recursivamente
	recorrerDirectorioRecursivo(tempDir)

	// 4. Monitorear cambios en archivo (simulado)
	monitorearCambiosArchivo(tempDir)

	// 5. Trabajar con archivos grandes
	trabajarConArchivosGrandes(tempDir)

	// 6. Operaciones con metadatos
	operacionesMetadatos(tempDir)
}

func crearArchivosEjemplo(baseDir string) {
	fmt.Println("\n1. CREANDO ARCHIVOS DE EJEMPLO:")

	// Crear estructura de directorios
	dirs := []string{
		"documentos",
		"imagenes",
		"scripts",
		"logs",
	}

	for _, dir := range dirs {
		dirPath := filepath.Join(baseDir, dir)
		os.Mkdir(dirPath, 0755)
		fmt.Printf("  Directorio creado: %s\n", dir)
	}

	// Crear archivos de ejemplo
	archivos := map[string]string{
		"documentos/readme.txt": "Este es un archivo README\nContiene información importante\n",
		"documentos/notas.md":   "# Notas\n\n- Punto 1\n- Punto 2\n",
		"imagenes/imagen1.jpg":  "Datos simulados de imagen JPG",
		"imagenes/imagen2.png":  "Datos simulados de imagen PNG",
		"scripts/script.sh":     "#!/bin/bash\necho 'Hola mundo'\n",
		"scripts/script.bat":    "@echo off\necho Hola mundo\n",
		"logs/app.log":          "2025-08-23 INFO: Aplicación iniciada\n2025-08-23 DEBUG: Procesando datos\n",
		"logs/error.log":        "2025-08-23 ERROR: Error simulado\n",
	}

	for archivo, contenido := range archivos {
		archivoPath := filepath.Join(baseDir, archivo)
		err := os.WriteFile(archivoPath, []byte(contenido), 0644)
		if err != nil {
			fmt.Printf("  Error creando %s: %v\n", archivo, err)
		} else {
			fmt.Printf("  Archivo creado: %s\n", archivo)
		}
	}
}

func buscarArchivosPorPatron(baseDir string) {
	fmt.Println("\n2. BÚSQUEDA DE ARCHIVOS POR PATRÓN:")

	// Buscar todos los archivos .txt
	matches, err := filepath.Glob(filepath.Join(baseDir, "*", "*.txt"))
	if err != nil {
		fmt.Printf("Error buscando archivos: %v\n", err)
		return
	}

	fmt.Println("Archivos .txt encontrados:")
	for _, match := range matches {
		relPath, _ := filepath.Rel(baseDir, match)
		fmt.Printf("  %s\n", relPath)
	}

	// Buscar archivos de log
	logMatches, _ := filepath.Glob(filepath.Join(baseDir, "logs", "*.log"))
	fmt.Println("Archivos .log encontrados:")
	for _, match := range logMatches {
		relPath, _ := filepath.Rel(baseDir, match)
		info, _ := os.Stat(match)
		fmt.Printf("  %s (%d bytes)\n", relPath, info.Size())
	}
}

func recorrerDirectorioRecursivo(baseDir string) {
	fmt.Println("\n3. RECORRIDO RECURSIVO DE DIRECTORIO:")

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(baseDir, path)
		if relPath == "." {
			return nil
		}

		indent := ""
		depth := len(filepath.SplitList(strings.ReplaceAll(relPath, string(filepath.Separator), string(filepath.ListSeparator))))
		for i := 0; i < depth-1; i++ {
			indent += "  "
		}

		if info.IsDir() {
			fmt.Printf("%s📁 %s/\n", indent, info.Name())
		} else {
			fmt.Printf("%s📄 %s (%d bytes)\n", indent, info.Name(), info.Size())
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error recorriendo directorio: %v\n", err)
	}
}

func monitorearCambiosArchivo(baseDir string) {
	fmt.Println("\n4. MONITOREO DE CAMBIOS (SIMULADO):")

	archivoTest := filepath.Join(baseDir, "test_monitoreo.txt")

	// Crear archivo inicial
	contenidoInicial := "Contenido inicial\n"
	os.WriteFile(archivoTest, []byte(contenidoInicial), 0644)

	// Obtener información inicial
	infoInicial, _ := os.Stat(archivoTest)
	fmt.Printf("Archivo creado: %s\n", archivoTest)
	fmt.Printf("Tamaño inicial: %d bytes\n", infoInicial.Size())
	fmt.Printf("Modificado: %s\n", infoInicial.ModTime().Format("2006-01-02 15:04:05"))

	// Simular cambio después de un tiempo
	time.Sleep(100 * time.Millisecond)

	// Modificar archivo
	file, _ := os.OpenFile(archivoTest, os.O_APPEND|os.O_WRONLY, 0644)
	file.WriteString("Línea agregada\n")
	file.Close()

	// Verificar cambios
	infoModificada, _ := os.Stat(archivoTest)
	fmt.Printf("Después de modificación:\n")
	fmt.Printf("Nuevo tamaño: %d bytes\n", infoModificada.Size())
	fmt.Printf("Modificado: %s\n", infoModificada.ModTime().Format("2006-01-02 15:04:05"))

	if infoModificada.ModTime().After(infoInicial.ModTime()) {
		fmt.Println("✅ Cambio detectado!")
	}
}

func trabajarConArchivosGrandes(baseDir string) {
	fmt.Println("\n5. TRABAJO CON ARCHIVOS GRANDES (SIMULADO):")

	archivoGrande := filepath.Join(baseDir, "archivo_grande.txt")

	// Crear archivo "grande" con múltiples líneas
	file, err := os.Create(archivoGrande)
	if err != nil {
		fmt.Printf("Error creando archivo grande: %v\n", err)
		return
	}
	defer file.Close()

	// Escribir datos línea por línea
	writer := bufio.NewWriter(file)
	for i := 1; i <= 1000; i++ {
		line := fmt.Sprintf("Línea número %04d - Datos de ejemplo para simular archivo grande\n", i)
		writer.WriteString(line)
	}
	writer.Flush()

	fmt.Printf("Archivo grande creado: %s\n", archivoGrande)

	// Leer archivo línea por línea (eficiente para archivos grandes)
	readFile, _ := os.Open(archivoGrande)
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		// Solo mostrar las primeras 5 líneas
		if lineCount <= 5 {
			fmt.Printf("  %s\n", scanner.Text())
		}
	}

	fmt.Printf("Total de líneas procesadas: %d\n", lineCount)

	// Obtener información del archivo
	info, _ := os.Stat(archivoGrande)
	fmt.Printf("Tamaño del archivo: %d bytes\n", info.Size())
}

func operacionesMetadatos(baseDir string) {
	fmt.Println("\n6. OPERACIONES CON METADATOS:")

	archivo := filepath.Join(baseDir, "metadatos_test.txt")
	contenido := "Archivo para probar metadatos\n"
	os.WriteFile(archivo, []byte(contenido), 0644)

	// Obtener información detallada
	info, err := os.Stat(archivo)
	if err != nil {
		fmt.Printf("Error obteniendo información: %v\n", err)
		return
	}

	fmt.Printf("Información detallada de '%s':\n", info.Name())
	fmt.Printf("  Tamaño: %d bytes\n", info.Size())
	fmt.Printf("  Modo: %v\n", info.Mode())
	fmt.Printf("  Modificado: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
	fmt.Printf("  Es directorio: %v\n", info.IsDir())
	fmt.Printf("  Es archivo regular: %v\n", info.Mode().IsRegular())

	// Cambiar tiempo de modificación
	nuevoTiempo := time.Now().Add(-24 * time.Hour) // Un día atrás
	err = os.Chtimes(archivo, nuevoTiempo, nuevoTiempo)
	if err != nil {
		fmt.Printf("Error cambiando tiempo: %v\n", err)
	} else {
		infoActualizada, _ := os.Stat(archivo)
		fmt.Printf("  Tiempo modificado a: %s\n", infoActualizada.ModTime().Format("2006-01-02 15:04:05"))
	}

	// Crear enlace simbólico (si es soportado)
	enlace := filepath.Join(baseDir, "enlace_test.txt")
	err = os.Symlink(archivo, enlace)
	if err != nil {
		fmt.Printf("Enlace simbólico no soportado o error: %v\n", err)
	} else {
		fmt.Printf("Enlace simbólico creado: %s -> %s\n", enlace, archivo)

		// Verificar si es enlace simbólico
		linkInfo, _ := os.Lstat(enlace)
		if linkInfo.Mode()&os.ModeSymlink != 0 {
			fmt.Println("  ✅ Es un enlace simbólico")

			// Leer destino del enlace
			target, err := os.Readlink(enlace)
			if err == nil {
				fmt.Printf("  Apunta a: %s\n", target)
			}
		}
	}
}
