package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// ============= ENCAPSULACIÓN =============

// CuentaBancaria - campos privados (minúsculas)
type CuentaBancaria struct {
	numero        string
	titular       string
	balance       float64
	fechaApertura time.Time
	activa        bool
	pin           string
	historial     []Transaccion
}

type Transaccion struct {
	id              int
	tipo            string
	monto           float64
	fecha           time.Time
	descripcion     string
	balanceAnterior float64
	balanceNuevo    float64
}

// Constructor público
func NuevaCuentaBancaria(numero, titular, pin string, balanceInicial float64) *CuentaBancaria {
	return &CuentaBancaria{
		numero:        numero,
		titular:       titular,
		balance:       balanceInicial,
		fechaApertura: time.Now(),
		activa:        true,
		pin:           pin,
		historial:     make([]Transaccion, 0),
	}
}

// Métodos getter públicos (acceso controlado a campos privados)
func (cb *CuentaBancaria) Numero() string {
	return cb.numero
}

func (cb *CuentaBancaria) Titular() string {
	return cb.titular
}

func (cb *CuentaBancaria) Balance() float64 {
	return cb.balance
}

func (cb *CuentaBancaria) FechaApertura() time.Time {
	return cb.fechaApertura
}

func (cb *CuentaBancaria) EstaActiva() bool {
	return cb.activa
}

// Métodos de negocio con validaciones
func (cb *CuentaBancaria) Depositar(monto float64, pin string) error {
	if !cb.validarPin(pin) {
		return errors.New("PIN incorrecto")
	}

	if !cb.activa {
		return errors.New("cuenta inactiva")
	}

	if monto <= 0 {
		return errors.New("el monto debe ser positivo")
	}

	balanceAnterior := cb.balance
	cb.balance += monto

	transaccion := Transaccion{
		id:              cb.generarIdTransaccion(),
		tipo:            "DEPÓSITO",
		monto:           monto,
		fecha:           time.Now(),
		descripcion:     fmt.Sprintf("Depósito de %.2f", monto),
		balanceAnterior: balanceAnterior,
		balanceNuevo:    cb.balance,
	}

	cb.historial = append(cb.historial, transaccion)
	return nil
}

func (cb *CuentaBancaria) Retirar(monto float64, pin string) error {
	if !cb.validarPin(pin) {
		return errors.New("PIN incorrecto")
	}

	if !cb.activa {
		return errors.New("cuenta inactiva")
	}

	if monto <= 0 {
		return errors.New("el monto debe ser positivo")
	}

	if monto > cb.balance {
		return errors.New("fondos insuficientes")
	}

	balanceAnterior := cb.balance
	cb.balance -= monto

	transaccion := Transaccion{
		id:              cb.generarIdTransaccion(),
		tipo:            "RETIRO",
		monto:           monto,
		fecha:           time.Now(),
		descripcion:     fmt.Sprintf("Retiro de %.2f", monto),
		balanceAnterior: balanceAnterior,
		balanceNuevo:    cb.balance,
	}

	cb.historial = append(cb.historial, transaccion)
	return nil
}

func (cb *CuentaBancaria) Transferir(monto float64, cuentaDestino *CuentaBancaria, pin string) error {
	if !cb.validarPin(pin) {
		return errors.New("PIN incorrecto")
	}

	if !cb.activa {
		return errors.New("cuenta origen inactiva")
	}

	if !cuentaDestino.activa {
		return errors.New("cuenta destino inactiva")
	}

	if monto <= 0 {
		return errors.New("el monto debe ser positivo")
	}

	if monto > cb.balance {
		return errors.New("fondos insuficientes")
	}

	// Retirar de cuenta origen
	balanceAnteriorOrigen := cb.balance
	cb.balance -= monto

	transaccionOrigen := Transaccion{
		id:              cb.generarIdTransaccion(),
		tipo:            "TRANSFERENCIA_SALIDA",
		monto:           monto,
		fecha:           time.Now(),
		descripcion:     fmt.Sprintf("Transferencia a %s (%.2f)", cuentaDestino.numero, monto),
		balanceAnterior: balanceAnteriorOrigen,
		balanceNuevo:    cb.balance,
	}

	// Depositar en cuenta destino
	balanceAnteriorDestino := cuentaDestino.balance
	cuentaDestino.balance += monto

	transaccionDestino := Transaccion{
		id:              cuentaDestino.generarIdTransaccion(),
		tipo:            "TRANSFERENCIA_ENTRADA",
		monto:           monto,
		fecha:           time.Now(),
		descripcion:     fmt.Sprintf("Transferencia desde %s (%.2f)", cb.numero, monto),
		balanceAnterior: balanceAnteriorDestino,
		balanceNuevo:    cuentaDestino.balance,
	}

	cb.historial = append(cb.historial, transaccionOrigen)
	cuentaDestino.historial = append(cuentaDestino.historial, transaccionDestino)

	return nil
}

