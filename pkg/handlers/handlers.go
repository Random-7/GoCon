package handlers

import (
	"fmt"
	"net/http"

	"github.com/Random-7/GoRcon/pkg/config"
	"github.com/Random-7/GoRcon/pkg/models"
	"github.com/Random-7/GoRcon/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

//Repository is the struct type
type Repository struct {
	App *config.AppConfig
}

//Creates the new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

//NewHandlers Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home renders the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.go.tmpl", &models.TemplateData{ActivePage: "Home"})
}

//Dashboard renders the Dashboard page - containing data pulled from RCON
func (m *Repository) Dashboard(w http.ResponseWriter, r *http.Request) {
	playercount, playerlist, err := Repo.App.Rcon.GetPlayers()
	if err != nil {
		fmt.Println("Error with loading player list", err)
	}

	stringMap := make(map[string]string)
	stringMap["playercount"] = fmt.Sprintf("%d", playercount)

	data := make(map[string]interface{})
	data["players"] = playerlist
	data["rconStatus"] = m.App.Rcon.ConnectionStatus

	render.RenderTemplate(w, "dashboard.page.go.tmpl", &models.TemplateData{ActivePage: "Dashboard",
		StringMap: stringMap, DataMap: data})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.go.tmpl", &models.TemplateData{ActivePage: "About"})
}

//Players renders the Players page - containing data pulled from RCON
func (m *Repository) Players(w http.ResponseWriter, r *http.Request) {
	playercount, playerlist, err := Repo.App.Rcon.GetPlayers()
	if err != nil {
		fmt.Println("Error with loading player list", err)
	}

	stringMap := make(map[string]string)
	stringMap["playercount"] = fmt.Sprintf("%d", playercount)

	data := make(map[string]interface{})
	data["players"] = playerlist

	render.RenderTemplate(w, "players.page.go.tmpl", &models.TemplateData{ActivePage: "Players",
		StringMap: stringMap, DataMap: data})
}

func (m *Repository) Commands(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "commands.page.go.tmpl", &models.TemplateData{ActivePage: "Commands"})
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.go.tmpl", &models.TemplateData{ActivePage: "Login"})
}

func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	//redirect to "/"
}
func (m *Repository) Admin(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "admin.page.go.tmpl", &models.TemplateData{ActivePage: "Admin"})
}
