package controllers

import (
	"apigo/back/data"
	u "apigo/back/utils"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// CreateUser : new user controller
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &data.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	// If Admin creation, we check if the current user is an admin
	if user.AccessLevel != 0 {
		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header
		if tokenHeader == "" {                       //Token is missing, returns with error code 403 Unauthorized
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &data.Token{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		if err != nil { //Malformed token, returns with http code 403 as usual
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		if tk.AccessLevel != 0 {
			resp := user.Create() //Create Admin
			u.Respond(w, resp)
			return
		} else if tk.AccessLevel == 0 {
			u.Respond(w, u.Message(false, "Error, you must have admin rights for this"))
			return
		}
	} else {
		resp := user.Create() //Create user
		u.Respond(w, resp)
		return
	}
}

//UpdateUser : update user's info
var UpdateUser = func(c *gin.Context) {
	w := c.Writer
	r := c.Request

	response := make(map[string]interface{})
	tokenHeader := r.Header.Get("Authorization") //Grab the token from the header
	if tokenHeader == "" {                       //Token is missing, returns with error code 403 Unauthorized
		response = u.Message(false, "Missing auth token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}
	splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		response = u.Message(false, "Invalid/Malformed auth token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}
	tokenPart := splitted[1] //Grab the token part, what we are truly interested in
	tk := &data.Token{}
	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})
	if err != nil { //Malformed token, returns with http code 403 as usual
		response = u.Message(false, "Malformed authentication token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}
	if !token.Valid { //Token is invalid, maybe not signed on this server
		response = u.Message(false, "Token is not valid.")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}
	//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
	if tk.AccessLevel == 0 {
		u.Respond(c.Writer, u.Message(false, "Error, you must have admin rights for this"))
		return
	}
	uuidParam, err := uuid.FromString(c.Params.ByName("uuid"))
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "Can't find specified uuid"))
		return
	}
	//Get with password to update everything
	user := data.GetUserWithPW(uuidParam)// Get user by uuid
	if user.Uuid.UUID != nil { // Assign every param to the newUser
		newUser := &data.User{}
		newUser.ID = user.ID
		newUser.LastName = user.LastName
		newUser.Uuid.UUID = user.Uuid.UUID
		newUser.Password = user.Password
		newUser.AccessLevel = user.AccessLevel
		newUser.FirstName = user.FirstName
		newUser.DateOfBirth = user.DateOfBirth
		newUser.ID = user.ID
		temp := &data.User{}
		err := json.NewDecoder(c.Request.Body).Decode(temp) //decode the request body into struct and failed if any error occur
		if err != nil {
			u.Respond(c.Writer, u.Message(false, "Errors in user parameters"))
			return
		}
		// Verifiy variables that can be modified
		if temp.Email != "" { 
			newUser.Email = temp.Email
		}
		if temp.Password != "" {
			newUser.Password = temp.Password
		}
		if temp.LastName != "" {
			newUser.LastName = temp.LastName
		}
		if temp.FirstName != "" {
			newUser.FirstName = temp.FirstName
		}
		resp := newUser.Update()
		// Display modified data in JSON message "success"
		c.JSON(200, resp)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

//DeleteUser : Delete a user
var DeleteUser = func(c *gin.Context) {
	contextUser := c.Request.Context().Value("user") // Get the user from gin context
	if contextUser.(map[string]string)["access_level"] == "0" {
		u.Respond(c.Writer, u.Message(false, "Error, you must have admin rights for this"))
		return
	}
	uuidParam, err := uuid.FromString(c.Params.ByName("uuid")) // Get user from uuid
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "Can't find specified uuid"))
		return
	}
	//Every steps to delete user
	user := &data.User{}
	user.ID = data.GetUser(uuidParam).ID
	if user.ID != 0 {
		// Delete user
		resp := user.Delete()
		// Display modified data in JSON message "success"
		c.JSON(200, resp)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"}) 
	}
}

// Authenticate : login user controller
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &data.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := data.Login(user.Email, user.Password, r.RemoteAddr)
	u.Respond(w, resp)
}