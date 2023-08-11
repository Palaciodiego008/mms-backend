package main

import (
	"mms-project/actions"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/calculate", actions.CalculateHandler)

	http.ListenAndServe(":8080", r)

}
