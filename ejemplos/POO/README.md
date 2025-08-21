# Programación Orientada a Objetos en Go

Este proyecto demuestra los conceptos fundamentales de la Programación Orientada a Objetos (POO) implementados en Go, mostrando cómo el lenguaje maneja estos conceptos de manera única.

## Estructura del Proyecto

### Archivos Principales

- **`main.go`** - Punto de entrada que ejecuta todos los ejemplos
- **`interfaces.go`** - Interfaces básicas y polimorfismo
- **`composicion.go`** - Composición y embedding (herencia en Go)
- **`estructuras_anidadas.go`** - Estructuras complejas anidadas
- **`encapsulacion.go`** - Encapsulación y control de acceso
- **`metodos.go`** - Métodos con diferentes tipos de receivers
- **`interface_vacia.go`** - Interface{} y type assertions
- **`patron_strategy.go`** - Patrón de diseño Strategy
- **`patron_observer.go`** - Patrón de diseño Observer
- **`sistema_biblioteca.go`** - Sistema completo integrando todos los conceptos

## Conceptos Demostrados

### 1. Interfaces y Polimorfismo (`interfaces.go`)
- Definición de interfaces
- Implementación implícita de interfaces
- Polimorfismo con diferentes tipos
- Interfaces compuestas
- Funciones que trabajan con interfaces

**Ejemplo:**
```go
type Forma interface {
    Area() float64
    Perimetro() float64
    String() string
}

type Rectangulo struct {
    Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
    return r.Ancho * r.Alto
}
```

### 2. Composición y Embedding (`composicion.go`)
- Embedding como alternativa a la herencia
- Composición de estructuras
- Promoción de métodos
- Jerarquías complejas

**Ejemplo:**
```go
type Persona struct {
    Nombre, Apellido string
    Edad int
}

type Empleado struct {
    Persona    // Embedding
    ID int
    Salario float64
}
```

### 3. Estructuras Anidadas (`estructuras_anidadas.go`)
- Estructuras dentro de estructuras
- Composición compleja
- Modelado de entidades del mundo real
- Gestión de relaciones entre objetos

### 4. Encapsulación (`encapsulacion.go`)
- Campos privados (minúsculas)
- Métodos getter y setter
- Control de acceso
- Validación en métodos de negocio

**Ejemplo:**
```go
type CuentaBancaria struct {
    numero  string  // privado
    balance float64 // privado
}

func (cb *CuentaBancaria) Depositar(monto float64, pin string) error {
    if !cb.validarPin(pin) {
        return errors.New("PIN incorrecto")
    }
    // lógica de depósito
}
```

### 5. Métodos y Receivers (`metodos.go`)
- Value receivers vs Pointer receivers
- Cuándo usar cada tipo
- Métodos que modifican vs métodos que solo leen
- Estructuras de datos como Stack y Lista Enlazada

### 6. Interface Vacía y Type Assertions (`interface_vacia.go`)
- `interface{}` para tipos genéricos
- Type assertions seguras
- Type switches
- Reflection básica
- Procesamiento genérico de datos

### 7. Patrón Strategy (`patron_strategy.go`)
- Implementación del patrón Strategy
- Intercambio dinámico de algoritmos
- Ejemplos: descuentos, ordenamiento, validaciones
- Flexibilidad en tiempo de ejecución

### 8. Patrón Observer (`patron_observer.go`)
- Implementación del patrón Observer
- Sistema de notificaciones
- Desacoplamiento entre componentes
- Ejemplos: noticias, monitoreo de temperatura

### 9. Sistema Completo (`sistema_biblioteca.go`)
- Integración de todos los conceptos POO
- Sistema de gestión de biblioteca
- Múltiples tipos de elementos (libros, revistas, DVDs)
- Gestión de usuarios y préstamos
- Búsquedas y estadísticas

## Características Únicas de Go

### 1. No hay Herencia Clásica
Go no tiene herencia tradicional, pero usa **embedding** para lograr comportamiento similar:
```go
type Animal struct {
    nombre string
}

type Perro struct {
    Animal  // embedding, no herencia
    raza string
}
```

### 2. Interfaces Implícitas
No necesitas declarar explícitamente que implementas una interfaz:
```go
type Writer interface {
    Write([]byte) (int, error)
}

// Cualquier tipo con método Write implementa Writer automáticamente
```

### 3. Composition over Inheritance
Go favorece la composición sobre la herencia, lo que lleva a diseños más flexibles.

### 4. Value vs Pointer Receivers
```go
// Value receiver - no modifica el original
func (r Rectangulo) Area() float64 { ... }

// Pointer receiver - puede modificar el original
func (r *Rectangulo) Escalar(factor float64) { ... }
```

## Ejecución

Para ejecutar todos los ejemplos:
```bash
go run .
```

Para compilar:
```bash
go build .
```

Para ejecutar el binario compilado:
```bash
./poo        # Linux/Mac
./poo.exe    # Windows
```

## Ejemplos de Salida

El programa ejecuta automáticamente todos los ejemplos, mostrando:

1. **Interfaces**: Diferentes formas geométricas con polimorfismo
2. **Composición**: Empleados, clientes y contactos
3. **Estructuras anidadas**: Sistema empresarial complejo
4. **Encapsulación**: Cuentas bancarias y gestión de empleados
5. **Métodos**: Puntos, vectores, contadores y estructuras de datos
6. **Type assertions**: Procesamiento genérico de tipos
7. **Strategy**: Algoritmos intercambiables
8. **Observer**: Sistema de notificaciones
9. **Sistema biblioteca**: Ejemplo completo e integrado

## Patrones de Diseño Implementados

- **Strategy Pattern**: Para algoritmos intercambiables
- **Observer Pattern**: Para notificaciones y eventos
- **Factory Pattern**: Constructores de objetos
- **Repository Pattern**: En el sistema de biblioteca

## Mejores Prácticas Demostradas

1. **Naming conventions**: Público (mayúscula) vs privado (minúscula)
2. **Error handling**: Manejo explícito de errores
3. **Interface design**: Interfaces pequeñas y enfocadas
4. **Composition**: Prefer composition over inheritance
5. **Encapsulation**: Control de acceso a datos
6. **Testing-friendly**: Código diseñado para ser testeable

Este proyecto sirve como una guía completa para entender cómo aplicar conceptos de POO en Go, respetando las características únicas del lenguaje.
