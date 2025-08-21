package main

import (
	//"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pauljc-probeplus/bookings/internal/config"
	"github.com/pauljc-probeplus/bookings/internal/handlers"
	"github.com/pauljc-probeplus/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber=": 8082"
var app config.AppConfig
var session *scs.SessionManager


func main(){
	
	 
	// change this to true when in production
	app.InProduction=false
	
	session = scs.New()
	session.Lifetime = 24* time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session=session

	
	template_cache,err:= render.CreateTemplateCache()
	if err!=nil{
		log.Fatal("cannot create template cache")
	}
	
	app.TemplateCache=template_cache
	
	app.UseCache=false
	

	repo:=handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	
	render.NewTemplates(&app)
	
	
	fmt.Println("Starting server at port",portNumber)
	
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),

	}
	err = srv.ListenAndServe()
	log.Fatal(err)



}