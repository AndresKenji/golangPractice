package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
// 	m := validPath.FindStringSubmatch(r.URL.Path)
// 	if m == nil {
// 		http.NotFound(w, r)
// 		return "", errors.New("Titulo invalido")
// 	}

// 	return m[2], nil
// }

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[2])

	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title := r.URL.Path[len("/view/"):]

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusTemporaryRedirect)
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpName string, data *Page) {
	err := templates.ExecuteTemplate(w, tmpName+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	fmt.Println("Iniciando server en puerto 8801")
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola mundo")
	})
	log.Fatal(http.ListenAndServe(":8801", nil))

}
