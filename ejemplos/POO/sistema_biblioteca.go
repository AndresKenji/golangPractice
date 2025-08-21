package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// ============= SISTEMA COMPLETO DE BIBLIOTECA =============

// Interfaces principales
type Prestable interface {
	ObtenerID() string
	ObtenerTitulo() string
	EstaDisponible() bool
	Prestar() error
	Devolver() error
	ObtenerTipo() string
}

type Busquedable interface {
	BuscarPorTitulo(titulo string) []Prestable
	BuscarPorAutor(autor string) []Prestable
	BuscarPorCategoria(categoria string) []Prestable
}

type Notificable interface {
	EnviarNotificacion(usuario *Usuario, mensaje string)
}

// Struct base para elementos de biblioteca
type ElementoBiblioteca struct {
	id            string
	titulo        string
	categoria     string
	disponible    bool
	fechaPrestamo *time.Time
	usuarioActual *Usuario
}

func (eb *ElementoBiblioteca) ObtenerID() string {
	return eb.id
}

func (eb *ElementoBiblioteca) ObtenerTitulo() string {
	return eb.titulo
}

func (eb *ElementoBiblioteca) EstaDisponible() bool {
	return eb.disponible
}

func (eb *ElementoBiblioteca) prestarElemento(usuario *Usuario) error {
	if !eb.disponible {
		return errors.New("elemento no disponible")
	}

	eb.disponible = false
	ahora := time.Now()
	eb.fechaPrestamo = &ahora
	eb.usuarioActual = usuario
	return nil
}

func (eb *ElementoBiblioteca) devolverElemento() error {
	if eb.disponible {
		return errors.New("elemento no está prestado")
	}

	eb.disponible = true
	eb.fechaPrestamo = nil
	eb.usuarioActual = nil
	return nil
}

// Libro - hereda de ElementoBiblioteca
type Libro struct {
	ElementoBiblioteca
	autor          string
	isbn           string
	paginas        int
	editorial      string
	añoPublicacion int
}

func NuevoLibro(id, titulo, autor, isbn, categoria string, paginas, año int, editorial string) *Libro {
	return &Libro{
		ElementoBiblioteca: ElementoBiblioteca{
			id:         id,
			titulo:     titulo,
			categoria:  categoria,
			disponible: true,
		},
		autor:          autor,
		isbn:           isbn,
		paginas:        paginas,
		editorial:      editorial,
		añoPublicacion: año,
	}
}

func (l *Libro) Prestar() error {
	return l.prestarElemento(l.usuarioActual)
}

func (l *Libro) Devolver() error {
	return l.devolverElemento()
}

func (l *Libro) ObtenerTipo() string {
	return "Libro"
}

func (l *Libro) ObtenerAutor() string {
	return l.autor
}

func (l *Libro) ObtenerISBN() string {
	return l.isbn
}

func (l *Libro) String() string {
	estado := "Disponible"
	if !l.disponible {
		estado = fmt.Sprintf("Prestado a %s", l.usuarioActual.Nombre())
	}
	return fmt.Sprintf("📚 %s por %s (%s) - %s", l.titulo, l.autor, l.categoria, estado)
}

// Revista - hereda de ElementoBiblioteca
type Revista struct {
	ElementoBiblioteca
	numero    int
	mes       string
	año       int
	editor    string
	articulos []string
}

func NuevaRevista(id, titulo, categoria string, numero int, mes string, año int, editor string) *Revista {
	return &Revista{
		ElementoBiblioteca: ElementoBiblioteca{
			id:         id,
			titulo:     titulo,
			categoria:  categoria,
			disponible: true,
		},
		numero:    numero,
		mes:       mes,
		año:       año,
		editor:    editor,
		articulos: make([]string, 0),
	}
}

func (r *Revista) Prestar() error {
	return r.prestarElemento(r.usuarioActual)
}

func (r *Revista) Devolver() error {
	return r.devolverElemento()
}

func (r *Revista) ObtenerTipo() string {
	return "Revista"
}

func (r *Revista) AgregarArticulo(articulo string) {
	r.articulos = append(r.articulos, articulo)
}

