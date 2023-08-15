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
	r := chi.NewRouter()
	// Configure CORS settings
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
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

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":3000", r)
}
