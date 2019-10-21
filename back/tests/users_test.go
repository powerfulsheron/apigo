package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"apigo/back/controllers"
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
			rr.Body.String(), expected)
	}
}