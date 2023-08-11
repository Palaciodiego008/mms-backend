package actions

import (
	"encoding/json"
	"fmt"
	"mms-project/internal"
	"net/http"
)

type InputData struct {
	Lambda float64 `json:"lambda"`
	Mu     float64 `json:"mu"`
	S      float64 `json:"s"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintf(w, "Average Number of Customers in the System (L): %.6f\n", L)
	fmt.Fprintf(w, "Average Number of Customers in the Queue (Lq): %.6f\n", Lq)
	fmt.Fprintf(w, "Average Time a Customer Spends in the System (W): %.6f\n", W)
	fmt.Fprintf(w, "Average Time a Customer Spends in the Queue (Wq): %.6f\n", Wq)
}
