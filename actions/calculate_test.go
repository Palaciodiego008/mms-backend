package actions_test

import (
	"mms-project/actions"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	payload := `{"lambda": 2.0, "mu": 3.0, "s": 4.0}`
	req, err := http.NewRequest("POST", "/calculate", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(actions.CalculateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Average Number of Customers in the System (L): 0.666737\n" +
		"Average Number of Customers in the Queue (Lq): 0.000071\n" +
		"Average Time a Customer Spends in the System (W): 0.333369\n" +
		"Average Time a Customer Spends in the Queue (Wq): 0.000035\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body:\n got:\n%s\n want:\n%s", rr.Body.String(), expected)
	}
}