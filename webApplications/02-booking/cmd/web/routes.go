package main

import (
	"bookings/internal/config"
	"bookings/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// usando pat agregar al import "github.com/bmizerany/pat"
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// usando chi


	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)


	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)
	mux.Get("/generals-quarters",handlers.Repo.Generals)
	mux.Get("/majors-suite",handlers.Repo.Majors)
	
	mux.Get("/search-availability",handlers.Repo.Availability)
	mux.Post("/search-availability",handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json",handlers.Repo.AvailabilityJson)

	mux.Get("/contact",handlers.Repo.Contact)
	
	mux.Get("/make-reservation",handlers.Repo.Reservation)
	mux.Post("/make-reservation",handlers.Repo.PostReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}