// Método privado para validar PIN
func (cb *CuentaBancaria) validarPin(pin string) bool {
	return cb.pin == pin
}

// Método privado para generar ID de transacción
func (cb *CuentaBancaria) generarIdTransaccion() int {
	return rand.Intn(1000000) + 1
}

// Método para cambiar PIN (requiere PIN actual)
func (cb *CuentaBancaria) CambiarPin(pinActual, pinNuevo string) error {
	if !cb.validarPin(pinActual) {
		return errors.New("PIN actual incorrecto")
	}

	if len(pinNuevo) < 4 {
		return errors.New("el nuevo PIN debe tener al menos 4 dígitos")
	}

	cb.pin = pinNuevo
	return nil
}

// Método para obtener historial (con autenticación)
func (cb *CuentaBancaria) ObtenerHistorial(pin string) ([]Transaccion, error) {
	if !cb.validarPin(pin) {
		return nil, errors.New("PIN incorrecto")
	}

	// Retornar copia del historial para evitar modificaciones externas
	historialCopia := make([]Transaccion, len(cb.historial))
	copy(historialCopia, cb.historial)
	return historialCopia, nil
}

// Método para activar/desactivar cuenta (solo para administración)
func (cb *CuentaBancaria) CambiarEstado(activa bool, pinAdmin string) error {
	// En un sistema real, esto requeriría autenticación de administrador
	if pinAdmin != "ADMIN123" {
		return errors.New("acceso denegado - PIN de administrador incorrecto")
	}

	cb.activa = activa
	estado := "activada"
	if !activa {
		estado = "desactivada"
	}

	transaccion := Transaccion{
		id:              cb.generarIdTransaccion(),
		tipo:            "ADMIN",
		monto:           0,
		fecha:           time.Now(),
		descripcion:     fmt.Sprintf("Cuenta %s por administrador", estado),
		balanceAnterior: cb.balance,
		balanceNuevo:    cb.balance,
	}

	cb.historial = append(cb.historial, transaccion)
	return nil
}

func (cb *CuentaBancaria) String() string {
	estado := "ACTIVA"
	if !cb.activa {
		estado = "INACTIVA"
	}
	return fmt.Sprintf("Cuenta %s - %s (Estado: %s, Balance: $%.2f)",
		cb.numero, cb.titular, estado, cb.balance)
}

// Sistema de gestión de empleados con encapsulación
type SistemaEmpleados struct {
	empleados   map[int]*EmpleadoSistema
	siguienteID int
	adminPin    string
}

type EmpleadoSistema struct {
	id           int
	nombre       string
	apellido     string
	email        string
	departamento string
	salario      float64
	fechaIngreso time.Time
	activo       bool
	evaluaciones []Evaluacion
}

type Evaluacion struct {
	fecha       time.Time
	puntuacion  int // 1-10
	comentarios string
	evaluador   string
}

// Constructor del sistema
func NuevoSistemaEmpleados(adminPin string) *SistemaEmpleados {
	return &SistemaEmpleados{
		empleados:   make(map[int]*EmpleadoSistema),
		siguienteID: 1000,
		adminPin:    adminPin,
	}
}

// Métodos públicos con control de acceso
func (se *SistemaEmpleados) AgregarEmpleado(nombre, apellido, email, departamento string, salario float64, adminPin string) (int, error) {
	if !se.validarAdmin(adminPin) {
		return 0, errors.New("acceso denegado")
	}

	empleado := &EmpleadoSistema{
		id:           se.siguienteID,
		nombre:       nombre,
		apellido:     apellido,
		email:        email,
		departamento: departamento,
		salario:      salario,
		fechaIngreso: time.Now(),
		activo:       true,
		evaluaciones: make([]Evaluacion, 0),
	}

	se.empleados[se.siguienteID] = empleado
	id := se.siguienteID
	se.siguienteID++

	return id, nil
}

