package render

import (
	"bookings/internal/config"
	"bookings/internal/models"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
)

var app *config.AppConfig

// NewTemplates sets new config for the template package
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData{
	td.Flash = app.Session.PopString(r.Context(),"flash")
	td.Error = app.Session.PopString(r.Context(),"error")
	td.Warning = app.Session.PopString(r.Context(),"warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request,tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// Get the template cache from app config
		tc = app.TemplateCache
	}else {
		tc, _ = CreateTemplateCache()
	}


	// get requested template from cache 
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Coul not get template from template cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}
	// get all of the files named *.page.tmple from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending whit *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
