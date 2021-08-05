package handlers

import (
	"github.com/Fietzorama/Bookings/pkg/config"
	"github.com/Fietzorama/Bookings/pkg/models"
	"github.com/Fietzorama/Bookings/pkg/render"
	"net/http"
)

// GENERAL:
// Handlers reading a template from disk storing in the template cache.
// And every time the appropriate handler is called, I serve the appropriate template

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository Type
// repository pattern - common pattern
// that allows it allows us to swap components
// out of our application with a minimal changes required to the code base.
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
// NewRepo takes "a" as a pointer to the application config
// NewRepo returns a pointer to a repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{ //returns the referenc to a Repository
		App: a, // populate in "App" from type "Repository
	}
}

// NewHandlers sets the repository for the handlers
// NewHandlers create new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home Handler respond to the request to the main page
// m *Repository links all of these functions,
// all handlers together with repository
// all handlers have access to that repository and inside the repo.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// Every time somebody hits that home page for that user session
	remoteIP := r.RemoteAddr

	//Storing the remote IP Address as a string in the session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	//TemplateData passing a empty Template, else it would be error
	render.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{})
}

// About Handler respond to request to the about page
// m *Repository links all of these functions,
// all handlers together with repository
// all handlers have access to that repository and inside the repo.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// Create a new StringMap (see TemplateData.go)
	stringMap := make(map[string]string)

	// Assign the value of test "hello, again"
	stringMap["Test"] = "hello, again."

	//Pull IP out of the session
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP // from TemplateData

	// send the data to the template
	// TemplateData need a value
	render.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
