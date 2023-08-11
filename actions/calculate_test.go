package actions_test

import (
	"bytes"
	"mms-project/actions"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	payload := "2.0\n3.0\n4.0\n"
	req, err := http.NewRequest("GET", "/calculate", bytes.NewBufferString(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(actions.CalculateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Average Number of Customers in the System (L): 2.666667\n" +
		"Average Number of Customers in the Queue (Lq): 1.333333\n" +
		"Average Time a Customer Spends in the System (W): 1.166667\n" +
		"Average Time a Customer Spends in the Queue (Wq): 0.583333\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
