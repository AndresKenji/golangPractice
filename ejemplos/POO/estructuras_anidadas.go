package main

import (
	"fmt"
	"time"
)

// ============= ESTRUCTURAS ANIDADAS =============

// Coordenadas geográficas
type Coordenadas struct {
	Latitud  float64
	Longitud float64
}

func (c Coordenadas) String() string {
	return fmt.Sprintf("(%.6f, %.6f)", c.Latitud, c.Longitud)
}

// Información de contacto detallada
type InformacionContacto struct {
	Email    string
	Telefono string
	Movil    string
	Sitio    string
	Redes    map[string]string // red social -> usuario
}

func (ic InformacionContacto) String() string {
	redes := ""
	for red, usuario := range ic.Redes {
		redes += fmt.Sprintf("%s: @%s, ", red, usuario)
	}
	if len(redes) > 0 {
		redes = redes[:len(redes)-2] // Quitar última coma
	}
	return fmt.Sprintf("Email: %s, Tel: %s, Móvil: %s, Web: %s, Redes: {%s}",
		ic.Email, ic.Telefono, ic.Movil, ic.Sitio, redes)
}

// Dirección completa con coordenadas
type DireccionCompleta struct {
	Calle        string
	Numero       string
	Piso         string
	Departamento string
	CodigoPostal string
	Ciudad       string
	Provincia    string
	Pais         string
	Coordenadas  Coordenadas // Estructura anidada
}

func (dc DireccionCompleta) DireccionCorta() string {
	direccion := fmt.Sprintf("%s %s", dc.Calle, dc.Numero)
	if dc.Piso != "" {
		direccion += fmt.Sprintf(", Piso %s", dc.Piso)
	}
	if dc.Departamento != "" {
		direccion += fmt.Sprintf(", Dpto %s", dc.Departamento)
	}
	return direccion
}

func (dc DireccionCompleta) String() string {
	return fmt.Sprintf("%s, %s %s, %s, %s - %s",
		dc.DireccionCorta(), dc.Ciudad, dc.CodigoPostal, dc.Provincia, dc.Pais, dc.Coordenadas)
}

// Empresa con múltiples estructuras anidadas
type Empresa struct {
	Nombre      string
	RazonSocial string
	CUIT        string
	Sector      string
	Fundacion   time.Time
	Sede        DireccionCompleta   // Estructura anidada
	Contacto    InformacionContacto // Estructura anidada
	Sucursales  []DireccionCompleta // Slice de estructuras
	Empleados   []EmpleadoEmpresa   // Slice de estructuras
	Finanzas    FinanzasEmpresa     // Estructura anidada
}

// Empleado específico de empresa
type EmpleadoEmpresa struct {
	ID           int
	Nombre       string
	Apellido     string
	Cargo        string
	Departamento string
	FechaIngreso time.Time
	Salario      SalarioEmpleado     // Estructura anidada
	Direccion    DireccionCompleta   // Estructura anidada
	Contacto     InformacionContacto // Estructura anidada
	Supervisor   *EmpleadoEmpresa    // Puntero a otra estructura del mismo tipo
	Subordinados []*EmpleadoEmpresa  // Slice de punteros
}

// Información salarial detallada
type SalarioEmpleado struct {
	SalarioBase       float64
	Bonos             float64
	Comisiones        float64
	Deducciones       float64
	Moneda            string
	FechaUltimoCambio time.Time
}

func (se SalarioEmpleado) SalarioNeto() float64 {
	return se.SalarioBase + se.Bonos + se.Comisiones - se.Deducciones
}

func (se SalarioEmpleado) String() string {
	return fmt.Sprintf("Base: %.2f %s, Bonos: %.2f, Comisiones: %.2f, Deducciones: %.2f, Neto: %.2f %s",
		se.SalarioBase, se.Moneda, se.Bonos, se.Comisiones, se.Deducciones, se.SalarioNeto(), se.Moneda)
}

