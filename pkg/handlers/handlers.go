package handlers

import (
	"Hotel/pkg/config"
	"Hotel/pkg/models"
	"Hotel/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.Appconfig
}

func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the Home page handler
// in order for a function to respond to a request from a web, it has to handle two params
// because of the repository, we change this function with a new receiver, so all these function can have the access to repo
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
