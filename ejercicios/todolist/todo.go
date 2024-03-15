package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tarea struct {
	Id         int
	Nombre     string
	Detalle    string
	Finalizado bool `default:"false"`
}

var contadorID int = 1
var reader = bufio.NewReader(os.Stdin)
func leerEntrada() string {
    entrada, _ := reader.ReadString('\n')
	entrada = strings.TrimSpace(entrada)
    return entrada
}

func (t *Tarea) modificarTarea() {
	var detalleNuevo string
	fmt.Println("Ingresa el nuevo detalle de la tarea")
	detalleNuevo = leerEntrada()
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

func agregarTarea(tareas *[]Tarea){
	var nombre, detalle string
	fmt.Print("Nombre:")
	nombre = leerEntrada()
	fmt.Print("Detalle:")
	detalle = leerEntrada()
	nuevaTarea := Tarea {Id:contadorID ,Nombre: nombre,Detalle: detalle} 
	*tareas = append(*tareas, nuevaTarea)
	contadorID ++
}

func listarTareas(tareas *[]Tarea){
	for _, tarea := range *tareas{
		fmt.Printf("%v) %s: %s Finalizado: %t \n",tarea.Id, tarea.Nombre, tarea.Detalle, tarea.Finalizado)
	}
}
func remove(slice []Tarea, s int) []Tarea {
    return append(slice[:s], slice[s+1:]...)
}

func main() {
	misTareas := []Tarea{}
	continuar := true

	for continuar {
		fmt.Print(`
		Menu
		Por favor elije una opción
		1) Agregar Tarea
		2) Lista de Tareas
		3) Finalizar Tarea
		4) Eliminar Tarea
		`)
		var opcion int
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			agregarTarea(&misTareas)
		case 2:
			listarTareas(&misTareas)
		case 3:
			var idTarea int
			fmt.Print("Id:")
			fmt.Scanln(&idTarea)
			misTareas[idTarea].finalizar()
		case 4:
			var idTarea int
			fmt.Print("Id:")
			fmt.Scanln(&idTarea)
			misTareas = remove(misTareas, idTarea)
		}

	}


}