func (se *SistemaEmpleados) ObtenerEmpleado(id int, adminPin string) (*EmpleadoSistema, error) {
	if !se.validarAdmin(adminPin) {
		return nil, errors.New("acceso denegado")
	}

	empleado, existe := se.empleados[id]
	if !existe {
		return nil, errors.New("empleado no encontrado")
	}

	// Retornar copia para evitar modificaciones externas no controladas
	copia := *empleado
	return &copia, nil
}

func (se *SistemaEmpleados) ActualizarSalario(id int, nuevoSalario float64, adminPin string) error {
	if !se.validarAdmin(adminPin) {
		return errors.New("acceso denegado")
	}

	empleado, existe := se.empleados[id]
	if !existe {
		return errors.New("empleado no encontrado")
	}

	if nuevoSalario < 0 {
		return errors.New("el salario no puede ser negativo")
	}

	empleado.salario = nuevoSalario
	return nil
}

func (se *SistemaEmpleados) AgregarEvaluacion(id int, puntuacion int, comentarios, evaluador, adminPin string) error {
	if !se.validarAdmin(adminPin) {
		return errors.New("acceso denegado")
	}

	empleado, existe := se.empleados[id]
	if !existe {
		return errors.New("empleado no encontrado")
	}

	if puntuacion < 1 || puntuacion > 10 {
		return errors.New("la puntuación debe estar entre 1 y 10")
	}

	evaluacion := Evaluacion{
		fecha:       time.Now(),
		puntuacion:  puntuacion,
		comentarios: comentarios,
		evaluador:   evaluador,
	}

	empleado.evaluaciones = append(empleado.evaluaciones, evaluacion)
	return nil
}

func (se *SistemaEmpleados) ListarEmpleados(adminPin string) ([]EmpleadoSistema, error) {
	if !se.validarAdmin(adminPin) {
		return nil, errors.New("acceso denegado")
	}

	lista := make([]EmpleadoSistema, 0, len(se.empleados))
	for _, empleado := range se.empleados {
		lista = append(lista, *empleado)
	}

	return lista, nil
}

// Método privado para validar administrador
func (se *SistemaEmpleados) validarAdmin(pin string) bool {
	return se.adminPin == pin
}

func (es EmpleadoSistema) NombreCompleto() string {
	return fmt.Sprintf("%s %s", es.nombre, es.apellido)
}

func (es EmpleadoSistema) PromedioEvaluaciones() float64 {
	if len(es.evaluaciones) == 0 {
		return 0
	}

	suma := 0
	for _, eval := range es.evaluaciones {
		suma += eval.puntuacion
	}

	return float64(suma) / float64(len(es.evaluaciones))
}

