package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) mount() http.Handler {
	// Create a new instance of chi router
	mux := chi.NewRouter()

	// Middleware
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	// Routes for v1
	mux.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthHandler)
	})

	return mux
}

func (app *application) run() error {
	// Call the mount() method to create the routes and middleware for the application.
	mux := app.mount()

	// Create a new http.Server struct to hold the server configuration settings.
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Inform the user that the server is starting.
	log.Printf("Starting server on %s\n", srv.Addr)

	// Run the server
	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
