package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tarea struct {
	Nombre     string
	Detalle    string
	Finalizado bool `default:"false"`
}

type ListaTareas struct {
	tareas []Tarea
}

func (l *ListaTareas ) agregarTarea(t Tarea){
	l.tareas = append(l.tareas, t)
}
func (l *ListaTareas ) completarTarea(index int){
	l.tareas[index].Finalizado = true
}
func (l *ListaTareas ) editarTarea(index int, t Tarea){
	l.tareas[index] = t
}
func (l *ListaTareas ) eliminarTarea(index int){
	l.tareas = append(l.tareas[:index],l.tareas[index+1:]...)
}
func (t *ListaTareas ) listarTareas(){
	for i, tarea := range t.tareas{
		fmt.Printf("%v) %s: %s Finalizado: %t \n",i, tarea.Nombre, tarea.Detalle, tarea.Finalizado)
	}
}

var reader = bufio.NewReader(os.Stdin)
func leerEntrada() string {
    entrada, _ := reader.ReadString('\n')
	entrada = strings.TrimSpace(entrada)
    return entrada
}


func main() {
	misTareas := ListaTareas{}
	continuar := true

	for continuar {
		fmt.Print(`
********************************
*Menu                          *
*Por favor elije una opción    *
*1) Agregar Tarea              *
*2) Lista de Tareas            *
*3) Finalizar Tarea            *
*4) Eliminar Tarea             *
*5) Salir                      *
********************************
>`)
		var opcion int
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Agregar nueva tarea")
			fmt.Println("===============================================")
			fmt.Print("Nombre:")
			t.Nombre = leerEntrada()
			fmt.Print("Detalle:")
			t.Detalle = leerEntrada()
			fmt.Println("===============================================")
			misTareas.agregarTarea(t)
		case 2:
			misTareas.listarTareas()
		case 3:
			var indice int
			fmt.Print("Tarea:")
			fmt.Scan(&indice)
			misTareas.completarTarea(indice)			
		case 4:
			var indice int
			fmt.Print("Tarea:")
			fmt.Scan(&indice)
			misTareas.eliminarTarea(indice)					
		case 5:
			continuar = false

		}

	}


}