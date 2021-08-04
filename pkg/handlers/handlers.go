package handlers

import (
	"github.com/Fietzorama/Bookings/pkg/config"
	"github.com/Fietzorama/Bookings/pkg/models"
	"github.com/Fietzorama/Bookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository Type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home respond the request to the main page (main page handler)
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	//Grap the remote IP address of visitor at the FIRST visit
	remoteIP := r.RemoteAddr
	//Storing the IP Address as a string in the session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{}) //&TemplateData{} empty!
}

// About respond the request to the about page (about page handler)
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["Test"] = "hello, again."

	//Pull that value out of the session
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip") // if nothing in String is empty
		
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// grab the remote IP address of the person visiting my site and store it in the home page.
// in the home handler
