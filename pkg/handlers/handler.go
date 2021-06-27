package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/agniadvani/bookings/pkg/config"
	"github.com/agniadvani/bookings/pkg/models"
	render "github.com/agniadvani/bookings/pkg/renders"
)

//A variable used to reference outide the package scope values
var Repo *Repository

//Struct that holds the variable app, can be referenced outside the package scope using New Repo
type Repository struct {
	app *config.AppConfig
}

//Repository init function
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

//Handler init function
func NewHandler(r *Repository) {
	Repo = r
}

//Hnadler for "/home"
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r, "home.page.html", &models.TemplateData{})
}

//Handler for "/about"
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r, "about.page.html", &models.TemplateData{})
}

//Handler for "/contact"
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"contact.page.html",&models.TemplateData{})
}
//Handler for "/majors-suite"
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"majors.page.html",&models.TemplateData{})
}
//Handler for "/presidential-suite"
func (m *Repository) Presidential(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"presidential.page.html",&models.TemplateData{})
}
//Handler for "/search-availability"
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"search-availability.page.html",&models.TemplateData{})
}
//Handler for Post on "/search-availability"
func (m *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Post method..."))
}

type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message"`
}
//Handler for "/search-availability-json"
func (m *Repository) SearchAvailabilityJson(w http.ResponseWriter, r *http.Request){
	resp := jsonResponse{
		OK: true,
		Message: "Available!",
	}

	bs,err := json.MarshalIndent(resp,"","     ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bs))
	w.Header().Set("Content-Type","application/json")
	w.Write(bs)
}
//Handler for "/make-reservation"
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"make-reservation.page.html",&models.TemplateData{})
}

