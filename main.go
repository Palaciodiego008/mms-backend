package main

import (
	"fmt"
	"mms-project/actions"
	"net/http"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	port := 8080
	r := chi.NewRouter()
	// Configure CORS settings
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodPost},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Use the CORS middleware
	r.Use(corsOptions.Handler)

	// Use some basic middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/calculate", actions.CalculateHandler)
	r.Post("/general-stats", actions.GeneralStatsHandler) // New endpoint for general statistics
	r.Post("/experiment", actions.ExperimentHandler)      // New endpoint for parameter experiments

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
