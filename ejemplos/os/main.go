package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	// Flags de línea de comandos
	verbose    = flag.Bool("verbose", false, "Activar modo verbose")
	fileName   = flag.String("file", "test.txt", "Nombre del archivo a procesar")
	outputDir  = flag.String("output", "output", "Directorio de salida")
	iterations = flag.Int("iterations", 1, "Número de iteraciones")

	// Flags para ejecutar ejemplos específicos
	all           = flag.Bool("all", true, "Ejecutar todos los ejemplos básicos (por defecto)")
	system        = flag.Bool("system", false, "Solo ejemplos de información del sistema")
	env           = flag.Bool("env", false, "Solo ejemplos de variables de entorno")
	envAdvanced   = flag.Bool("env-advanced", false, "Ejemplos avanzados de variables de entorno")
	files         = flag.Bool("files", false, "Solo ejemplos básicos de archivos")
	filesAdvanced = flag.Bool("files-advanced", false, "Ejemplos avanzados de archivos")
	dirs          = flag.Bool("dirs", false, "Solo ejemplos de directorios")
	paths         = flag.Bool("paths", false, "Solo ejemplos de rutas")
	commands      = flag.Bool("commands", false, "Solo ejemplos de comandos del sistema")
	flags         = flag.Bool("flags", false, "Solo ejemplos de flags")
	flagsAdvanced = flag.Bool("flags-advanced", false, "Ejemplos avanzados de flags")
	help          = flag.Bool("help", false, "Mostrar ayuda detallada")
)

