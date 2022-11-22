package handlers

import (
	"fmt"
	"net/http"

	"github.com/atuprosper/go-project/pkg/config"
	"github.com/atuprosper/go-project/pkg/models"
	"github.com/atuprosper/go-project/pkg/render"
)

// Creating a Repository pattern
// This variable is the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// This function creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// This function NewHandlers, sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// A function with a 'reciever' m, of type 'Repository'. This will give our handler function access to everything in the config file
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// Get the IP address of user
	remoteIP := r.RemoteAddr
	// store the IP in the site wide config which is available via the paramter 'm'
	// It takes in three parameters. The context, the name (which can be anything), and the value
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	//Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	// Get the remote IP and pass it to the template
	getRemoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = getRemoteIP
	fmt.Println("Your IP address is", getRemoteIP)

	// Send the data to the template
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}
