package controllers

import (
	"apigo/back/data"
	u "apigo/back/utils"
	"encoding/json"
	"net/http"
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
	contextUser := r.Context().Value("user")
	if user.AccessLevel != 0 && contextUser.(map[string]interface{})["access_level"]!=0 {
		resp := user.Create() //Create Admin
		u.Respond(w, resp)
		return
	} else if user.AccessLevel != 0 &&  contextUser.(map[string]interface{})["access_level"]==0 {
		u.Respond(w, u.Message(false, "Error, you must have admin rights for this"))
		return
	} else {
		resp := user.Create() //Create user
		u.Respond(w, resp)
		return
	}
}

var UpdateUser = func(c *gin.Context) {
	contextUser := c.Request.Context().Value("user")
	if contextUser.(map[string]interface{})["access_level"]==0 {
		u.Respond(c.Writer, u.Message(false, "Error, you must have admin rights for this"))
		return
	}
	uuidParam, err := uuid.FromString(c.Params.ByName("uuid"))
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "Can't find specified uuid"))
		return
	}
	user := data.GetUser(uuidParam)
	if user.Uuid.UUID != nil {
		newUser := &data.User{}
		err := json.NewDecoder(c.Request.Body).Decode(newUser) //decode the request body into struct and failed if any error occur
		if err != nil {
			u.Respond(c.Writer, u.Message(false, "Errors in user parameters"))
			return
		}
		newUser.ID = user.ID
		// Validate user
		resp := newUser.Update()
		// Display modified data in JSON message "success"
		c.JSON(200, gin.H{"success": resp})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

var DeleteUser = func(c *gin.Context) {
	contextUser := c.Request.Context().Value("user")
	if contextUser.(map[string]interface{})["access_level"]==0 {
		u.Respond(c.Writer, u.Message(false, "Error, you must have admin rights for this"))
		return
	}
	uuidParam, err := uuid.FromString(c.Params.ByName("uuid"))
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "Can't find specified uuid"))
		return
	}
	user := &data.User{}
	user.ID = data.GetUser(uuidParam).ID
	if user.ID == 0 {
		// Delete user
		resp := user.Delete()
		// Display modified data in JSON message "success"
		c.JSON(200, gin.H{"success": resp})
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