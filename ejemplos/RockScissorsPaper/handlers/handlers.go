package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	templateDir = "./templates"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.html","templates/base.html"))
	err := tpl.ExecuteTemplate(w,"base", nil)
	if err != nil {
		http.Error(w, "Error al renderizar el template:", http.StatusInternalServerError)
	}
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear nuevo Juego")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pagina de juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "jugar")
}


func RenderTemplate(w http.ResponseWriter, base,page string, data any){
	tpl := template.Must(template.ParseFiles(base,templateDir+page))
	err := tpl.ExecuteTemplate(w,"base", nil)
	if err != nil {
		http.Error(w, "Error al renderizar el template:", http.StatusInternalServerError)
	}
}