func main() {
	// Personalizar el mensaje de uso
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Ejemplos del módulo OS en Go\n\n")
		fmt.Fprintf(os.Stderr, "Uso: %s [opciones]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Opciones generales:\n")
		fmt.Fprintf(os.Stderr, "  -verbose\t\tActivar modo verbose\n")
		fmt.Fprintf(os.Stderr, "  -file string\t\tNombre del archivo a procesar (default \"test.txt\")\n")
		fmt.Fprintf(os.Stderr, "  -output string\tDirectorio de salida (default \"output\")\n")
		fmt.Fprintf(os.Stderr, "  -iterations int\tNúmero de iteraciones (default 1)\n\n")

		fmt.Fprintf(os.Stderr, "Ejemplos específicos (usar solo uno):\n")
		fmt.Fprintf(os.Stderr, "  -all\t\t\tEjecutar todos los ejemplos básicos (default)\n")
		fmt.Fprintf(os.Stderr, "  -system\t\tSolo información del sistema\n")
		fmt.Fprintf(os.Stderr, "  -env\t\t\tSolo variables de entorno básicas\n")
		fmt.Fprintf(os.Stderr, "  -env-advanced\t\tVariables de entorno avanzadas\n")
		fmt.Fprintf(os.Stderr, "  -files\t\tSolo archivos básicos\n")
		fmt.Fprintf(os.Stderr, "  -files-advanced\tArchivos avanzados\n")
		fmt.Fprintf(os.Stderr, "  -dirs\t\t\tSolo directorios\n")
		fmt.Fprintf(os.Stderr, "  -paths\t\tSolo rutas\n")
		fmt.Fprintf(os.Stderr, "  -commands\t\tSolo comandos del sistema\n")
		fmt.Fprintf(os.Stderr, "  -flags\t\tSolo ejemplos de flags básicos\n")
		fmt.Fprintf(os.Stderr, "  -flags-advanced\tEjemplos de flags avanzados\n")
		fmt.Fprintf(os.Stderr, "  -help\t\t\tMostrar esta ayuda\n\n")

		fmt.Fprintf(os.Stderr, "Ejemplos de uso:\n")
		fmt.Fprintf(os.Stderr, "  %s                    # Ejecutar todos los ejemplos básicos\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -system           # Solo información del sistema\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -files -verbose   # Solo archivos con modo verbose\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -env-advanced     # Variables de entorno avanzadas\n", os.Args[0])
	}

	// Parsear flags de línea de comandos
	flag.Parse()

	// Mostrar ayuda si se solicita
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Determinar qué ejemplos ejecutar
	runSpecificExamples := *system || *env || *envAdvanced || *files || *filesAdvanced ||
		*dirs || *paths || *commands || *flags || *flagsAdvanced

	fmt.Println("=== EJEMPLOS DEL MÓDULO OS EN GO ===")
	if *verbose {
		fmt.Printf("Modo verbose activado\n")
		fmt.Printf("Archivo: %s, Output: %s, Iteraciones: %d\n", *fileName, *outputDir, *iterations)
	}
	fmt.Println()

	// Ejecutar ejemplos según los flags
	if *system || (!runSpecificExamples && *all) {
		ejemploInformacionSistema()
	}

	if *env || (!runSpecificExamples && *all) {
		ejemploVariablesEntorno()
	}

	if *envAdvanced {
		fmt.Println("=== EJEMPLOS AVANZADOS DE VARIABLES DE ENTORNO ===")
		fmt.Println("Para ejecutar estos ejemplos, use uno de estos comandos:")
		fmt.Println("  go run env_operations.go main.go")
		fmt.Println("  go build -o env_examples env_operations.go main.go && ./env_examples")
		fmt.Println()
		fmt.Println("O descomente la función main() en env_operations.go y ejecute:")
		fmt.Println("  go run env_operations.go")
		fmt.Println()

		// Ejecutar una versión simplificada aquí
		fmt.Println("Vista previa (versión simplificada):")
		ejemploVariablesEntornoSimplificado()
	}

	if *flags || (!runSpecificExamples && *all) {
		ejemploArgumentosYFlags()
	}

	if *flagsAdvanced {
		fmt.Println("=== EJEMPLOS AVANZADOS DE FLAGS ===")
		fmt.Println("Para ejecutar estos ejemplos, use uno de estos comandos:")
		fmt.Println("  go run flags_example.go main.go")
		fmt.Println("  go build -o flags_examples flags_example.go main.go && ./flags_examples")
		fmt.Println()
		fmt.Println("O descomente la función main() en flags_example.go y ejecute:")
		fmt.Println("  go run flags_example.go")
		fmt.Println()

		// Ejecutar una versión simplificada aquí
		fmt.Println("Vista previa (versión simplificada):")
		ejemploFlagsSimplificado()
	}

	if *files || (!runSpecificExamples && *all) {
		ejemploArchivosBasicos()
		ejemploInformacionArchivos()
		ejemploPermisos()
		ejemploLecturaEscritura()
	}

	if *filesAdvanced {
		fmt.Println("=== EJEMPLOS AVANZADOS DE ARCHIVOS ===")
		fmt.Println("Para ejecutar estos ejemplos, use uno de estos comandos:")
		fmt.Println("  go run file_operations.go main.go")
		fmt.Println("  go build -o file_examples file_operations.go main.go && ./file_examples")
		fmt.Println()
		fmt.Println("O descomente la función main() en file_operations.go y ejecute:")
		fmt.Println("  go run file_operations.go")
		fmt.Println()

		// Ejecutar una versión simplificada aquí
		fmt.Println("Vista previa (versión simplificada):")
		ejemploArchivosSimplificado()
	}

	if *dirs || (!runSpecificExamples && *all) {
		ejemploDirectorios()
	}

	if *paths || (!runSpecificExamples && *all) {
		ejemploRutas()
	}

	if *commands || (!runSpecificExamples && *all) {
		ejemploComandosSistema()
		ejemploSeñales()
		ejemploEntradaSalida()
	}

	// Solo mostrar mensaje final si se ejecutó algo
	if runSpecificExamples || *all {
		fmt.Println("=== FIN DE EJEMPLOS ===")
	} else {
		fmt.Println("No se especificó ningún ejemplo. Use -help para ver las opciones disponibles.")
	}
}

