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

	expected := `{"L":0.6667374130409149,"Lq":0.00007074637424831977,"W":0.33336870652045747,"Wq":0.000035373187124159885}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body:\n got:\n%s\n want:\n%s", rr.Body.String(), expected)
	}
}
