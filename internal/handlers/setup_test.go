package handlers

import (
	"fmt"
	"encoding/gob"
	"net/http"
	"time"
	"github.com/alexedwards/scs/v2"
	"github.com/pauljc-probeplus/bookings/internal/config"
	"github.com/pauljc-probeplus/bookings/internal/models"
	"github.com/pauljc-probeplus/bookings/internal/render"
	"log"
	"html/template"
	"github.com/justinas/nosurf"
	"path/filepath"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates="./../../templates"
var functions=template.FuncMap{}


func getRoutes() http.Handler{
	
	gob.Register(models.Reservation{})
	 
	// change this to true when in production
	app.InProduction=false
		
	session = scs.New()
	session.Lifetime = 24* time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	
	app.Session=session
	
		
	template_cache,err:= CreateTestTemplateCache()
	if err!=nil{
		log.Fatal("cannot create template cache")
		
	}
		
	app.TemplateCache=template_cache
		
	app.UseCache=true
		
	
	repo:=NewRepo(&app)
	NewHandlers(repo)
		
	render.NewTemplates(&app)

	mux:=chi.NewRouter()
	
	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/",http.HandlerFunc(Repo.Home))
	mux.Get("/about",http.HandlerFunc(Repo.About))
	mux.Get("/generals-quarters",http.HandlerFunc(Repo.Generals))
	mux.Get("/majors-suite",http.HandlerFunc(Repo.Majors))
	
	mux.Get("/search-availibility",http.HandlerFunc(Repo.Availibility))
	mux.Post("/search-availibility",http.HandlerFunc(Repo.PostAvailibility))
	mux.Post("/search-availibility-json",http.HandlerFunc(Repo.AvailibilityJSON))
	
	
	mux.Get("/contact",http.HandlerFunc(Repo.Contact))

	mux.Get("/make-reservation",http.HandlerFunc(Repo.Reservation))
	mux.Post("/make-reservation",http.HandlerFunc(Repo.PostReservation))
	mux.Get("/reservation-summary",http.HandlerFunc(Repo.ReservationSummary))



	fileServer:=http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*",http.StripPrefix("/static",fileServer))


	return mux

}

func NoSurf(next http.Handler) http.Handler{
	csrfHandler:=nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler)http.Handler{
	return session.LoadAndSave(next)
}

func CreateTestTemplateCache()(map[string]*template.Template,error){
	myCache:=map[string]*template.Template{}

	//get all of the files named .page.html from templates folder
	pages,err:=filepath.Glob(fmt.Sprintf("%s/*.page.html",pathToTemplates))
	if err!=nil{
		return myCache,err
	}

	//range through all files ending with *.page.html
	for _,page:=range(pages){
		name:=filepath.Base(page)
		template_set,err:=template.New(name).ParseFiles(page)
		if err!=nil{
			return myCache,err
		}

		matches,err:=filepath.Glob(fmt.Sprintf("%s/*.layout.html",pathToTemplates))
		if err!=nil{
			return myCache,err
		}

		if len(matches)>0{
			template_set,err =template_set.ParseGlob(fmt.Sprintf("%s/*.layout.html",pathToTemplates))
			if err!=nil{
				return myCache,err
			}
		}
		myCache[name]=template_set
	}
	return myCache,err
}