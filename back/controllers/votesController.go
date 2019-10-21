package controllers

import (
	"apigo/back/data"
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
	if contextUser.(map[string]string)["access_level"] == "0" {
		c.JSON(401, gin.H{"Error": "You are not authorized"})
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

	if resp.Uuid.UUID != nil {
		c.JSON(200, gin.H{"success": resp})
	} else {
		c.JSON(404, gin.H{"error": "Not found"})
	}
}

// UpdateVote : PUT/UPDATE a vote
var UpdateVote = func(c *gin.Context) {

	contextUser := c.Request.Context().Value("user")
	if contextUser.(map[string]string)["access_level"] == "0" {
		c.JSON(401, gin.H{"Error": "You are not authorized"})
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

		if newVote.Title != "" {
			vote.Title = newVote.Title
		}
		if newVote.Description != "" {
			vote.Description = newVote.Description
		}

		resp := vote.Update()
		c.JSON(200, gin.H{"success": resp})
	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

// DeleteVote : DELETE a vote
var DeleteVote = func(c *gin.Context) {

	contextUser := c.Request.Context().Value("user")
	if contextUser.(map[string]string)["access_level"] == "0" {
		c.JSON(401, gin.H{"Error": "You are not authorized"})
		return
	}

	uuidParam, err := uuid.FromString(c.Params.ByName("uuid"))
	if err != nil {
		c.JSON(400, gin.H{"Error": "You must provide the uuid in the params"})
		return
	}

	vote := data.GetVote(uuidParam)

	if vote.Uuid.UUID != nil {
		vote.Delete()
		c.JSON(204, gin.H{"success": "Deleted successfuly"})
	} else {
		c.JSON(404, gin.H{"error": "Vote not found"})
	}
}

// Vote : POST vote for a purpose
var Vote = func(c *gin.Context) {

	contextUser := c.Request.Context().Value("user")
	uuidUserStr := contextUser.(map[string]string)["uuid"]
	uuidUser, err := uuid.FromString(uuidUserStr)

	if err != nil {
		c.JSON(400, gin.H{"Error": "Invalid uuid"})
		return
	}

	uuidParam, err := uuid.FromString(c.Params.ByName("uuid"))
	if err != nil {
		c.JSON(400, gin.H{"Error": "You must provide the uuid in the params"})
		return
	}

	vote := data.GetVote(uuidParam)
	if vote.Uuid.UUID != nil {
		var resp = vote.Append(uuidUser)
		c.JSON(200, gin.H{"success": resp})

	} else {
		c.JSON(404, gin.H{"error": "User not found"})
	}
}