func ejemploEncapsulacion() {
	fmt.Println("  === EJEMPLO DE CUENTAS BANCARIAS ===")

	// Crear cuentas bancarias
	cuenta1 := NuevaCuentaBancaria("001-12345", "Ana García", "1234", 1000.0)
	cuenta2 := NuevaCuentaBancaria("001-67890", "Carlos López", "5678", 500.0)

	fmt.Printf("    %s\n", cuenta1)
	fmt.Printf("    %s\n", cuenta2)

	// Operaciones válidas
	fmt.Println("\n  Operaciones bancarias:")

	err := cuenta1.Depositar(200, "1234")
	if err != nil {
		fmt.Printf("    Error en depósito: %v\n", err)
	} else {
		fmt.Printf("    Depósito exitoso: %s\n", cuenta1)
	}

	err = cuenta1.Retirar(150, "1234")
	if err != nil {
		fmt.Printf("    Error en retiro: %v\n", err)
	} else {
		fmt.Printf("    Retiro exitoso: %s\n", cuenta1)
	}

	err = cuenta1.Transferir(300, cuenta2, "1234")
	if err != nil {
		fmt.Printf("    Error en transferencia: %v\n", err)
	} else {
		fmt.Printf("    Transferencia exitosa:\n")
		fmt.Printf("      Origen: %s\n", cuenta1)
		fmt.Printf("      Destino: %s\n", cuenta2)
	}

	// Intentos de operaciones inválidas
	fmt.Println("\n  Intentos de operaciones inválidas:")

	err = cuenta1.Depositar(100, "0000") // PIN incorrecto
	if err != nil {
		fmt.Printf("    ✓ Depósito rechazado: %v\n", err)
	}

	err = cuenta2.Retirar(10000, "5678") // Fondos insuficientes
	if err != nil {
		fmt.Printf("    ✓ Retiro rechazado: %v\n", err)
	}

	// Mostrar historial
	fmt.Println("\n  Historial de transacciones:")
	historial, err := cuenta1.ObtenerHistorial("1234")
	if err != nil {
		fmt.Printf("    Error al obtener historial: %v\n", err)
	} else {
		for i, trans := range historial {
			fmt.Printf("    [%d] %s: %s - $%.2f (Balance: $%.2f -> $%.2f)\n",
				i+1, trans.fecha.Format("15:04:05"), trans.tipo, trans.monto,
				trans.balanceAnterior, trans.balanceNuevo)
		}
	}

	// Cambiar PIN
	err = cuenta1.CambiarPin("1234", "9876")
	if err != nil {
		fmt.Printf("    Error al cambiar PIN: %v\n", err)
	} else {
		fmt.Println("    PIN cambiado exitosamente")
	}

	// Administración de cuenta
	err = cuenta1.CambiarEstado(false, "ADMIN123") // Desactivar cuenta
	if err != nil {
		fmt.Printf("    Error al cambiar estado: %v\n", err)
	} else {
		fmt.Printf("    Cuenta desactivada: %s\n", cuenta1)
	}

	fmt.Println("\n  === EJEMPLO DE SISTEMA DE EMPLEADOS ===")

	// Crear sistema de empleados
	sistema := NuevoSistemaEmpleados("ADMIN2024")

	// Agregar empleados
	id1, err := sistema.AgregarEmpleado("María", "Rodríguez", "maria@empresa.com", "Desarrollo", 4500, "ADMIN2024")
	if err != nil {
		fmt.Printf("    Error al agregar empleado: %v\n", err)
	} else {
		fmt.Printf("    Empleado agregado con ID: %d\n", id1)
	}

	id2, err := sistema.AgregarEmpleado("Pedro", "Sánchez", "pedro@empresa.com", "Marketing", 3800, "ADMIN2024")
	if err != nil {
		fmt.Printf("    Error al agregar empleado: %v\n", err)
	} else {
		fmt.Printf("    Empleado agregado con ID: %d\n", id2)
	}

	// Intentar agregar empleado sin permisos
	_, err = sistema.AgregarEmpleado("Juan", "Pérez", "juan@empresa.com", "Ventas", 3500, "WRONG_PIN")
	if err != nil {
		fmt.Printf("    ✓ Acceso denegado (correcto): %v\n", err)
	}

	// Actualizar salario
	err = sistema.ActualizarSalario(id1, 5000, "ADMIN2024")
	if err != nil {
		fmt.Printf("    Error al actualizar salario: %v\n", err)
	} else {
		fmt.Printf("    Salario actualizado para empleado %d\n", id1)
	}

	// Agregar evaluaciones
	sistema.AgregarEvaluacion(id1, 9, "Excelente desempeño en el proyecto", "Gerente TI", "ADMIN2024")
	sistema.AgregarEvaluacion(id1, 8, "Muy buen trabajo en equipo", "Lead Developer", "ADMIN2024")
	sistema.AgregarEvaluacion(id2, 7, "Cumple con los objetivos", "Gerente Marketing", "ADMIN2024")

	// Listar empleados
	empleados, err := sistema.ListarEmpleados("ADMIN2024")
	if err != nil {
		fmt.Printf("    Error al listar empleados: %v\n", err)
	} else {
		fmt.Println("\n  Lista de empleados:")
		for _, emp := range empleados {
			fmt.Printf("    ID: %d - %s (%s) - $%.2f - Promedio evaluaciones: %.1f\n",
				emp.id, emp.NombreCompleto(), emp.departamento, emp.salario, emp.PromedioEvaluaciones())
		}
	}

	// Intentar acceder sin permisos
	_, err = sistema.ListarEmpleados("WRONG_PIN")
	if err != nil {
		fmt.Printf("    ✓ Acceso denegado a lista de empleados: %v\n", err)
	}
}
