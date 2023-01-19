package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://0.0.0.0:3000", "*"},
	})
	mux.Use(corsMiddleware.Handler)

	mux.Get("/", app.Home)

	mux.Get("/movies", app.AllMovies)

	return mux
}
