package book

import "fmt"

// Para el uso de polimorfismo
type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

// Con atributos publicos
// type Book struct {
// 	Title  string
// 	Author string
// 	Pages  int
// }
type Book struct {
	title  string
	author string
	pages  int
}

// Simular metodo constructor
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

// Ya que los campos son privados se crean getters y setters
func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

type TextBook struct {
	// La composición simula la herencia
	Book
	editorial string
	level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nNivel: %s\n",
		b.title, b.author, b.pages, b.editorial, b.level)
}
