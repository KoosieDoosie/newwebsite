package handlers

import (
	"github.com/KoosieDoosie/newwebsite/pkg/config"
	"github.com/KoosieDoosie/newwebsite/pkg/models"
	"github.com/KoosieDoosie/newwebsite/pkg/render"
	"net/http"
)

// The respository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repsitory
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Newhandler sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "Remote_IP", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	remoteIP := m.App.Session.GetString(r.Context(), "Remote_IP")

	stringMap["Remote_IP"] = remoteIP

	//send the data to the template

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
