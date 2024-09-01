package handlers

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"bookings/pkg/render"
	"encoding/json"
	"log"
	"net/http"
)


var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	
	stringMap["remote_ip"] = remoteIP

	// send data to template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and display form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "make-reservation.page.tmpl",&models.TemplateData{})
}

// Generals renders the  Generals room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "generals.page.tmpl",&models.TemplateData{})
}

// Majors renders the  Majors room page 
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "majors.page.tmpl",&models.TemplateData{})
}


// Availability renders the search availability page 
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "search-availability.page.tmpl",&models.TemplateData{})
}

// PostAvailability renders the search availability page 
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request){
	// start := r.Form.Get("start")
	// end := r.Form.Get("end")
	w.Write([]byte("Posted to search availability"))
}

type jsonResponse struct{
	OK bool `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJson handles request for availability and send JSOn response
func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request){
	resp := jsonResponse{
		OK: true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp,"","    ")
	if err != nil{
		log.Println(err)
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(out)
}

// Contact renders the search availability page 
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "contact.page.tmpl",&models.TemplateData{})
}