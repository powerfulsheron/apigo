package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"apigo/back/router"

	"github.com/bxcodec/faker"
)

type VoteMock struct {
	Title       string `faker:"sentence"`
	Description string `faker:"paragraph"`
}

// Perform request
func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// POST /votes
func TestCreateVote(t *testing.T) {

	// Faker
	voteMock := VoteMock{}
	err := faker.FakeData(&voteMock)
	if err != nil {
		fmt.Println(err)
	}

	// Request
	router := router.VoteRouter()
	var jsonStr = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	w := performRequest(router, "POST", "/votes", jsonStr)

	// Response status
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Response body
	var responseBodyArray map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBodyArray); err != nil {
		t.Fatal(err)
	}

	var responseArray map[string]interface{}
	responseArray = responseBodyArray["success"].(map[string]interface{})

	var voteArray map[string]interface{}
	voteArray = responseArray["vote"].(map[string]interface{})

	// Grab the value & whether or not it exists
	title, titleExists := voteArray["vote"]
	description, descExists := voteArray["description"]

	if titleExists == false || title != voteMock.Title {
		t.Errorf("Nil title or wrong value")
	}

	if descExists == false || description != voteMock.Description {
		t.Errorf("Nil title or wrong value")
	}
}
