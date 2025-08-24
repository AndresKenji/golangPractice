# Ejemplos del Módulo OS en Go

Este directorio contiene ejemplos completos del uso del módulo `os` en Go, incluyendo manejo de archivos, flags, variables de entorno, y operaciones del sistema.

## Archivos Incluidos

### 📋 `main.go` - Programa Principal Modular
**Archivo principal con sistema de flags para ejecutar ejemplos específicos:**

- **Sistema de flags avanzado** para ejecutar ejemplos selectivos
- **12 secciones de ejemplos básicos** disponibles
- **Integración con ejemplos avanzados** en archivos separados
- **Ayuda detallada** con ejemplos de uso

### 🚩 `flags_example.go`
**Ejemplo avanzado de manejo de flags**

### 📁 `file_operations.go`
**Operaciones avanzadas con archivos**

### 🌍 `env_operations.go`
**Manejo avanzado de variables de entorno**

### 📖 `README.md`
**Documentación completa**

### ⚙️ `go.mod`
**Módulo Go**

## Cómo Ejecutar con Flags

### 🎯 **Ejemplos Básicos Selectivos:**

```bash
# Ver todas las opciones disponibles
go run main.go -help

# Ejecutar todos los ejemplos básicos (comportamiento por defecto)
go run main.go
go run main.go -all

# Solo información del sistema
go run main.go -system

# Solo variables de entorno básicas
go run main.go -env

# Solo ejemplos de archivos
go run main.go -files

# Solo directorios
go run main.go -dirs

# Solo rutas
go run main.go -paths

# Solo comandos del sistema
go run main.go -commands

# Solo flags básicos
go run main.go -flags
```

### 🌟 **Ejemplos Avanzados:**

```bash
# Variables de entorno avanzadas (con instrucciones)
go run main.go -env-advanced

# Archivos avanzados (con instrucciones)
go run main.go -files-advanced

# Flags avanzados (con instrucciones)
go run main.go -flags-advanced
```

### ⚡ **Opciones de Configuración:**

```bash
# Activar modo verbose
go run main.go -system -verbose

# Especificar archivo y configuración
go run main.go -files -file="mi_archivo.txt" -output="resultados" -iterations=3

# Combinaciones múltiples
go run main.go -env -dirs -verbose
```

## Flags Disponibles

### 📋 **Flags de Configuración:**
| Flag | Tipo | Default | Descripción |
|------|------|---------|-------------|
| `-verbose` | bool | false | Activar modo verbose |
| `-file` | string | "test.txt" | Nombre del archivo a procesar |
| `-output` | string | "output" | Directorio de salida |
| `-iterations` | int | 1 | Número de iteraciones |
| `-help` | bool | false | Mostrar ayuda detallada |

### 🎯 **Flags de Ejemplos Específicos:**
| Flag | Descripción |
|------|-------------|
| `-all` | Ejecutar todos los ejemplos básicos (default) |
| `-system` | Solo información del sistema |
| `-env` | Solo variables de entorno básicas |
| `-env-advanced` | Variables de entorno avanzadas |
| `-files` | Solo archivos básicos |
| `-files-advanced` | Archivos avanzados |
| `-dirs` | Solo directorios |
| `-paths` | Solo rutas |
| `-commands` | Solo comandos del sistema |
| `-flags` | Solo ejemplos de flags básicos |
| `-flags-advanced` | Ejemplos de flags avanzados |

## Ejemplos de Uso Prácticos

### 🔧 **Desarrollo y Debugging:**
```bash
# Ver información del sistema para debugging
go run main.go -system -verbose

# Verificar configuración de archivos
go run main.go -files -paths -verbose

# Revisar variables de entorno
go run main.go -env -verbose
```

