package main

import (
	"fmt"
	"math"
)

// ============= INTERFACES Y POLIMORFISMO =============

// Interfaz básica para figuras geométricas
type Forma interface {
	Area() float64
	Perimetro() float64
	String() string
}

// Interfaz para objetos que pueden moverse
type Movible interface {
	Mover(x, y float64)
	Posicion() (float64, float64)
}

// Interfaz compuesta
type FormaMovible interface {
	Forma
	Movible
}

// Rectángulo implementa Forma
type Rectangulo struct {
	Ancho, Alto float64
	x, y        float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

func (r Rectangulo) String() string {
	return fmt.Sprintf("Rectángulo(%.1fx%.1f) en (%.1f,%.1f)", r.Ancho, r.Alto, r.x, r.y)
}

func (r *Rectangulo) Mover(x, y float64) {
	r.x += x
	r.y += y
}

func (r Rectangulo) Posicion() (float64, float64) {
	return r.x, r.y
}

// Círculo implementa Forma
type Circulo struct {
	Radio float64
	x, y  float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

func (c Circulo) String() string {
	return fmt.Sprintf("Círculo(r=%.1f) en (%.1f,%.1f)", c.Radio, c.x, c.y)
}

func (c *Circulo) Mover(x, y float64) {
	c.x += x
	c.y += y
}

func (c Circulo) Posicion() (float64, float64) {
	return c.x, c.y
}

// Triángulo implementa Forma
type Triangulo struct {
	Base, Altura float64
	x, y         float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (t Triangulo) Perimetro() float64 {
	// Asumiendo triángulo isósceles para simplicidad
	lado := math.Sqrt((t.Base/2)*(t.Base/2) + t.Altura*t.Altura)
	return t.Base + 2*lado
}

func (t Triangulo) String() string {
	return fmt.Sprintf("Triángulo(base=%.1f, altura=%.1f) en (%.1f,%.1f)", t.Base, t.Altura, t.x, t.y)
}

func (t *Triangulo) Mover(x, y float64) {
	t.x += x
	t.y += y
}

func (t Triangulo) Posicion() (float64, float64) {
	return t.x, t.y
}

// Funciones que trabajan con interfaces
func imprimirInfoForma(f Forma) {
	fmt.Printf("  %s - Área: %.2f, Perímetro: %.2f\n", f.String(), f.Area(), f.Perimetro())
}

func moverForma(fm FormaMovible, dx, dy float64) {
	fmt.Printf("  Moviendo %s por (%.1f, %.1f)\n", fm.String(), dx, dy)
	fm.Mover(dx, dy)
	x, y := fm.Posicion()
	fmt.Printf("  Nueva posición: (%.1f, %.1f)\n", x, y)
}

func calcularAreaTotal(formas []Forma) float64 {
	total := 0.0
	for _, forma := range formas {
		total += forma.Area()
	}
	return total
}

func ejemploInterfaces() {
	// Crear diferentes formas
	rect := &Rectangulo{Ancho: 5, Alto: 3, x: 0, y: 0}
	circ := &Circulo{Radio: 2.5, x: 1, y: 1}
	tri := &Triangulo{Base: 4, Altura: 6, x: 2, y: 2}

	// Slice de interfaces
	formas := []Forma{rect, circ, tri}

	fmt.Println("  Formas creadas:")
	for _, forma := range formas {
		imprimirInfoForma(forma)
	}

	fmt.Printf("\n  Área total: %.2f\n", calcularAreaTotal(formas))

	// Demostrar FormaMovible
	fmt.Println("\n  Moviendo formas:")
	formasMovibles := []FormaMovible{rect, circ, tri}

	for i, forma := range formasMovibles {
		moverForma(forma, float64(i+1), float64(i+1))
	}

	fmt.Println("\n  Formas después del movimiento:")
	for _, forma := range formas {
		imprimirInfoForma(forma)
	}
}
