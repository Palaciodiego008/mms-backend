package main

import (
	"fmt"
	"mms-project/actions"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Post("/calculate", actions.CalculateHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
