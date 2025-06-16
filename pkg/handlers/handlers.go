package handlers

import (
	"net/http"

	"github.com/pauljc-probeplus/bookings/pkg/config"
	"github.com/pauljc-probeplus/bookings/pkg/models"
	"github.com/pauljc-probeplus/bookings/pkg/render"
)

//Repository struct
type Repository struct{
	App *config.AppConfig
}

//Repository variable
var Repo *Repository

//Creates a new repository
func NewRepo(a *config.AppConfig)*Repository{
	return &Repository{
		App:a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo=r
}

func (m *Repository)Home(w http.ResponseWriter, r *http.Request){
	remoteIp:=r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIp)
	render.RenderTemplate(w,"home.page.html",&models.TemplateData{})
}

func (m *Repository)About(w http.ResponseWriter,r *http.Request){
	
	//perform some logic
	stringMap:=make(map[string]string)
	stringMap["test"]="Hello, again."

	remoteIp:=m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"]=remoteIp

	
	render.RenderTemplate(w,"about.page.html",&models.TemplateData{StringMap: stringMap})
}
