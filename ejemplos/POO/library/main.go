package main

import (
	//"fmt"
	"library/animal"
	"library/book"
)

func main() {
	// Crear un objeto cuando no se tiene ninguna de sus propiedades privadas ni un metodo contructor
	// var mybook = book.Book {
	// 	Title: "Moby Dick",
	// 	Author: "Herman Melville",
	// 	Pages: 300,
	// }

	// Crear objeto con su "constructor" cuando sus atributos son privados
	mybook := book.NewBook("Moby Dick", "Herman Melville", 300)

	mybook.PrintInfo()

	mybook.SetTitle("Moby Dick [New Edition]")
	//fmt.Println(mybook.GetTitle())

	myTextBook := book.NewTextBook("Mil años de soledad", "Gabriel Garcia Marquez", 471, "booklet", "secundaria")
	//myTextBook.PrintInfo()

	book.Print(myTextBook)
	book.Print(myTextBook)

	/////////////////////

	miPerro := animal.Perro{Nombre: "Cody"}
	miGato := animal.Gato{Nombre: "Lio"}
	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

}
