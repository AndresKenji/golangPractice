package main

import "fmt"

type Tarea struct {
	Id         int
	Nombre     string
	Detalle    string
	Finalizado bool
}

func (t *Tarea) modificarTarea() {
	var detalleNuevo string
	fmt.Println("Ingresa el nuevo detalle de la tarea")
	fmt.Scanln(&detalleNuevo)
	t.Detalle = detalleNuevo
	fmt.Println("Detalle modificado")
}

func (t *Tarea) finalizar(){
	t.Finalizado = !t.Finalizado
	if t.Finalizado {
		fmt.Println("Tarea finalizada!")
	}else {
		fmt.Println("Tarea iniciada!")
	}
}

func main() {
	misTareas := []Tarea{}
	continuar := true

	for continuar {
		fmt.Print(`
		Menu
		1) Agregar Tarea
		2) Lista de Tareas
		3) Eliminar Tarea
		`)
	}


}