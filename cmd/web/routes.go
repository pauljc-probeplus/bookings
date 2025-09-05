package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pauljc-probeplus/bookings/internal/config"
	"github.com/pauljc-probeplus/bookings/internal/handlers"
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
	mux.Get("/generals-quarters",http.HandlerFunc(handlers.Repo.Generals))
	mux.Get("/majors-suite",http.HandlerFunc(handlers.Repo.Majors))
	
	mux.Get("/search-availibility",http.HandlerFunc(handlers.Repo.Availibility))
	mux.Post("/search-availibility",http.HandlerFunc(handlers.Repo.PostAvailibility))
	mux.Post("/search-availibility-json",http.HandlerFunc(handlers.Repo.AvailibilityJSON))
	
	
	mux.Get("/contact",http.HandlerFunc(handlers.Repo.Contact))

	mux.Get("/make-reservation",http.HandlerFunc(handlers.Repo.Reservation))
	mux.Post("/make-reservation",http.HandlerFunc(handlers.Repo.PostReservation))
	mux.Get("/reservation-summary",http.HandlerFunc(handlers.Repo.ReservationSummary))



	fileServer:=http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*",http.StripPrefix("/static",fileServer))


	return mux

}