// Finanzas de la empresa
type FinanzasEmpresa struct {
	Capital     float64
	Ingresos    IngresosPorPeriodo
	Gastos      GastosPorPeriodo
	Inversiones []Inversion
	Moneda      string
}

type IngresosPorPeriodo struct {
	Mensual    float64
	Trimestral float64
	Anual      float64
}

type GastosPorPeriodo struct {
	Operativos float64
	Personal   float64
	Marketing  float64
	Tecnologia float64
	Otros      float64
}

func (gpp GastosPorPeriodo) Total() float64 {
	return gpp.Operativos + gpp.Personal + gpp.Marketing + gpp.Tecnologia + gpp.Otros
}

type Inversion struct {
	Tipo        string
	Monto       float64
	Fecha       time.Time
	Descripcion string
}

// Métodos para Empresa
func (e *Empresa) AgregarEmpleado(empleado EmpleadoEmpresa) {
	e.Empleados = append(e.Empleados, empleado)
	fmt.Printf("  Empleado %s %s agregado a %s\n", empleado.Nombre, empleado.Apellido, e.Nombre)
}

func (e *Empresa) AgregarSucursal(sucursal DireccionCompleta) {
	e.Sucursales = append(e.Sucursales, sucursal)
	fmt.Printf("  Nueva sucursal agregada en %s\n", sucursal.Ciudad)
}

func (e Empresa) NumeroEmpleados() int {
	return len(e.Empleados)
}

func (e Empresa) NumeroSucursales() int {
	return len(e.Sucursales)
}

func (e Empresa) MasasSalarialTotal() float64 {
	total := 0.0
	for _, emp := range e.Empleados {
		total += emp.Salario.SalarioNeto()
	}
	return total
}

// Métodos para EmpleadoEmpresa
func (ee *EmpleadoEmpresa) AsignarSupervisor(supervisor *EmpleadoEmpresa) {
	ee.Supervisor = supervisor
	supervisor.Subordinados = append(supervisor.Subordinados, ee)
	fmt.Printf("  %s %s ahora reporta a %s %s\n",
		ee.Nombre, ee.Apellido, supervisor.Nombre, supervisor.Apellido)
}

func (ee EmpleadoEmpresa) TieneSubordinados() bool {
	return len(ee.Subordinados) > 0
}

func (ee EmpleadoEmpresa) NombreCompleto() string {
	return fmt.Sprintf("%s %s", ee.Nombre, ee.Apellido)
}

func (ee EmpleadoEmpresa) String() string {
	supervisor := "Sin supervisor"
	if ee.Supervisor != nil {
		supervisor = ee.Supervisor.NombreCompleto()
	}
	return fmt.Sprintf("ID: %d, %s (%s) - %s, Supervisor: %s, Subordinados: %d, Salario: %s",
		ee.ID, ee.NombreCompleto(), ee.Cargo, ee.Departamento, supervisor, len(ee.Subordinados), ee.Salario)
}