// 1. Información del sistema y proceso
func ejemploInformacionSistema() {
	fmt.Println("1. INFORMACIÓN DEL SISTEMA Y PROCESO:")

	// Información del proceso actual
	fmt.Printf("PID del proceso actual: %d\n", os.Getpid())
	fmt.Printf("PID del proceso padre: %d\n", os.Getppid())

	// Información del usuario
	fmt.Printf("UID del usuario: %d\n", os.Getuid())
	fmt.Printf("GID del grupo: %d\n", os.Getgid())

	// Directorio de trabajo actual
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error obteniendo directorio actual: %v\n", err)
	} else {
		fmt.Printf("Directorio de trabajo actual: %s\n", pwd)
	}

	// Directorio home del usuario
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error obteniendo directorio home: %v\n", err)
	} else {
		fmt.Printf("Directorio home: %s\n", homeDir)
	}

	// Directorio temporal
	tempDir := os.TempDir()
	fmt.Printf("Directorio temporal: %s\n", tempDir)

	// Hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error obteniendo hostname: %v\n", err)
	} else {
		fmt.Printf("Hostname: %s\n", hostname)
	}

	fmt.Println()
}

// 2. Variables de entorno
func ejemploVariablesEntorno() {
	fmt.Println("2. VARIABLES DE ENTORNO:")

	// Obtener variable de entorno específica
	path := os.Getenv("PATH")
	fmt.Printf("PATH: %s\n", path[:100]+"...") // Solo mostrar los primeros 100 caracteres

	// Obtener variable con valor por defecto
	customVar := os.Getenv("MI_VARIABLE_CUSTOM")
	if customVar == "" {
		customVar = "valor_por_defecto"
	}
	fmt.Printf("MI_VARIABLE_CUSTOM: %s\n", customVar)

	// Establecer variable de entorno
	os.Setenv("MI_NUEVA_VARIABLE", "mi_valor")
	fmt.Printf("MI_NUEVA_VARIABLE: %s\n", os.Getenv("MI_NUEVA_VARIABLE"))

	// Verificar si existe una variable
	if value, exists := os.LookupEnv("HOME"); exists {
		fmt.Printf("HOME existe: %s\n", value)
	} else {
		fmt.Println("HOME no existe")
	}

	// Listar todas las variables de entorno (solo las primeras 5)
	fmt.Println("Primeras 5 variables de entorno:")
	allEnv := os.Environ()
	for i, env := range allEnv {
		if i >= 5 {
			break
		}
		fmt.Printf("  %s\n", env)
	}

	// Limpiar variable de entorno
	os.Unsetenv("MI_NUEVA_VARIABLE")
	fmt.Printf("MI_NUEVA_VARIABLE después de limpiar: '%s'\n", os.Getenv("MI_NUEVA_VARIABLE"))

	fmt.Println()
}

