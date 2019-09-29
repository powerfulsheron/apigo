package data

import (
	"apigo/back/database"
	"apigo/back/models"
	u "apigo/back/utils"
)

// Vote : vote model
type Vote models.Vote

// Create an account
func (vote *Vote) Create() map[string]interface{} {
	database.GetDB().Create(vote)

	// Response
	response := u.Message(true, "Vote has been created")
	response["vote"] = vote
	return response
}

// Update an account
func (vote *Vote) Update() map[string]interface{} {
	database.GetDB().Save(vote)

	// Response
	response := u.Message(true, "Vote has been updated")
	response["vote"] = vote
	return response
}

// Delete an account
func (vote *Vote) Delete() map[string]interface{} {
	database.GetDB().Delete(vote)

	// Response
	response := u.Message(true, "Vote has been deleted")
	response["vote"] = vote
	return response
}

// GetVote from DB
func GetVote(uuid string) Vote {
	vote := Vote{}
	database.GetDB().Table("votes").Where("uuid = ?", uuid).First(&vote)
	return vote
}

// GetVotes get all votes from DB
func GetVotes() []Vote {
	votes := []Vote{}
	database.GetDB().Find(&votes)
	return votes
}