func (r *Revista) String() string {
	estado := "Disponible"
	if !r.disponible {
		estado = fmt.Sprintf("Prestado a %s", r.usuarioActual.Nombre())
	}
	return fmt.Sprintf("📰 %s #%d (%s %d) - %s", r.titulo, r.numero, r.mes, r.año, estado)
}

// DVD - hereda de ElementoBiblioteca
type DVD struct {
	ElementoBiblioteca
	director string
	duracion int // en minutos
	año      int
	genero   string
	actores  []string
}

func NuevoDVD(id, titulo, director, categoria, genero string, duracion, año int) *DVD {
	return &DVD{
		ElementoBiblioteca: ElementoBiblioteca{
			id:         id,
			titulo:     titulo,
			categoria:  categoria,
			disponible: true,
		},
		director: director,
		duracion: duracion,
		año:      año,
		genero:   genero,
		actores:  make([]string, 0),
	}
}

func (d *DVD) Prestar() error {
	return d.prestarElemento(d.usuarioActual)
}

func (d *DVD) Devolver() error {
	return d.devolverElemento()
}

func (d *DVD) ObtenerTipo() string {
	return "DVD"
}

func (d *DVD) AgregarActor(actor string) {
	d.actores = append(d.actores, actor)
}

func (d *DVD) String() string {
	estado := "Disponible"
	if !d.disponible {
		estado = fmt.Sprintf("Prestado a %s", d.usuarioActual.Nombre())
	}
	return fmt.Sprintf("🎬 %s dirigida por %s (%d min) - %s", d.titulo, d.director, d.duracion, estado)
}

// Usuario del sistema
type Usuario struct {
	id                 string
	nombre             string
	apellido           string
	email              string
	telefono           string
	fechaRegistro      time.Time
	activo             bool
	elementosPrestados []Prestable
	historialPrestamos []HistorialPrestamo
	limiteElementos    int
}

type HistorialPrestamo struct {
	Elemento        Prestable
	FechaPrestamo   time.Time
	FechaDevolucion *time.Time
	Devuelto        bool
}

func NuevoUsuario(id, nombre, apellido, email, telefono string) *Usuario {
	return &Usuario{
		id:                 id,
		nombre:             nombre,
		apellido:           apellido,
		email:              email,
		telefono:           telefono,
		fechaRegistro:      time.Now(),
		activo:             true,
		elementosPrestados: make([]Prestable, 0),
		historialPrestamos: make([]HistorialPrestamo, 0),
		limiteElementos:    3,
	}
}

func (u *Usuario) ID() string {
	return u.id
}

func (u *Usuario) Nombre() string {
	return u.nombre + " " + u.apellido
}

func (u *Usuario) Email() string {
	return u.email
}

func (u *Usuario) PuedePrestar() bool {
	return u.activo && len(u.elementosPrestados) < u.limiteElementos
}

func (u *Usuario) ElementosPrestados() []Prestable {
	return u.elementosPrestados
}

func (u *Usuario) NumeroElementosPrestados() int {
	return len(u.elementosPrestados)
}

func (u *Usuario) String() string {
	estado := "Activo"
	if !u.activo {
		estado = "Inactivo"
	}
	return fmt.Sprintf("👤 %s (%s) - %s - Prestados: %d/%d",
		u.Nombre(), u.email, estado, len(u.elementosPrestados), u.limiteElementos)
}

// Sistema de notificaciones
type SistemaNotificaciones struct {
	notificaciones []Notificacion
}

type Notificacion struct {
	Usuario *Usuario
	Mensaje string
	Fecha   time.Time
	Tipo    string
	Leida   bool
}

func NuevoSistemaNotificaciones() *SistemaNotificaciones {
	return &SistemaNotificaciones{
		notificaciones: make([]Notificacion, 0),
	}
}

func (sn *SistemaNotificaciones) EnviarNotificacion(usuario *Usuario, mensaje string) {
	notificacion := Notificacion{
		Usuario: usuario,
		Mensaje: mensaje,
		Fecha:   time.Now(),
		Tipo:    "General",
		Leida:   false,
	}

	sn.notificaciones = append(sn.notificaciones, notificacion)
	fmt.Printf("    📧 Notificación a %s: %s\n", usuario.Nombre(), mensaje)
}