### 📚 **Aprendizaje Incremental:**
```bash
# Paso 1: Conceptos básicos del sistema
go run main.go -system

# Paso 2: Variables de entorno
go run main.go -env

# Paso 3: Archivos básicos
go run main.go -files

# Paso 4: Directorios y rutas
go run main.go -dirs -paths

# Paso 5: Comandos del sistema
go run main.go -commands
```

### 🚀 **Ejemplos Avanzados:**
```bash
# Para ejecutar ejemplos avanzados de variables de entorno:
go run env_operations.go main.go

# Para ejecutar ejemplos avanzados de archivos:
go run file_operations.go main.go

# Para ejecutar ejemplos avanzados de flags:
go run flags_example.go main.go
```

## Contenido de los Ejemplos

### ✅ **Ejemplos Básicos (main.go):**

1. **Información del Sistema y Proceso**
   - PID del proceso actual y padre
   - UID/GID del usuario
   - Directorio de trabajo, home y temporal
   - Hostname del sistema

2. **Variables de Entorno**
   - Obtener, establecer y eliminar variables
   - Verificar existencia con `LookupEnv`
   - Listar todas las variables

3. **Argumentos y Flags de Línea de Comandos**
   - Parseo de flags con el paquete `flag`
   - Manejo de argumentos del programa
   - Flags personalizados (verbose, file, output, iterations)

4. **Trabajar con Archivos Básicos**
   - Crear, renombrar, copiar y eliminar archivos
   - Verificar existencia de archivos
   - Lectura y escritura básica

5. **Directorios y Navegación**
   - Crear y eliminar directorios
   - Navegación entre directorios
   - Listar contenido de directorios

6. **Información de Archivos y Directorios**
   - Obtener metadatos con `os.Stat`
   - Verificar permisos y propiedades
   - Información de tamaño y timestamps

7. **Permisos de Archivos**
   - Cambiar permisos con `os.Chmod`
   - Verificar permisos específicos

8. **Lectura y Escritura de Archivos**
   - Lectura/escritura completa de archivos
   - Lectura línea por línea con `bufio.Scanner`
   - Append a archivos existentes

9. **Trabajar con Rutas**
   - Construcción de rutas con `filepath.Join`
   - Separación de componentes de ruta
   - Rutas absolutas y relativas
   - Búsqueda con patrones glob

10. **Ejecutar Comandos del Sistema**
    - Ejecutar comandos con `os/exec`
    - Capturar salida de comandos
    - Comandos multiplataforma

11. **Señales del Sistema**
    - Conceptos de manejo de señales
    - Ejemplos de SIGINT, SIGTERM

12. **Entrada, Salida y Error Estándar**
    - Escribir a stdout/stderr
    - Verificar si la salida es un terminal
    - Conceptos de lectura de stdin

### 🚩 `flags_example.go`
**Ejemplo avanzado de manejo de flags:**
- Flags personalizados con validación
- Mensaje de ayuda personalizado
- Procesamiento de listas (emails)
- Flags booleanos y numéricos

### 📁 `file_operations.go`
**Operaciones avanzadas con archivos:**
- Creación de estructuras de directorios complejas
- Búsqueda de archivos por patrón
- Recorrido recursivo de directorios
- Monitoreo de cambios en archivos
- Manejo eficiente de archivos grandes
- Operaciones con metadatos y enlaces simbólicos

### 🌍 `env_operations.go`
**Manejo avanzado de variables de entorno:**
- Listado y agrupación de variables
- Análisis de la variable PATH
- Configuración desde variables de entorno
- Manejo de variables temporales

## Cómo Ejecutar

### Ejecutar el programa principal:
```bash
go run main.go
```

### Ejecutar con flags personalizados:
```bash
go run main.go -verbose -file="mi_archivo.txt" -output="resultados" -iterations=5
```

### Ver ayuda de flags:
```bash
go run main.go -help
```

### Ejecutar con argumentos adicionales:
```bash
go run main.go -verbose arg1 arg2 arg3
```

## Ejemplos de Flags Disponibles

