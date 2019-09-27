package controllers

import (
	u "apigo/back/utils"
	"encoding/json"
	"net/http"
	"apigo/back/data"
)

// GetVotes : GET a list of votes
var GetVotes = func(w http.ResponseWriter, r *http.Request) {
	votes := data.GetVotes() // Get votes
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(votes)
}

// CreateVote : POST a vote
var CreateVote = func(w http.ResponseWriter, r *http.Request) {

	vote := &data.Vote{}
	err := json.NewDecoder(r.Body).Decode(vote) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := vote.Create() //Create user
	u.Respond(w, resp)
}

/*
 --- NOT IMPLEMENTED ---
*/

// GetVote : GET a single vote
var GetVote = func(w http.ResponseWriter, r *http.Request) {

	// not implemented
}

// UpdateVote : PUT/UPDATE a vote
var UpdateVote = func(w http.ResponseWriter, r *http.Request) {

	// not implemented

}

// DeleteVote : DELETE a vote
var DeleteVote = func(w http.ResponseWriter, r *http.Request) {

	// not implemented
}
