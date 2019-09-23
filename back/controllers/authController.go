package controllers

import (
	"apigo/back/data"
	u "apigo/back/utils"
	"encoding/json"
	"net/http"
)

// CreateAccount : new user controller
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &data.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

// Authenticate : login user controller
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &data.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := data.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
