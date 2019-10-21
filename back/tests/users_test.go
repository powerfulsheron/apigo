package test

import (
	"apigo/back/controllers"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/bxcodec/faker"
	"fmt"
	"encoding/json"
)


type UserMock struct {
	Email       string    `faker:"email"`
	Password    string    `faker:"password"`
	AccessLevel int 
	FirstName   string    `faker:"first_name"`
	LastName    string    `faker:"last_name"`
}

func TestCreateUser(t *testing.T) { 

	userMock := UserMock{}
	userMock.AccessLevel = 0
	err := faker.FakeData(&userMock)
	if err != nil {
		fmt.Println(err)
	}
	
	var jsonStr = []byte(`{"email":"`+userMock.Email+`","pass":"`+userMock.Password+`"}`)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseBodyArray map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBodyArray); err != nil {
		t.Fatal(err)
	}

	var userArray map[string]interface{}
	userArray = responseBodyArray["user"].(map[string]interface{})

	got := `{"email":"`+userArray["email"].(string)+`","pass":"`+userArray["pass"].(string)+`"}`
	expected := `{"email":"`+userMock.Email+`","pass":""}`

	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expected)
	}
}

func TestUserLogin(t *testing.T) { 

	userMock := UserMock{}
	userMock.AccessLevel = 0
	err := faker.FakeData(&userMock)
	if err != nil {
		fmt.Println(err)
	}
	
	var jsonStr = []byte(`{"email":"`+userMock.Email+`","pass":"`+userMock.Password+`"}`)

	createReq, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
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

	loginReq, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
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

	var responseBodyArray map[string]interface{}
	if err := json.Unmarshal(loginRecorder.Body.Bytes(), &responseBodyArray); err != nil {
		t.Fatal(err)
	}

	got := `{"message":"`+responseBodyArray["message"].(string)+`","status":`+ fmt.Sprintf("%t", responseBodyArray["status"].(bool))+`}`
	expected := `{"message":"Logged In","status":true}`

	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expected)
	}
}
