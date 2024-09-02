package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"bookings/internal/config"
	"bookings/internal/handlers"
	"bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8801"
var app config.AppConfig
var session *scs.SessionManager




func main() {

	// change this to true when in production
	app.InProduction = false

	// session para manejo de sesiones con el paquete github.com/alexedwards/scs/v2
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // para que la sesión persista incluso despues de cerrar el navegador
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)



	// http.HandleFunc("/", repo.Home)
	// http.HandleFunc("/about", repo.About)

	fmt.Printf("Starting application on port %s \n", portNumber)
	// _ = http.ListenAndServe(portNumber,nil)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