// 3. Argumentos de línea de comandos y flags
func ejemploArgumentosYFlags() {
	fmt.Println("3. ARGUMENTOS DE LÍNEA DE COMANDOS Y FLAGS:")

	// Argumentos de línea de comandos
	fmt.Printf("Nombre del programa: %s\n", os.Args[0])
	fmt.Printf("Número total de argumentos: %d\n", len(os.Args))

	if len(os.Args) > 1 {
		fmt.Println("Argumentos:")
		for i, arg := range os.Args[1:] {
			fmt.Printf("  Arg[%d]: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("No se proporcionaron argumentos adicionales")
	}

	// Flags parseados
	fmt.Printf("Flag verbose: %v\n", *verbose)
	fmt.Printf("Flag file: %s\n", *fileName)
	fmt.Printf("Flag output: %s\n", *outputDir)
	fmt.Printf("Flag iterations: %d\n", *iterations)

	// Argumentos no parseados por flags
	nonFlagArgs := flag.Args()
	if len(nonFlagArgs) > 0 {
		fmt.Println("Argumentos no parseados por flags:")
		for i, arg := range nonFlagArgs {
			fmt.Printf("  NonFlag[%d]: %s\n", i, arg)
		}
	}

	fmt.Println()
}

// 4. Trabajar con archivos básicos
func ejemploArchivosBasicos() {
	fmt.Println("4. TRABAJAR CON ARCHIVOS BÁSICOS:")

	// Crear un archivo
	fileName := "ejemplo.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creando archivo: %v\n", err)
		return
	}
	defer file.Close()

	// Escribir al archivo
	content := "Este es un archivo de ejemplo\nCreado con Go\n"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error escribiendo al archivo: %v\n", err)
		return
	}

	fmt.Printf("Archivo '%s' creado exitosamente\n", fileName)

	// Verificar si el archivo existe
	if _, err := os.Stat(fileName); err == nil {
		fmt.Printf("El archivo '%s' existe\n", fileName)
	} else if os.IsNotExist(err) {
		fmt.Printf("El archivo '%s' no existe\n", fileName)
	} else {
		fmt.Printf("Error verificando archivo: %v\n", err)
	}

	// Renombrar archivo
	newFileName := "ejemplo_renombrado.txt"
	err = os.Rename(fileName, newFileName)
	if err != nil {
		fmt.Printf("Error renombrando archivo: %v\n", err)
	} else {
		fmt.Printf("Archivo renombrado de '%s' a '%s'\n", fileName, newFileName)
	}

	// Copiar archivo (implementación manual)
	copiedFileName := "ejemplo_copia.txt"
	err = copyFile(newFileName, copiedFileName)
	if err != nil {
		fmt.Printf("Error copiando archivo: %v\n", err)
	} else {
		fmt.Printf("Archivo copiado a '%s'\n", copiedFileName)
	}

	// Leer archivo completo
	data, err := os.ReadFile(newFileName)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v\n", err)
	} else {
		fmt.Printf("Contenido del archivo:\n%s", string(data))
	}

	// Limpiar archivos creados
	filesToClean := []string{newFileName, copiedFileName}
	for _, file := range filesToClean {
		if err := os.Remove(file); err != nil {
			fmt.Printf("Error eliminando '%s': %v\n", file, err)
		} else {
			fmt.Printf("Archivo '%s' eliminado\n", file)
		}
	}

	fmt.Println()
}

// 5. Directorios y navegación
func ejemploDirectorios() {
	fmt.Println("5. DIRECTORIOS Y NAVEGACIÓN:")

	// Crear directorio
	dirName := "ejemplo_directorio"
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Printf("Error creando directorio: %v\n", err)
	} else {
		fmt.Printf("Directorio '%s' creado\n", dirName)
	}

	// Crear directorios anidados
	nestedDir := filepath.Join(dirName, "subdir1", "subdir2")
	err = os.MkdirAll(nestedDir, 0755)
	if err != nil {
		fmt.Printf("Error creando directorios anidados: %v\n", err)
	} else {
		fmt.Printf("Directorios anidados creados: %s\n", nestedDir)
	}

	// Cambiar directorio de trabajo
	originalDir, _ := os.Getwd()
	err = os.Chdir(dirName)
	if err != nil {
		fmt.Printf("Error cambiando directorio: %v\n", err)
	} else {
		currentDir, _ := os.Getwd()
		fmt.Printf("Cambiado al directorio: %s\n", currentDir)
	}

	// Volver al directorio original
	os.Chdir(originalDir)
	fmt.Printf("Vuelto al directorio original: %s\n", originalDir)

	// Listar contenido del directorio
	entries, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Printf("Error leyendo directorio: %v\n", err)
	} else {
		fmt.Printf("Contenido del directorio '%s':\n", dirName)
		for _, entry := range entries {
			if entry.IsDir() {
				fmt.Printf("  [DIR]  %s\n", entry.Name())
			} else {
				fmt.Printf("  [FILE] %s\n", entry.Name())
			}
		}
	}

	// Limpiar directorios creados
	err = os.RemoveAll(dirName)
	if err != nil {
		fmt.Printf("Error eliminando directorio: %v\n", err)
	} else {
		fmt.Printf("Directorio '%s' y todo su contenido eliminado\n", dirName)
	}

	fmt.Println()
}

