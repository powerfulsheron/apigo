package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"apigo/back/router"

	"github.com/bxcodec/faker"
)

type VoteMock struct {
	Title       string `faker:"sentence"`
	Description string `faker:"paragraph"`
}

// Perform request
func performRequest(r http.Handler, method, path string, body []byte, jwt string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+jwt)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// POST /votes
func TestCreateVote(t *testing.T) {

	var token = os.Getenv("admin_jwt_for_test")

	// Faker
	voteMock := VoteMock{}
	errVoteMock := faker.FakeData(&voteMock)
	if errVoteMock != nil {
		fmt.Println(errVoteMock)
	}

	// Request
	router := router.VoteRouter()
	var jsonStr = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	w := performRequest(router, "POST", "/votes", jsonStr, token)

	// Response status
	if status := w.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestVoteVote(t *testing.T) {

	var token = os.Getenv("admin_jwt_for_test")

	// Faker
	voteMock := VoteMock{}
	errVoteMock := faker.FakeData(&voteMock)
	if errVoteMock != nil {
		fmt.Println(errVoteMock)
	}

	// Request
	router := router.VoteRouter()
	var jsonStrNewVote = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	w := performRequest(router, "POST", "/votes", jsonStrNewVote, token)

	// Response status
	if status := w.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Recupération du vote
	var response map[string]interface{}
	if responseError := json.Unmarshal(w.Body.Bytes(), &response); responseError != nil {
		t.Fatal(responseError)
	}

	var innerResponse map[string]interface{}
	innerResponse = response["response"].(map[string]interface{})

	var voteResponse map[string]interface{}
	voteResponse = innerResponse["vote"].(map[string]interface{})

	// Request
	var jsonStr = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	voteRequest := performRequest(router, "POST", "/votes/"+voteResponse["uuid"].(string), jsonStr, token)

	// Response status
	if status := voteRequest.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDeleteVote(t *testing.T) {

	var token = os.Getenv("admin_jwt_for_test")

	// Faker
	voteMock := VoteMock{}
	errVoteMock := faker.FakeData(&voteMock)
	if errVoteMock != nil {
		fmt.Println(errVoteMock)
	}

	// Request
	router := router.VoteRouter()
	var jsonStrNewVote = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	w := performRequest(router, "POST", "/votes", jsonStrNewVote, token)

	// Response status
	if status := w.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Recupération du vote
	var response map[string]interface{}
	if responseError := json.Unmarshal(w.Body.Bytes(), &response); responseError != nil {
		t.Fatal(responseError)
	}

	var innerResponse map[string]interface{}
	innerResponse = response["response"].(map[string]interface{})

	var voteResponse map[string]interface{}
	voteResponse = innerResponse["vote"].(map[string]interface{})

	// Request
	var jsonStr = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	voteRequest := performRequest(router, "DELETE", "/votes/"+voteResponse["uuid"].(string), jsonStr, token)

	// Response status
	if status := voteRequest.Code; status != 204 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdateVote(t *testing.T) {

	var token = os.Getenv("admin_jwt_for_test")

	// Faker
	voteMock := VoteMock{}
	errVoteMock := faker.FakeData(&voteMock)
	if errVoteMock != nil {
		fmt.Println(errVoteMock)
	}

	// Request
	router := router.VoteRouter()
	var jsonStrNewVote = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	w := performRequest(router, "POST", "/votes", jsonStrNewVote, token)

	// Response status
	if status := w.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Recupération du vote
	var response map[string]interface{}
	if responseError := json.Unmarshal(w.Body.Bytes(), &response); responseError != nil {
		t.Fatal(responseError)
	}

	var innerResponse map[string]interface{}
	innerResponse = response["response"].(map[string]interface{})

	var voteResponse map[string]interface{}
	voteResponse = innerResponse["vote"].(map[string]interface{})

	// Request
	var jsonStr = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	voteRequest := performRequest(router, "PUT", "/votes/"+voteResponse["uuid"].(string), jsonStr, token)

	// Response status
	if status := voteRequest.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestGetVote(t *testing.T) {

	var token = os.Getenv("admin_jwt_for_test")

	// Faker
	voteMock := VoteMock{}
	errVoteMock := faker.FakeData(&voteMock)
	if errVoteMock != nil {
		fmt.Println(errVoteMock)
	}

	// Request
	router := router.VoteRouter()
	var jsonStrNewVote = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	w := performRequest(router, "POST", "/votes", jsonStrNewVote, token)

	// Response status
	if status := w.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Recupération du vote
	var response map[string]interface{}
	if responseError := json.Unmarshal(w.Body.Bytes(), &response); responseError != nil {
		t.Fatal(responseError)
	}

	var innerResponse map[string]interface{}
	innerResponse = response["response"].(map[string]interface{})

	var voteResponse map[string]interface{}
	voteResponse = innerResponse["vote"].(map[string]interface{})

	// Request
	var jsonStr = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	voteRequest := performRequest(router, "GET", "/votes/"+voteResponse["uuid"].(string), jsonStr, token)

	// Response status
	if status := voteRequest.Code; status != 200 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
