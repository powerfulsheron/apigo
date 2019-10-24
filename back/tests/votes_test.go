package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"apigo/back/controllers"
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

	/*===========================
			BEFORE TEST
	===========================*/
	// REGISTER
	userMock := UserMock{}
	userMock.AccessLevel = 0
	errUserMock := faker.FakeData(&userMock)
	if errUserMock != nil {
		fmt.Println(errUserMock)
	}

	var jsonStrUser = []byte(`{"email":"` + userMock.Email + `","pass":"` + userMock.Password + `"}`)

	req, ererrReqRegister := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStrUser))
	if ererrReqRegister != nil {
		t.Fatal(ererrReqRegister)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// LOGIN
	createReq, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStrUser))
	if err != nil {
		t.Fatal(err)
	}
	createReq.Header.Set("Content-Type", "application/json")
	createRecorder := httptest.NewRecorder()
	createHandler := http.HandlerFunc(controllers.CreateUser)
	createHandler.ServeHTTP(createRecorder, createReq)
	if status := createRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	loginReq, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStrUser))
	if err != nil {
		t.Fatal(err)
	}
	loginReq.Header.Set("Content-Type", "application/json")
	loginRecorder := httptest.NewRecorder()
	loginHandler := http.HandlerFunc(controllers.Authenticate)
	loginHandler.ServeHTTP(loginRecorder, loginReq)
	if status := loginRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseBodyArrayUser map[string]interface{}
	if err := json.Unmarshal(loginRecorder.Body.Bytes(), &responseBodyArrayUser); err != nil {
		t.Fatal(err)
	}

	var token = responseBodyArrayUser["jwt"]
	/*===========================
				TEST
	===========================*/
	// Faker
	voteMock := VoteMock{}
	errVoteMock := faker.FakeData(&voteMock)
	if errVoteMock != nil {
		fmt.Println(errVoteMock)
	}

	// Request
	router := router.VoteRouter()
	var jsonStr = []byte(`{"title":"` + voteMock.Title + `","description":"` + voteMock.Description + `"}`)
	w := performRequest(router, "POST", "/votes", jsonStr, token.(string))

	// Response status
	if status := w.Code; status != 401 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
