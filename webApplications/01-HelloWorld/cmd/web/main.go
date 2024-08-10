package main

import (
	"fmt"
	"log"
	"net/http"

	"helloworld/pkg/config"
	"helloworld/pkg/handlers"
	"helloworld/pkg/render"
)

const portNumber = ":8801"




func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber,nil)
}