func ejemploEstructurasAnidadas() {
	// Crear coordenadas para las direcciones
	coordenadasSede := Coordenadas{Latitud: -34.6037, Longitud: -58.3816}     // Buenos Aires
	coordenadasSucursal := Coordenadas{Latitud: -31.4201, Longitud: -64.1888} // Córdoba

	// Crear direcciones completas
	sede := DireccionCompleta{
		Calle:        "Av. Corrientes",
		Numero:       "1234",
		Piso:         "10",
		Departamento: "A",
		CodigoPostal: "C1043",
		Ciudad:       "Buenos Aires",
		Provincia:    "Ciudad Autónoma de Buenos Aires",
		Pais:         "Argentina",
		Coordenadas:  coordenadasSede,
	}

	sucursal := DireccionCompleta{
		Calle:        "San Martín",
		Numero:       "567",
		CodigoPostal: "X5000",
		Ciudad:       "Córdoba",
		Provincia:    "Córdoba",
		Pais:         "Argentina",
		Coordenadas:  coordenadasSucursal,
	}

	// Crear información de contacto de la empresa
	contactoEmpresa := InformacionContacto{
		Email:    "info@tecnoinnovadora.com",
		Telefono: "+54 11 4567-8900",
		Movil:    "+54 9 11 5555-1234",
		Sitio:    "www.tecnoinnovadora.com",
		Redes: map[string]string{
			"LinkedIn":  "tecnoinnovadora",
			"Twitter":   "tecnoinnov",
			"Instagram": "tecnoinnovadora_ar",
		},
	}

	// Crear finanzas
	finanzas := FinanzasEmpresa{
		Capital: 5000000,
		Ingresos: IngresosPorPeriodo{
			Mensual:    250000,
			Trimestral: 750000,
			Anual:      3000000,
		},
		Gastos: GastosPorPeriodo{
			Operativos: 50000,
			Personal:   120000,
			Marketing:  25000,
			Tecnologia: 30000,
			Otros:      15000,
		},
		Moneda: "USD",
	}

	finanzas.Inversiones = []Inversion{
		{
			Tipo:        "Tecnología",
			Monto:       100000,
			Fecha:       time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
			Descripcion: "Servidores y infraestructura cloud",
		},
		{
			Tipo:        "Marketing Digital",
			Monto:       25000,
			Fecha:       time.Date(2024, 3, 10, 0, 0, 0, 0, time.UTC),
			Descripcion: "Campaña publicitaria online",
		},
	}

	// Crear empresa con todas las estructuras anidadas
	empresa := &Empresa{
		Nombre:      "TecnoInnovadora S.A.",
		RazonSocial: "Tecnología e Innovación Sociedad Anónima",
		CUIT:        "30-12345678-9",
		Sector:      "Tecnología",
		Fundacion:   time.Date(2015, 6, 15, 0, 0, 0, 0, time.UTC),
		Sede:        sede,
		Contacto:    contactoEmpresa,
		Sucursales:  []DireccionCompleta{sucursal},
		Finanzas:    finanzas,
	}

	fmt.Println("  Empresa creada:")
	fmt.Printf("    %s (CUIT: %s)\n", empresa.Nombre, empresa.CUIT)
	fmt.Printf("    Sector: %s, Fundada: %s\n", empresa.Sector, empresa.Fundacion.Format("2006-01-02"))
	fmt.Printf("    Sede: %s\n", empresa.Sede)
	fmt.Printf("    Contacto: %s\n", empresa.Contacto)

	// Crear empleados con estructuras anidadas complejas
	directorGeneral := EmpleadoEmpresa{
		ID:           1001,
		Nombre:       "Roberto",
		Apellido:     "Fernández",
		Cargo:        "Director General",
		Departamento: "Dirección",
		FechaIngreso: time.Date(2015, 6, 15, 0, 0, 0, 0, time.UTC),
		Salario: SalarioEmpleado{
			SalarioBase:       15000,
			Bonos:             3000,
			Comisiones:        2000,
			Deducciones:       1500,
			Moneda:            "USD",
			FechaUltimoCambio: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Direccion: DireccionCompleta{
			Calle:        "Libertador",
			Numero:       "8901",
			CodigoPostal: "C1425",
			Ciudad:       "Buenos Aires",
			Provincia:    "Ciudad Autónoma de Buenos Aires",
			Pais:         "Argentina",
			Coordenadas:  Coordenadas{Latitud: -34.5755, Longitud: -58.4084},
		},
		Contacto: InformacionContacto{
			Email:    "roberto.fernandez@tecnoinnovadora.com",
			Telefono: "+54 11 4567-8901",
			Movil:    "+54 9 11 6666-1001",
			Redes: map[string]string{
				"LinkedIn": "roberto-fernandez-dg",
			},
		},
	}

	gerenteTI := EmpleadoEmpresa{
		ID:           1002,
		Nombre:       "Laura",
		Apellido:     "Martínez",
		Cargo:        "Gerente de TI",
		Departamento: "Tecnología",
		FechaIngreso: time.Date(2016, 3, 1, 0, 0, 0, 0, time.UTC),
		Salario: SalarioEmpleado{
			SalarioBase:       8000,
			Bonos:             1500,
			Comisiones:        500,
			Deducciones:       800,
			Moneda:            "USD",
			FechaUltimoCambio: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Direccion: DireccionCompleta{
			Calle:        "Santa Fe",
			Numero:       "2345",
			Piso:         "5",
			Departamento: "B",
			CodigoPostal: "C1123",
			Ciudad:       "Buenos Aires",
			Provincia:    "Ciudad Autónoma de Buenos Aires",
			Pais:         "Argentina",
			Coordenadas:  Coordenadas{Latitud: -34.5956, Longitud: -58.3772},
		},
		Contacto: InformacionContacto{
			Email: "laura.martinez@tecnoinnovadora.com",
			Movil: "+54 9 11 7777-1002",
			Redes: map[string]string{
				"LinkedIn": "laura-martinez-ti",
				"GitHub":   "lmartinez-dev",
			},
		},
	}

	desarrollador := EmpleadoEmpresa{
		ID:           1003,
		Nombre:       "Diego",
		Apellido:     "Rodríguez",
		Cargo:        "Desarrollador Senior",
		Departamento: "Tecnología",
		FechaIngreso: time.Date(2018, 9, 15, 0, 0, 0, 0, time.UTC),
		Salario: SalarioEmpleado{
			SalarioBase:       5000,
			Bonos:             800,
			Comisiones:        200,
			Deducciones:       400,
			Moneda:            "USD",
			FechaUltimoCambio: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Direccion: DireccionCompleta{
			Calle:        "Rivadavia",
			Numero:       "5678",
			CodigoPostal: "C1406",
			Ciudad:       "Buenos Aires",
			Provincia:    "Ciudad Autónoma de Buenos Aires",
			Pais:         "Argentina",
			Coordenadas:  Coordenadas{Latitud: -34.6398, Longitud: -58.4631},
		},
		Contacto: InformacionContacto{
			Email: "diego.rodriguez@tecnoinnovadora.com",
			Movil: "+54 9 11 8888-1003",
			Redes: map[string]string{
				"GitHub":  "drodriguez-dev",
				"Twitter": "diego_codes",
			},
		},
	}

	// Agregar empleados a la empresa
	empresa.AgregarEmpleado(directorGeneral)
	empresa.AgregarEmpleado(gerenteTI)
	empresa.AgregarEmpleado(desarrollador)

	// Establecer jerarquía
	empresa.Empleados[1].AsignarSupervisor(&empresa.Empleados[0]) // Gerente TI reporta a Director
	empresa.Empleados[2].AsignarSupervisor(&empresa.Empleados[1]) // Desarrollador reporta a Gerente TI

	fmt.Println("\n  Información de empleados:")
	for i, emp := range empresa.Empleados {
		fmt.Printf("    [%d] %s\n", i+1, emp)
	}

	fmt.Printf("\n  Estadísticas de la empresa:\n")
	fmt.Printf("    Empleados: %d\n", empresa.NumeroEmpleados())
	fmt.Printf("    Sucursales: %d\n", empresa.NumeroSucursales())
	fmt.Printf("    Masa salarial total: %.2f %s\n", empresa.MasasSalarialTotal(), "USD")
	fmt.Printf("    Gastos mensuales totales: %.2f %s\n", empresa.Finanzas.Gastos.Total(), empresa.Finanzas.Moneda)
	fmt.Printf("    Ingresos anuales: %.2f %s\n", empresa.Finanzas.Ingresos.Anual, empresa.Finanzas.Moneda)

	fmt.Println("\n  Inversiones realizadas:")
	for i, inv := range empresa.Finanzas.Inversiones {
		fmt.Printf("    [%d] %s: %.2f %s (%s) - %s\n",
			i+1, inv.Tipo, inv.Monto, empresa.Finanzas.Moneda, inv.Fecha.Format("2006-01-02"), inv.Descripcion)
	}
}