// 6. Información de archivos y directorios
func ejemploInformacionArchivos() {
	fmt.Println("6. INFORMACIÓN DE ARCHIVOS Y DIRECTORIOS:")

	// Crear archivo temporal para obtener información
	tempFile, err := os.CreateTemp("", "info_ejemplo_*.txt")
	if err != nil {
		fmt.Printf("Error creando archivo temporal: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Escribir algo al archivo
	tempFile.WriteString("Contenido de ejemplo para obtener información")

	// Obtener información del archivo
	fileInfo, err := os.Stat(tempFile.Name())
	if err != nil {
		fmt.Printf("Error obteniendo información del archivo: %v\n", err)
		return
	}

	fmt.Printf("Información del archivo '%s':\n", fileInfo.Name())
	fmt.Printf("  Tamaño: %d bytes\n", fileInfo.Size())
	fmt.Printf("  Modo: %v\n", fileInfo.Mode())
	fmt.Printf("  Tiempo de modificación: %v\n", fileInfo.ModTime())
	fmt.Printf("  Es directorio: %v\n", fileInfo.IsDir())

	// Verificar permisos específicos
	mode := fileInfo.Mode()
	fmt.Printf("  Permisos específicos:\n")
	fmt.Printf("    Lectura propietario: %v\n", mode&0400 != 0)
	fmt.Printf("    Escritura propietario: %v\n", mode&0200 != 0)
	fmt.Printf("    Ejecución propietario: %v\n", mode&0100 != 0)

	// Obtener información del directorio actual
	currentDir, _ := os.Getwd()
	dirInfo, err := os.Stat(currentDir)
	if err != nil {
		fmt.Printf("Error obteniendo información del directorio: %v\n", err)
	} else {
		fmt.Printf("Información del directorio actual:\n")
		fmt.Printf("  Nombre: %s\n", dirInfo.Name())
		fmt.Printf("  Es directorio: %v\n", dirInfo.IsDir())
		fmt.Printf("  Modo: %v\n", dirInfo.Mode())
	}

	fmt.Println()
}

// 7. Permisos de archivos
func ejemploPermisos() {
	fmt.Println("7. PERMISOS DE ARCHIVOS:")

	// Crear archivo para cambiar permisos
	fileName := "permisos_ejemplo.txt"
	err := os.WriteFile(fileName, []byte("Archivo para probar permisos"), 0644)
	if err != nil {
		fmt.Printf("Error creando archivo: %v\n", err)
		return
	}
	defer os.Remove(fileName)

	// Obtener permisos actuales
	fileInfo, _ := os.Stat(fileName)
	fmt.Printf("Permisos actuales: %v\n", fileInfo.Mode())

	// Cambiar permisos del archivo
	err = os.Chmod(fileName, 0755)
	if err != nil {
		fmt.Printf("Error cambiando permisos: %v\n", err)
	} else {
		fileInfo, _ = os.Stat(fileName)
		fmt.Printf("Nuevos permisos: %v\n", fileInfo.Mode())
	}

	// Cambiar propietario (solo funciona con permisos adecuados)
	// err = os.Chown(fileName, os.Getuid(), os.Getgid())
	// if err != nil {
	// 	fmt.Printf("Error cambiando propietario: %v\n", err)
	// } else {
	// 	fmt.Println("Propietario cambiado exitosamente")
	// }

	fmt.Println()
}

// 8. Lectura y escritura de archivos
func ejemploLecturaEscritura() {
	fmt.Println("8. LECTURA Y ESCRITURA DE ARCHIVOS:")

	fileName := "lectura_escritura.txt"

	// Escribir archivo completo de una vez
	content := "Primera línea\nSegunda línea\nTercera línea\n"
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error escribiendo archivo: %v\n", err)
		return
	}
	defer os.Remove(fileName)

	fmt.Printf("Archivo '%s' escrito exitosamente\n", fileName)

	// Leer archivo completo
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v\n", err)
		return
	}
	fmt.Printf("Contenido leído:\n%s", string(data))

	// Escribir al archivo línea por línea
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error abriendo archivo para append: %v\n", err)
		return
	}
	defer file.Close()

	// Escribir líneas adicionales
	additionalLines := []string{
		"Cuarta línea agregada\n",
		"Quinta línea agregada\n",
	}

	for _, line := range additionalLines {
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Printf("Error escribiendo línea: %v\n", err)
			return
		}
	}

	fmt.Println("Líneas adicionales agregadas")

	// Leer archivo línea por línea
	fmt.Println("Leyendo archivo línea por línea:")
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error abriendo archivo para lectura: %v\n", err)
		return
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)
	lineNumber := 1
	for scanner.Scan() {
		fmt.Printf("  Línea %d: %s\n", lineNumber, scanner.Text())
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error leyendo archivo: %v\n", err)
	}

	fmt.Println()
}

