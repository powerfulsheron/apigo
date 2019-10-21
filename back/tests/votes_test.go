package test

import (
	"apigo/back/controllers"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// POST /votes
func TestCreateVote(t *testing.T) {
	var jsonStr = []byte(`{"Title": "Brexit 2.0","Description": "Le vote du Brexit 2.0"}`)

	req, err := http.NewRequest("POST", "/votes", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateVote)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

// GET /votes
func TestGetVotes(t *testing.T) {
	req, err := http.NewRequest("GET", "/votes", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetVotes)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}
