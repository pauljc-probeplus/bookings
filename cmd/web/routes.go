package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pauljc-probeplus/bookings/pkg/config"
	"github.com/pauljc-probeplus/bookings/pkg/handlers"
)

func routes(app *config.AppConfig)http.Handler{
	
	/*mux:=pat.New()
	mux.Get("/",http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about",http.HandlerFunc(handlers.Repo.About))*/

	mux:=chi.NewRouter()
	
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/",http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about",http.HandlerFunc(handlers.Repo.About))

	fileServer:=http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*",http.StripPrefix("/static",fileServer))


	return mux

}