// Biblioteca principal - integra todo el sistema
type Biblioteca struct {
	nombre         string
	elementos      []Prestable
	usuarios       map[string]*Usuario
	notificaciones *SistemaNotificaciones
	estadisticas   EstadisticasBiblioteca
}

type EstadisticasBiblioteca struct {
	TotalPrestamos      int
	ElementosActivos    int
	UsuariosRegistrados int
	ElementosPrestados  int
}

func NuevaBiblioteca(nombre string) *Biblioteca {
	return &Biblioteca{
		nombre:         nombre,
		elementos:      make([]Prestable, 0),
		usuarios:       make(map[string]*Usuario),
		notificaciones: NuevoSistemaNotificaciones(),
		estadisticas:   EstadisticasBiblioteca{},
	}
}

// Implementación de Busquedable
func (b *Biblioteca) BuscarPorTitulo(titulo string) []Prestable {
	var resultados []Prestable
	tituloLower := strings.ToLower(titulo)

	for _, elemento := range b.elementos {
		if strings.Contains(strings.ToLower(elemento.ObtenerTitulo()), tituloLower) {
			resultados = append(resultados, elemento)
		}
	}

	return resultados
}

func (b *Biblioteca) BuscarPorAutor(autor string) []Prestable {
	var resultados []Prestable
	autorLower := strings.ToLower(autor)

	for _, elemento := range b.elementos {
		if libro, ok := elemento.(*Libro); ok {
			if strings.Contains(strings.ToLower(libro.ObtenerAutor()), autorLower) {
				resultados = append(resultados, elemento)
			}
		}
	}

	return resultados
}

func (b *Biblioteca) BuscarPorCategoria(categoria string) []Prestable {
	var resultados []Prestable
	categoriaLower := strings.ToLower(categoria)

	for _, elemento := range b.elementos {
		if eb, ok := elemento.(*Libro); ok {
			if strings.Contains(strings.ToLower(eb.categoria), categoriaLower) {
				resultados = append(resultados, elemento)
			}
		} else if er, ok := elemento.(*Revista); ok {
			if strings.Contains(strings.ToLower(er.categoria), categoriaLower) {
				resultados = append(resultados, elemento)
			}
		} else if ed, ok := elemento.(*DVD); ok {
			if strings.Contains(strings.ToLower(ed.categoria), categoriaLower) {
				resultados = append(resultados, elemento)
			}
		}
	}

	return resultados
}

// Métodos de gestión
func (b *Biblioteca) AgregarElemento(elemento Prestable) {
	b.elementos = append(b.elementos, elemento)
	b.estadisticas.ElementosActivos++
	fmt.Printf("    ✅ Elemento agregado: %s\n", elemento.ObtenerTitulo())
}

func (b *Biblioteca) RegistrarUsuario(usuario *Usuario) {
	b.usuarios[usuario.ID()] = usuario
	b.estadisticas.UsuariosRegistrados++
	b.notificaciones.EnviarNotificacion(usuario,
		fmt.Sprintf("Bienvenido a %s. Tu cuenta ha sido activada.", b.nombre))
}

func (b *Biblioteca) PrestarElemento(idElemento, idUsuario string) error {
	// Buscar usuario
	usuario, existe := b.usuarios[idUsuario]
	if !existe {
		return errors.New("usuario no encontrado")
	}

	if !usuario.PuedePrestar() {
		return errors.New("usuario no puede prestar más elementos")
	}

	// Buscar elemento
	var elemento Prestable
	for _, e := range b.elementos {
		if e.ObtenerID() == idElemento {
			elemento = e
			break
		}
	}

	if elemento == nil {
		return errors.New("elemento no encontrado")
	}

	if !elemento.EstaDisponible() {
		return errors.New("elemento no disponible")
	}

	// Realizar préstamo
	err := elemento.Prestar()
	if err != nil {
		return err
	}

	// Actualizar usuario
	usuario.elementosPrestados = append(usuario.elementosPrestados, elemento)
	usuario.historialPrestamos = append(usuario.historialPrestamos, HistorialPrestamo{
		Elemento:      elemento,
		FechaPrestamo: time.Now(),
		Devuelto:      false,
	})

	// Actualizar estadísticas
	b.estadisticas.TotalPrestamos++
	b.estadisticas.ElementosPrestados++

	// Enviar notificación
	b.notificaciones.EnviarNotificacion(usuario,
		fmt.Sprintf("Has prestado: %s", elemento.ObtenerTitulo()))

	return nil
}

