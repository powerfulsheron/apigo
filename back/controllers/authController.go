package controllers

import (
	"apigo/back/data"
	u "apigo/back/utils"
	"encoding/json"
	"net/http"
)

// CreateUser : new user controller
var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &data.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Create() //Create user
	u.Respond(w, resp)
}

// Authenticate : login user controller
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &data.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := data.Login(user.Email, user.Password)
	u.Respond(w, resp)
}