// 9. Trabajar con rutas
func ejemploRutas() {
	fmt.Println("9. TRABAJAR CON RUTAS:")

	// Ruta de ejemplo
	examplePath := filepath.Join("home", "usuario", "documentos", "archivo.txt")
	fmt.Printf("Ruta construida: %s\n", examplePath)

	// Separar componentes de la ruta
	dir := filepath.Dir(examplePath)
	base := filepath.Base(examplePath)
	ext := filepath.Ext(examplePath)

	fmt.Printf("Directorio: %s\n", dir)
	fmt.Printf("Nombre base: %s\n", base)
	fmt.Printf("Extensión: %s\n", ext)

	// Ruta absoluta
	absPath, err := filepath.Abs(examplePath)
	if err != nil {
		fmt.Printf("Error obteniendo ruta absoluta: %v\n", err)
	} else {
		fmt.Printf("Ruta absoluta: %s\n", absPath)
	}

	// Limpiar ruta
	messyPath := "home//usuario/../usuario/./documentos/archivo.txt"
	cleanPath := filepath.Clean(messyPath)
	fmt.Printf("Ruta desordenada: %s\n", messyPath)
	fmt.Printf("Ruta limpia: %s\n", cleanPath)

	// Verificar si la ruta es absoluta
	fmt.Printf("¿Es absoluta '%s'? %v\n", examplePath, filepath.IsAbs(examplePath))
	fmt.Printf("¿Es absoluta '%s'? %v\n", absPath, filepath.IsAbs(absPath))

	// Dividir ruta en directorio y archivo
	splitDir, splitFile := filepath.Split(examplePath)
	fmt.Printf("División - Dir: '%s', Archivo: '%s'\n", splitDir, splitFile)

	// Buscar archivos con patrón
	fmt.Println("Archivos .go en el directorio actual:")
	matches, err := filepath.Glob("*.go")
	if err != nil {
		fmt.Printf("Error buscando archivos: %v\n", err)
	} else {
		for _, match := range matches {
			fmt.Printf("  %s\n", match)
		}
	}

	fmt.Println()
}

// 10. Ejecutar comandos del sistema
func ejemploComandosSistema() {
	fmt.Println("10. EJECUTAR COMANDOS DEL SISTEMA:")

	// Comando simple
	cmd := exec.Command("echo", "Hola desde comando del sistema")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error ejecutando comando: %v\n", err)
	} else {
		fmt.Printf("Salida del comando: %s", string(output))
	}

	// Comando con múltiples argumentos
	var listCmd *exec.Cmd
	if strings.Contains(strings.ToLower(os.Getenv("OS")), "windows") {
		listCmd = exec.Command("cmd", "/c", "dir")
	} else {
		listCmd = exec.Command("ls", "-la")
	}

	listOutput, err := listCmd.Output()
	if err != nil {
		fmt.Printf("Error ejecutando comando de listado: %v\n", err)
	} else {
		fmt.Printf("Primeras líneas del listado:\n")
		lines := strings.Split(string(listOutput), "\n")
		for i, line := range lines {
			if i >= 5 { // Solo mostrar las primeras 5 líneas
				fmt.Println("  ...")
				break
			}
			fmt.Printf("  %s\n", line)
		}
	}

	// Comando con entrada estándar
	echoCmd := exec.Command("echo", "Texto procesado")
	echoOutput, err := echoCmd.Output()
	if err != nil {
		fmt.Printf("Error ejecutando echo: %v\n", err)
	} else {
		fmt.Printf("Echo output: %s", string(echoOutput))
	}

	fmt.Println()
}

