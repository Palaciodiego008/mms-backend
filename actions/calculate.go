package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

// This handler is used to calculate additional statistics
func GeneralStatsHandler(w http.ResponseWriter, r *http.Request) {
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

	// Perform calculations for general statistics
	L, Lq, W, Wq := internal.MmsModel(inputData.Lambda, inputData.Mu, inputData.S)

	// Additional statistics calculation
	avgQueueLength := Lq / Wq
	avgSystemTime := 1 / W
	// Add more calculations as needed

	// Create a map to hold the results
	results := map[string]float64{
		"L":              L,
		"Lq":             Lq,
		"W":              W,
		"Wq":             Wq,
		"AvgQueueLength": avgQueueLength,
		"AvgSystemTime":  avgSystemTime,
		// Add more results as needed
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

func ExperimentHandler(w http.ResponseWriter, r *http.Request) {
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

	// Perform calculations for different parameter combinations
	results := make(map[string]map[string]float64)
	for lambda := inputData.Lambda - 1.0; lambda <= inputData.Lambda+1.0; lambda += 0.5 {
		for mu := inputData.Mu - 0.1; mu <= inputData.Mu+0.1; mu += 0.05 {
			for s := inputData.S - 1.0; s <= inputData.S+1.0; s += 0.5 {
				L, Lq, W, Wq := internal.MmsModel(lambda, mu, s)

				parameterCombination := fmt.Sprintf("λ=%.2f, μ=%.2f, s=%.2f", lambda, mu, s)
				results[parameterCombination] = map[string]float64{
					"L":  L,
					"Lq": Lq,
					"W":  W,
					"Wq": Wq,
				}
			}
		}
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

// Endpoint para recomendaciones automáticas
func AutoRecommendationHandler(w http.ResponseWriter, r *http.Request) {
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

	// Perform calculations and logic for auto recommendations
	// This can be based on specific rules or optimization algorithms

	recommendedLambda := inputData.Lambda * 1.5 // Example recommendation
	recommendedMu := inputData.Mu * 1.2         // Example recommendation
	recommendedS := inputData.S + 1.0           // Example recommendation

	// Create a map to hold the recommendation results
	recommendations := map[string]float64{
		"RecommendedLambda": recommendedLambda,
		"RecommendedMu":     recommendedMu,
		"RecommendedS":      recommendedS,
	}

	// Convert the recommendations map to JSON
	responseJSON, err := json.Marshal(recommendations)
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

func VisualizationHandler(w http.ResponseWriter, r *http.Request) {
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

	// Perform calculations and logic for visualization data
	// This could involve generating data for graphs or charts

	// Create example data for visualization
	// This data structure should match the format expected by your visualization library
	visualizationData := map[string]interface{}{
		"labels": []string{"Queue Length", "System Time"},
		"data":   []float64{10, 5}, // Example data
	}

	// Convert the visualization data to JSON
	responseJSON, err := json.Marshal(visualizationData)
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

func GenerateAndCalculateHandler(w http.ResponseWriter, r *http.Request) {
	// Checking if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var chatInput struct {
		Scenario string `json:"scenario"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&chatInput); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Checking if a scenario is provided
	if chatInput.Scenario == "" {
		http.Error(w, "No scenario provided", http.StatusBadRequest)
		return
	}

	// Perform an API call to ChatGPT to generate a response based on the scenario
	chatResponse, err := CallChatGPTAPI(chatInput.Scenario)
	if err != nil {
		http.Error(w, "Error calling ChatGPT API", http.StatusInternalServerError)
		return
	}

	// Assuming chatResponse is a JSON response from ChatGPT
	var chatResponseData struct {
		Response string `json:"response"`
	}
	err = json.Unmarshal([]byte(chatResponse), &chatResponseData)
	if err != nil {
		http.Error(w, "Error parsing ChatGPT response", http.StatusInternalServerError)
		return
	}

	// Parse the response from ChatGPT to extract relevant parameters for the MMS model
	// For this example, we'll assume some parameters
	lambda := 8.0
	mu := 10.0
	s := 3.0

	// Perform calculations using MMS model with the generated parameters
	L, Lq, W, Wq := internal.MmsModel(lambda, mu, s)

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

// CallChatGPTAPI realiza la llamada a la API de ChatGPT y devuelve la respuesta
func CallChatGPTAPI(scenario string) (string, error) {
	// URL de la API de ChatGPT
	chatGPTAPIURL := "https://api.chatgpt.com/v1/chat"
	apiKey := "TU_CLAVE_DE_API_AQUI" // Cambia esto a tu clave de API real

	// Crear el payload para la solicitud a la API
	requestData := map[string]interface{}{
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": scenario},
		},
	}

	// Convertir el payload a formato JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	// Crear la solicitud HTTP POST
	req, err := http.NewRequest("POST", chatGPTAPIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	// Agregar encabezados requeridos
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Realizar la solicitud HTTP
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var responseBody bytes.Buffer
	_, err = io.Copy(&responseBody, resp.Body)
	if err != nil {
		return "", err
	}

	return responseBody.String(), nil
}
