package controllers

import (
	"github.com/gin-gonic/gin"
	"apigo/back/data"
	"encoding/json"
)


// CreateAccount : new user controller
var CreateAccount = func(c *gin.Context) {

	account := &data.Account{}
	err := json.NewDecoder(c.Request.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		c.JSON(500, "Error")
		return
	}

	resp := account.Create() //Create account
	c.JSON(200, resp)
}
// Authenticate : login user controller
var Authenticate = func(c *gin.Context) {

	account := &data.Account{}
	err := json.NewDecoder(c.Request.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		c.JSON(500, "Error")
		return
	}
	
	resp := data.Login(account.Email, account.Password)
	c.JSON(200, resp)
}