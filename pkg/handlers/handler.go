package handler

import (
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
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

//Handler for "/about"
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}

//Handler for "/contact"
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"contact.page.html",&models.TemplateData{})
}
//Handler for "/majors-suite"
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"majors.page.html",&models.TemplateData{})
}
//Handler for "/presidential-suite"
func (m *Repository) Presidential(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"presidential.page.html",&models.TemplateData{})
}
//Handler for "/search-availability"
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"search-availability.page.html",&models.TemplateData{})
}
//Handler for "/make-reservation"
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,"make-reservation.page.html",&models.TemplateData{})
}