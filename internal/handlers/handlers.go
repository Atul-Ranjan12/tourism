package handlers

import (
	"net/http"

	"github.com/Atul-Ranjan12/tourism/internal/config"
	"github.com/Atul-Ranjan12/tourism/internal/models"
	"github.com/Atul-Ranjan12/tourism/internal/render"
)

// Initialize the repository for the application
type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Function to create a new test Repository
func NewTestingRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) ShowLogin(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "login.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ShowSignUp(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "signup.page.tmpl", &models.TemplateData{})
}
