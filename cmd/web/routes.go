package main

import (
	"net/http"

	"github.com/agniadvani/bookings/pkg/config"
	handler "github.com/agniadvani/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//Using Chi Router
func router(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/contact", handler.Repo.Contact)
	mux.Get("/majors-suite", handler.Repo.Majors)
	mux.Get("/presidential-suite",handler.Repo.Presidential)
	mux.Get("/search-availability", handler.Repo.SearchAvailability)
	mux.Get("/make-reservation", handler.Repo.MakeReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static",fileServer))
	return mux
}
