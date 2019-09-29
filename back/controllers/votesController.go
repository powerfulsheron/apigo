package controllers

import (
	"apigo/back/data"
	u "apigo/back/utils"
	"encoding/json"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetVotes : GET a list of votes
var GetVotes = func(c *gin.Context) {
	votes := data.GetVotes() // Get votes
	c.Writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(votes)
}

// CreateVote : POST a vote
var CreateVote = func(c *gin.Context) {

	contextUser := c.Request.Context().Value("user")
	if contextUser.(map[string]interface{})["access_level"] == 0 {
		c.JSON(401, gin.H{"Error": "You are not authorized to do this"})
		return
	}

	vote := &data.Vote{}
	err := json.NewDecoder(c.Request.Body).Decode(vote) //decode the request body into struct and failed if any error occur
	if err != nil {
		c.JSON(400, gin.H{"Error": "You must provide all params"})
		return
	}

	resp := vote.Create() //Create user
	c.JSON(200, gin.H{"success": resp})
}

// GetVote : GET a single vote
var GetVote = func(c *gin.Context) {
	uuidParam, err := uuid.FromString(c.Params.ByName("uuid"))

	if err != nil {
		c.JSON(400, gin.H{"Error": "You must provide the uuid in the params"})
		return
	}

	resp := data.GetVote(uuidParam)
	c.JSON(200, gin.H{"success": resp})
}

// UpdateVote : PUT/UPDATE a vote
var UpdateVote = func(c *gin.Context) {

	contextUser := c.Request.Context().Value("user")
	if contextUser.(map[string]interface{})["access_level"] == 0 {
		u.Respond(c.Writer, u.Message(false, "Error, you must have admin rights for this"))
		return
	}

	uuidParam, err := uuid.FromString(c.Params.ByName("uuid"))
	if err != nil {
		c.JSON(400, gin.H{"Error": "You must provide the uuid in the params"})
		return
	}

	vote := data.GetVote(uuidParam)
	if vote.Uuid.UUID != nil {
		newVote := &data.Vote{}
		err := json.NewDecoder(c.Request.Body).Decode(newVote)
		if err != nil {
			c.JSON(200, gin.H{"error": "You must provide all needed params"})
			return
		}

		newVote.ID = vote.ID
		resp := newVote.Update()
		c.JSON(200, gin.H{"success": resp})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

// DeleteVote : DELETE a vote
var DeleteVote = func(c *gin.Context) {

	contextUser := c.Request.Context().Value("user")
	if contextUser.(map[string]interface{})["access_level"] == 0 {
		u.Respond(c.Writer, u.Message(false, "Error, you must have admin rights for this"))
		return
	}

	uuidParam, err := uuid.FromString(c.Params.ByName("uuid"))
	if err != nil {
		c.JSON(400, gin.H{"Error": "You must provide the uuid in the params"})
		return
	}

	vote := &data.Vote{}
	vote.ID = data.GetVote(uuidParam).ID

	if vote.ID == 0 {
		vote.Delete()
		c.JSON(204, gin.H{"success": "Deleted successfuly"})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}
