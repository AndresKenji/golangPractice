package main

import (
	"fmt"
	"time"
)

// ============= COMPOSICIÓN Y EMBEDDING =============

// Struct base - Persona
type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
	Email    string
}

func (p Persona) NombreCompleto() string {
	return p.Nombre + " " + p.Apellido
}

func (p Persona) String() string {
	return fmt.Sprintf("%s (%d años) - %s", p.NombreCompleto(), p.Edad, p.Email)
}

func (p *Persona) CumplirAños() {
	p.Edad++
	fmt.Printf("  ¡Feliz cumpleaños %s! Ahora tiene %d años\n", p.NombreCompleto(), p.Edad)
}

// Dirección como componente
type Direccion struct {
	Calle     string
	Numero    int
	Ciudad    string
	CodigoZip string
	Pais      string
}

func (d Direccion) String() string {
	return fmt.Sprintf("%s %d, %s %s, %s", d.Calle, d.Numero, d.Ciudad, d.CodigoZip, d.Pais)
}

// Contacto combinando Persona y Dirección
type Contacto struct {
	Persona   // Embedding - hereda todos los campos y métodos
	Direccion // Composición
	Telefono  string
	FechaAlta time.Time
}

func (c Contacto) String() string {
	return fmt.Sprintf("Contacto: %s\n  Tel: %s\n  Dir: %s\n  Alta: %s",
		c.Persona.String(), c.Telefono, c.Direccion.String(), c.FechaAlta.Format("2006-01-02"))
}

// Empleado extiende Persona con información laboral
type Empleado struct {
	Persona      // Embedding
	ID           int
	Departamento string
	Salario      float64
	FechaIngreso time.Time
	Jefe         *Empleado // Referencia a otro empleado
}

func (e Empleado) String() string {
	jefeInfo := "Sin jefe asignado"
	if e.Jefe != nil {
		jefeInfo = e.Jefe.NombreCompleto()
	}
	return fmt.Sprintf("Empleado #%d: %s\n  Depto: %s, Salario: $%.2f\n  Jefe: %s\n  Ingreso: %s",
		e.ID, e.Persona.String(), e.Departamento, e.Salario, jefeInfo, e.FechaIngreso.Format("2006-01-02"))
}

func (e *Empleado) AsignarJefe(jefe *Empleado) {
	e.Jefe = jefe
	fmt.Printf("  %s ahora reporta a %s\n", e.NombreCompleto(), jefe.NombreCompleto())
}

func (e *Empleado) AumentarSalario(porcentaje float64) {
	salarioAnterior := e.Salario
	e.Salario += e.Salario * (porcentaje / 100)
	fmt.Printf("  Aumento salarial para %s: $%.2f -> $%.2f (+%.1f%%)\n",
		e.NombreCompleto(), salarioAnterior, e.Salario, porcentaje)
}

// Cliente extiende Contacto con información comercial
type Cliente struct {
	Contacto        // Embedding de Contacto (que ya incluye Persona)
	ID              int
	TipoCliente     string
	DescuentoActivo bool
	Compras         []Compra
}

type Compra struct {
	ID        int
	Fecha     time.Time
	Total     float64
	Productos []string
}

func (c *Cliente) AgregarCompra(compra Compra) {
	c.Compras = append(c.Compras, compra)
	fmt.Printf("  Nueva compra registrada para %s: $%.2f\n", c.NombreCompleto(), compra.Total)
}

func (c Cliente) TotalCompras() float64 {
	total := 0.0
	for _, compra := range c.Compras {
		total += compra.Total
	}
	return total
}

func (c Cliente) String() string {
	return fmt.Sprintf("Cliente #%d (%s): %s\n  Total compras: $%.2f (%d compras)\n  Descuento activo: %t",
		c.ID, c.TipoCliente, c.Contacto.String(), c.TotalCompras(), len(c.Compras), c.DescuentoActivo)
}

// Demostrar polimorfismo con interfaces
type Identificable interface {
	ObtenerID() int
	ObtenerTipo() string
}

func (e Empleado) ObtenerID() int      { return e.ID }
func (e Empleado) ObtenerTipo() string { return "Empleado" }

func (c Cliente) ObtenerID() int      { return c.ID }
func (c Cliente) ObtenerTipo() string { return "Cliente" }

func procesarIdentificable(i Identificable) {
	fmt.Printf("  Procesando %s con ID: %d\n", i.ObtenerTipo(), i.ObtenerID())
}

func ejemploComposicion() {
	// Crear personas base
	persona1 := Persona{
		Nombre:   "Ana",
		Apellido: "García",
		Edad:     28,
		Email:    "ana.garcia@email.com",
	}

	persona2 := Persona{
		Nombre:   "Carlos",
		Apellido: "López",
		Edad:     35,
		Email:    "carlos.lopez@email.com",
	}

	fmt.Println("  Personas creadas:")
	fmt.Printf("    %s\n", persona1)
	fmt.Printf("    %s\n", persona2)

	// Crear empleados usando embedding
	jefe := &Empleado{
		Persona:      persona2,
		ID:           1001,
		Departamento: "Tecnología",
		Salario:      75000,
		FechaIngreso: time.Date(2020, 3, 15, 0, 0, 0, 0, time.UTC),
	}

	empleado := &Empleado{
		Persona:      persona1,
		ID:           1002,
		Departamento: "Tecnología",
		Salario:      55000,
		FechaIngreso: time.Date(2022, 8, 10, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println("\n  Empleados:")
	fmt.Println(jefe)
	fmt.Println()
	fmt.Println(empleado)

	// Operaciones con empleados
	fmt.Println("\n  Operaciones:")
	empleado.AsignarJefe(jefe)
	empleado.AumentarSalario(10)
	empleado.CumplirAños() // Método heredado de Persona

	// Crear cliente con composición compleja
	direccion := Direccion{
		Calle:     "Av. Principal",
		Numero:    123,
		Ciudad:    "Madrid",
		CodigoZip: "28001",
		Pais:      "España",
	}

	contacto := Contacto{
		Persona: Persona{
			Nombre:   "María",
			Apellido: "Rodríguez",
			Edad:     42,
			Email:    "maria.rodriguez@email.com",
		},
		Direccion: direccion,
		Telefono:  "+34 600 123 456",
		FechaAlta: time.Now(),
	}

	cliente := &Cliente{
		Contacto:        contacto,
		ID:              2001,
		TipoCliente:     "Premium",
		DescuentoActivo: true,
	}

	// Agregar compras
	compra1 := Compra{
		ID:        3001,
		Fecha:     time.Now().AddDate(0, -1, 0),
		Total:     299.99,
		Productos: []string{"Laptop", "Mouse"},
	}

	compra2 := Compra{
		ID:        3002,
		Fecha:     time.Now().AddDate(0, 0, -15),
		Total:     89.50,
		Productos: []string{"Teclado", "Cable USB"},
	}

	cliente.AgregarCompra(compra1)
	cliente.AgregarCompra(compra2)

	fmt.Println("\n  Cliente completo:")
	fmt.Println(cliente)

	// Demostrar polimorfismo
	fmt.Println("\n  Polimorfismo con interfaces:")
	identificables := []Identificable{jefe, empleado, cliente}
	for _, item := range identificables {
		procesarIdentificable(item)
	}
}