func (b *Biblioteca) DevolverElemento(idElemento, idUsuario string) error {
	// Buscar usuario
	usuario, existe := b.usuarios[idUsuario]
	if !existe {
		return errors.New("usuario no encontrado")
	}

	// Buscar elemento en préstamos del usuario
	var elemento Prestable
	var indice int = -1

	for i, e := range usuario.elementosPrestados {
		if e.ObtenerID() == idElemento {
			elemento = e
			indice = i
			break
		}
	}

	if elemento == nil {
		return errors.New("elemento no está prestado por este usuario")
	}

	// Realizar devolución
	err := elemento.Devolver()
	if err != nil {
		return err
	}

	// Actualizar usuario
	usuario.elementosPrestados = append(usuario.elementosPrestados[:indice],
		usuario.elementosPrestados[indice+1:]...)

	// Actualizar historial
	for i := range usuario.historialPrestamos {
		if usuario.historialPrestamos[i].Elemento.ObtenerID() == idElemento &&
			!usuario.historialPrestamos[i].Devuelto {
			ahora := time.Now()
			usuario.historialPrestamos[i].FechaDevolucion = &ahora
			usuario.historialPrestamos[i].Devuelto = true
			break
		}
	}

	// Actualizar estadísticas
	b.estadisticas.ElementosPrestados--

	// Enviar notificación
	b.notificaciones.EnviarNotificacion(usuario,
		fmt.Sprintf("Has devuelto: %s", elemento.ObtenerTitulo()))

	return nil
}

func (b *Biblioteca) MostrarCatalogo() {
	fmt.Printf("    📚 Catálogo de %s:\n", b.nombre)
	for i, elemento := range b.elementos {
		fmt.Printf("      [%d] %s\n", i+1, elemento)
	}
}

func (b *Biblioteca) MostrarUsuarios() {
	fmt.Printf("    👥 Usuarios registrados en %s:\n", b.nombre)
	for _, usuario := range b.usuarios {
		fmt.Printf("      %s\n", usuario)
	}
}

func (b *Biblioteca) MostrarEstadisticas() {
	fmt.Printf("    📊 Estadísticas de %s:\n", b.nombre)
	fmt.Printf("      Total de elementos: %d\n", b.estadisticas.ElementosActivos)
	fmt.Printf("      Usuarios registrados: %d\n", b.estadisticas.UsuariosRegistrados)
	fmt.Printf("      Elementos prestados: %d\n", b.estadisticas.ElementosPrestados)
	fmt.Printf("      Total préstamos realizados: %d\n", b.estadisticas.TotalPrestamos)
}

