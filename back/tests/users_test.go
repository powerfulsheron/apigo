package test

import (
	"apigo/back/controllers"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/bxcodec/faker"
	"os"
	"apigo/back/router"
)

type UserMock struct {
	Email       string `faker:"email"`
	Password    string `faker:"password"`
	AccessLevel int
	FirstName   string `faker:"first_name"`
	LastName    string `faker:"last_name"`
}



func performRequestWithJwt(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("admin_jwt_for_test"))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestCreateUser(t *testing.T) { 

	userMock := UserMock{}
	userMock.AccessLevel = 0
	err := faker.FakeData(&userMock)
	if err != nil {
		fmt.Println(err)
	}

	var jsonStr = []byte(`{"email":"` + userMock.Email + `","pass":"` + userMock.Password + `"}`)

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

	got := `{"email":"` + userArray["email"].(string) + `","pass":"` + userArray["pass"].(string) + `"}`
	expected := `{"email":"` + userMock.Email + `","pass":""}`

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

	var jsonStr = []byte(`{"email":"` + userMock.Email + `","pass":"` + userMock.Password + `"}`)

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

	got := `{"message":"` + responseBodyArray["message"].(string) + `","status":` + fmt.Sprintf("%t", responseBodyArray["status"].(bool)) + `}`
	expected := `{"message":"Logged In","status":true}`

	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expected)
	}
}

func TestModifyUser(t *testing.T) { 

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

	var userArray map[string]interface{}
	if userArrayErr := json.Unmarshal(createRecorder.Body.Bytes(), &userArray); err != nil {
		t.Fatal(userArrayErr)
	}
	userArray = userArray["user"].(map[string]interface{})

	newUserMock := UserMock{}
	newUserMock.AccessLevel = 0
	errNewUser := faker.FakeData(&newUserMock)
	if errNewUser != nil {
		fmt.Println(errNewUser)
	}

	// Request
	router := router.VoteRouter()
	jsonStr = []byte(`{"last_name":"`+newUserMock.LastName+`","first_name":"`+newUserMock.FirstName+`","email":"`+newUserMock.Email+`"}`)
	modifyRecorder := performRequestWithJwt(router, "PUT", "/users/"+userArray["uuid"].(string), jsonStr)

	// Response status
	if status := modifyRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// modifyReq, modifyReqErr := http.NewRequest("PUT", "/users/"+userArray["uuid"].(string), bytes.NewBuffer(jsonStr))
	// if modifyReqErr != nil {
	// 	t.Fatal(modifyReqErr)
	// }
	// modifyReq.Header.Set("Content-Type", "application/json")
	// modifyReq.Header.Set("Authorization", "Bearer "+os.Getenv("admin_jwt_for_test"))
	// modifyRecorder := httptest.NewRecorder()
	// modifyHandler := http.HandlerFunc(gin.Wrap(controllers.UpdateUser))
	// modifyHandler.ServeHTTP(modifyRecorder, modifyReq)
	// if status := modifyRecorder.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v",
	// 		status, http.StatusOK)
	// }

	var responseModifiedUserArray map[string]interface{}
	if modifiedUserArrayErr := json.Unmarshal(modifyRecorder.Body.Bytes(), &responseModifiedUserArray); modifiedUserArrayErr != nil {
		t.Fatal(modifiedUserArrayErr)
	}

	var modifiedUserArray map[string]interface{}
	modifiedUserArray = responseModifiedUserArray["user"].(map[string]interface{})

	got := `{"last_name":"`+modifiedUserArray["last_name"].(string)+`","first_name":"`+modifiedUserArray["first_name"].(string)+`","email":"`+modifiedUserArray["email"].(string)+`"}`
	expected := `{"last_name":"`+newUserMock.LastName+`","first_name":"`+newUserMock.FirstName+`","email":"`+newUserMock.Email+`"}`

	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expected)
	}
}

func TestDeleteUser(t *testing.T) { 

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

	var userArray map[string]interface{}
	if userArrayErr := json.Unmarshal(createRecorder.Body.Bytes(), &userArray); err != nil {
		t.Fatal(userArrayErr)
	}
	userArray = userArray["user"].(map[string]interface{})

	newUserMock := UserMock{}
	newUserMock.AccessLevel = 0
	errNewUser := faker.FakeData(&newUserMock)
	if errNewUser != nil {
		fmt.Println(errNewUser)
	}

	// Request
	router := router.VoteRouter()
	jsonStr = []byte(`{"last_name":"`+newUserMock.LastName+`","first_name":"`+newUserMock.FirstName+`","email":"`+newUserMock.Email+`"}`)
	modifyRecorder := performRequestWithJwt(router, "DELETE", "/users/"+userArray["uuid"].(string),[]byte(``))

	// Response status
	if status := modifyRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var responseModifiedUserArray map[string]interface{}
	if modifiedUserArrayErr := json.Unmarshal(modifyRecorder.Body.Bytes(), &responseModifiedUserArray); modifiedUserArrayErr != nil {
		t.Fatal(modifiedUserArrayErr)
	}

	var modifiedUserArray map[string]interface{}
	modifiedUserArray = responseModifiedUserArray["user"].(map[string]interface{})

	got := `{"message":"`+responseModifiedUserArray["message"].(string)+`","status":`+fmt.Sprintf("%t", responseModifiedUserArray["status"].(bool))+`,"ID":`+fmt.Sprintf("%g", modifiedUserArray["ID"].(float64))+`}`
	expected := `{"message":"User has been Deleted","status":true,"ID":`+fmt.Sprintf("%g", userArray["ID"].(float64))+`}`

	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", got, expected)
	}
}