// 11. Señales del sistema
func ejemploSeñales() {
	fmt.Println("11. SEÑALES DEL SISTEMA:")

	// Crear un proceso hijo para demostrar señales
	fmt.Println("Las señales se manejarían aquí normalmente con os/signal")
	fmt.Println("Ejemplo conceptual de manejo de señales:")
	fmt.Println("  - SIGINT: Interrupción (Ctrl+C)")
	fmt.Println("  - SIGTERM: Terminación")
	fmt.Println("  - SIGUSR1/SIGUSR2: Señales definidas por usuario")

	// En un programa real, usarías:
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// go func() {
	//     <-c
	//     fmt.Println("Señal recibida, limpiando...")
	//     os.Exit(0)
	// }()

	fmt.Println()
}

// 12. Stdin, Stdout, Stderr
func ejemploEntradaSalida() {
	fmt.Println("12. ENTRADA, SALIDA Y ERROR ESTÁNDAR:")

	// Escribir a stdout
	fmt.Fprintln(os.Stdout, "Esto se escribe a stdout")

	// Escribir a stderr
	fmt.Fprintln(os.Stderr, "Esto se escribe a stderr")

	// Verificar si stdout es un terminal
	if fileInfo, err := os.Stdout.Stat(); err == nil {
		if fileInfo.Mode()&os.ModeCharDevice != 0 {
			fmt.Println("stdout es un terminal")
		} else {
			fmt.Println("stdout está siendo redirigido")
		}
	}

	// Ejemplo de lectura de stdin (comentado para no bloquear la ejecución)
	// fmt.Print("Ingresa texto (presiona Enter): ")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	//     fmt.Printf("Error leyendo stdin: %v\n", err)
	// } else {
	//     fmt.Printf("Leído desde stdin: %s", input)
	// }

	fmt.Println("(Lectura de stdin comentada para no bloquear la demo)")

	fmt.Println()
}

// Función auxiliar para copiar archivos
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// Funciones simplificadas para mostrar vista previa de ejemplos avanzados
func ejemploVariablesEntornoSimplificado() {
	fmt.Println("Variables de entorno importantes:")
	importantes := []string{"PATH", "HOME", "USERPROFILE", "TEMP", "TMP", "USER", "USERNAME"}
	for _, env := range importantes {
		if value := os.Getenv(env); value != "" {
			if len(value) > 50 {
				fmt.Printf("  %s=%s...\n", env, value[:50])
			} else {
				fmt.Printf("  %s=%s\n", env, value)
			}
		}
	}
	fmt.Printf("Total de variables: %d\n", len(os.Environ()))
}

func ejemploFlagsSimplificado() {
	fmt.Println("Ejemplo de flags avanzados disponibles:")
	fmt.Println("  -name string    : Nombre del usuario")
	fmt.Println("  -age int       : Edad del usuario")
	fmt.Println("  -emails string : Lista de emails separados por comas")
	fmt.Println("  -help          : Mostrar ayuda")
	fmt.Println("Para ver el ejemplo completo, ejecute los comandos mostrados arriba.")
}

func ejemploArchivosSimplificado() {
	fmt.Println("Operaciones avanzadas de archivos disponibles:")
	fmt.Println("  ✓ Creación de estructuras de directorios complejas")
	fmt.Println("  ✓ Búsqueda de archivos por patrón")
	fmt.Println("  ✓ Recorrido recursivo de directorios")
	fmt.Println("  ✓ Monitoreo de cambios en archivos")
	fmt.Println("  ✓ Manejo eficiente de archivos grandes")
	fmt.Println("  ✓ Operaciones con metadatos y enlaces simbólicos")
	fmt.Println("Para ver el ejemplo completo, ejecute los comandos mostrados arriba.")
}
