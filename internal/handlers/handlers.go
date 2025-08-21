package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pauljc-probeplus/bookings/internal/config"
	"github.com/pauljc-probeplus/bookings/internal/models"
	"github.com/pauljc-probeplus/bookings/internal/render"
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
	render.RenderTemplate(w,r,"home.page.html",&models.TemplateData{})
}

func (m *Repository)About(w http.ResponseWriter,r *http.Request){
	
	//perform some logic
	stringMap:=make(map[string]string)
	stringMap["test"]="Hello, again."

	remoteIp:=m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"]=remoteIp

	
	render.RenderTemplate(w,r,"about.page.html",&models.TemplateData{StringMap: stringMap})
}


// func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request){
// 	render.RenderTemplate(w,"make-reservation.html",&models.TemplateData{})
// }

//Generals render the room page
func (m *Repository) Generals(w http.ResponseWriter,r *http.Request){
	render.RenderTemplate(w,r,"generals.page.html",&models.TemplateData{})
}

//Majors render the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"majors.page.html",&models.TemplateData{})
}

//Availibility render the search availibility page
func (m *Repository) Availibility(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"search-availibility.page.html",&models.TemplateData{})
}

//PostAvailibility
func (m *Repository) PostAvailibility(w http.ResponseWriter, r *http.Request){
	start:= r.Form.Get("start")
	end:= r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s",start,end)))
}

type jsonResponse struct{
	OK bool `json:"ok"`
	Message string `json:"message"`
}
//AvailibilityJSON handles request for availibility and send JSON response
func (m *Repository)AvailibilityJSON(w http.ResponseWriter, r *http.Request){
	resp := jsonResponse {
		OK:true,
		Message:"Available!",
	}
	out,err:=json.MarshalIndent(resp,"","     ")
	if err !=nil{
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type","application/json")
	w.Write(out)
}

//Contact render the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"contact.page.html",&models.TemplateData{})
}

//Reservation renders the make a reservation page and displays the form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r,"reservation.page.html",&models.TemplateData{})
}