| Flag | Tipo | Valor por defecto | Descripción |
|------|------|-------------------|-------------|
| `-verbose` | bool | false | Activar modo verbose |
| `-file` | string | "test.txt" | Nombre del archivo a procesar |
| `-output` | string | "output" | Directorio de salida |
| `-iterations` | int | 1 | Número de iteraciones |
| `-help` | bool | false | Mostrar ayuda |

## Funcionalidades Demonstradas

### ✅ **Manejo de Archivos:**
- Crear, leer, escribir, eliminar archivos
- Copiar archivos (implementación manual)
- Renombrar archivos
- Verificar existencia

### ✅ **Manejo de Directorios:**
- Crear directorios simples y anidados
- Listar contenido de directorios
- Navegación entre directorios
- Eliminación recursiva

### ✅ **Variables de Entorno:**
- Leer variables existentes
- Establecer nuevas variables
- Eliminar variables
- Verificar existencia

### ✅ **Información del Sistema:**
- PID del proceso
- Información del usuario
- Directorios del sistema
- Hostname

### ✅ **Argumentos de Línea de Comandos:**
- Parseo de flags con tipo
- Validación de argumentos
- Ayuda automática
- Argumentos posicionales

### ✅ **Operaciones Avanzadas:**
- Ejecución de comandos del sistema
- Manejo de permisos de archivos
- Trabajo con rutas multiplataforma
- Búsqueda de archivos con patrones

## Notas Importantes

1. **Multiplataforma**: Los ejemplos están diseñados para funcionar en Windows y Unix/Linux
2. **Manejo de Errores**: Todos los ejemplos incluyen manejo apropiado de errores
3. **Limpieza**: Los archivos temporales se eliminan automáticamente
4. **Seguridad**: Se usan permisos apropiados para archivos y directorios

## Estructura de Archivos Creados Durante la Ejecución

Durante la ejecución, el programa crea temporalmente:
```
ejemplos/os/
├── ejemplo.txt (temporal)
├── ejemplo_directorio/ (temporal)
├── lectura_escritura.txt (temporal)
├── permisos_ejemplo.txt (temporal)
└── archivos temporales del sistema
```

Todos los archivos temporales se limpian automáticamente al finalizar la ejecución.

## 🚀 Características Avanzadas del Sistema de Flags

### **Ejecución Modular:**
- Solo ejecuta los ejemplos que necesitas
- Reduce el tiempo de ejecución para pruebas específicas
- Permite aprendizaje incremental

### **Modo Verbose:**
- Información adicional en cualquier ejemplo
- Detalles de configuración
- Útil para debugging

### **Instrucciones Contextuales:**
- Los ejemplos avanzados muestran exactamente cómo ejecutarlos
- Múltiples opciones de ejecución
- Vistas previas de funcionalidad

### **Ayuda Integrada:**
```bash
go run main.go -help  # Muestra todas las opciones disponibles
```

## 🎯 Casos de Uso Recomendados

### **Para Aprender Go:**
```bash
# Día 1: Conceptos básicos
go run main.go -system -env

# Día 2: Archivos y directorios
go run main.go -files -dirs

# Día 3: Rutas y comandos
go run main.go -paths -commands

# Día 4: Ejemplos avanzados
go run main.go -env-advanced
go run main.go -files-advanced
```

### **Para Desarrollo:**
```bash
# Verificar configuración del sistema
go run main.go -system -env -verbose

# Probar operaciones con archivos
go run main.go -files -dirs -verbose

# Debuggear problemas de rutas
go run main.go -paths -verbose
```

### **Para Referencia Rápida:**
```bash
# Ver solo lo que necesitas
go run main.go -[tipo_ejemplo] -verbose
```

---

**💡 Tip Avanzado**: Usa el sistema de flags para crear tus propios scripts de aprendizaje. Combina diferentes flags según tus necesidades específicas de desarrollo.