func ejemploSistemaBiblioteca() {
	fmt.Println("  === SISTEMA COMPLETO DE BIBLIOTECA ===")

	// Crear biblioteca
	biblioteca := NuevaBiblioteca("Biblioteca Central")

	// Crear elementos del catálogo
	libro1 := NuevoLibro("L001", "El Arte de la Programación", "Donald Knuth",
		"978-0201896831", "Informática", 650, 1997, "Addison-Wesley")
	libro2 := NuevoLibro("L002", "Clean Code", "Robert Martin",
		"978-0132350884", "Informática", 464, 2008, "Prentice Hall")
	libro3 := NuevoLibro("L003", "Cien años de soledad", "Gabriel García Márquez",
		"978-8437604947", "Literatura", 471, 1967, "Sudamericana")

	revista1 := NuevaRevista("R001", "National Geographic", "Ciencia",
		5, "Mayo", 2024, "National Geographic Society")
	revista1.AgregarArticulo("Cambio climático en la Antártida")
	revista1.AgregarArticulo("Nuevas especies marinas")

	dvd1 := NuevoDVD("D001", "Matrix", "Lana Wachowski", "Ciencia Ficción", "Acción", 136, 1999)
	dvd1.AgregarActor("Keanu Reeves")
	dvd1.AgregarActor("Laurence Fishburne")

	// Agregar elementos a la biblioteca
	fmt.Println("  Agregando elementos al catálogo:")
	biblioteca.AgregarElemento(libro1)
	biblioteca.AgregarElemento(libro2)
	biblioteca.AgregarElemento(libro3)
	biblioteca.AgregarElemento(revista1)
	biblioteca.AgregarElemento(dvd1)

	// Crear usuarios
	fmt.Println("\n  Registrando usuarios:")
	usuario1 := NuevoUsuario("U001", "Ana", "García", "ana.garcia@email.com", "+1234567890")
	usuario2 := NuevoUsuario("U002", "Carlos", "López", "carlos.lopez@email.com", "+0987654321")
	usuario3 := NuevoUsuario("U003", "María", "Rodríguez", "maria.rodriguez@email.com", "+5555555555")

	biblioteca.RegistrarUsuario(usuario1)
	biblioteca.RegistrarUsuario(usuario2)
	biblioteca.RegistrarUsuario(usuario3)

	// Mostrar catálogo inicial
	fmt.Println("\n  Catálogo inicial:")
	biblioteca.MostrarCatalogo()

	// Realizar préstamos
	fmt.Println("\n  Realizando préstamos:")

	err := biblioteca.PrestarElemento("L001", "U001")
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		fmt.Printf("    ✅ Préstamo exitoso\n")
	}

	err = biblioteca.PrestarElemento("R001", "U001")
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		fmt.Printf("    ✅ Préstamo exitoso\n")
	}

	err = biblioteca.PrestarElemento("L002", "U002")
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		fmt.Printf("    ✅ Préstamo exitoso\n")
	}

	// Intentar préstamo de elemento no disponible
	err = biblioteca.PrestarElemento("L001", "U003")
	if err != nil {
		fmt.Printf("    ✗ Error esperado: %v\n", err)
	}

	// Mostrar estado después de préstamos
	fmt.Println("\n  Estado después de préstamos:")
	biblioteca.MostrarCatalogo()
	fmt.Println()
	biblioteca.MostrarUsuarios()

	// Búsquedas
	fmt.Println("\n  Realizando búsquedas:")

	resultados := biblioteca.BuscarPorTitulo("Clean")
	fmt.Printf("    Búsqueda por título 'Clean': %d resultado(s)\n", len(resultados))
	for _, r := range resultados {
		fmt.Printf("      - %s\n", r.ObtenerTitulo())
	}

	resultados = biblioteca.BuscarPorAutor("Knuth")
	fmt.Printf("    Búsqueda por autor 'Knuth': %d resultado(s)\n", len(resultados))
	for _, r := range resultados {
		fmt.Printf("      - %s\n", r.ObtenerTitulo())
	}

	resultados = biblioteca.BuscarPorCategoria("Informática")
	fmt.Printf("    Búsqueda por categoría 'Informática': %d resultado(s)\n", len(resultados))
	for _, r := range resultados {
		fmt.Printf("      - %s\n", r.ObtenerTitulo())
	}

	// Devoluciones
	fmt.Println("\n  Realizando devoluciones:")

	err = biblioteca.DevolverElemento("L001", "U001")
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		fmt.Printf("    ✅ Devolución exitosa\n")
	}

	err = biblioteca.DevolverElemento("L002", "U002")
	if err != nil {
		fmt.Printf("    Error: %v\n", err)
	} else {
		fmt.Printf("    ✅ Devolución exitosa\n")
	}

	// Estado final
	fmt.Println("\n  Estado final:")
	biblioteca.MostrarEstadisticas()

	fmt.Println("\n  Elementos disponibles para préstamo:")
	for _, elemento := range biblioteca.elementos {
		if elemento.EstaDisponible() {
			fmt.Printf("    ✅ %s\n", elemento)
		}
	}

	fmt.Println("\n  Historial de préstamos de Ana García:")
	for i, historial := range usuario1.historialPrestamos {
		estado := "Prestado"
		if historial.Devuelto {
			estado = fmt.Sprintf("Devuelto el %s", historial.FechaDevolucion.Format("2006-01-02 15:04"))
		}
		fmt.Printf("    [%d] %s - Prestado el %s - %s\n",
			i+1, historial.Elemento.ObtenerTitulo(),
			historial.FechaPrestamo.Format("2006-01-02 15:04"), estado)
	}
}
