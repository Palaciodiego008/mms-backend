package actions

import (
	"encoding/json"
	"mms-project/internal"
	"net/http"
)

type InputData struct {
	Lambda float64 `json:"lambda"`
	Mu     float64 `json:"mu"`
	S      float64 `json:"s"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	// Checking if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var inputData InputData

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inputData); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Validating input data (lambda, mu, s)
	if inputData.Lambda <= 0 || inputData.Mu <= 0 || inputData.S <= 0 {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	L, Lq, W, Wq := internal.MmsModel(inputData.Lambda, inputData.Mu, inputData.S)

	// Create a map to hold the results
	results := map[string]float64{
		"L":  L,
		"Lq": Lq,
		"W":  W,
		"Wq": Wq,
	}

	// Convert the results map to JSON
	responseJSON, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Write the JSON response to the writer
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
