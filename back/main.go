package main

import (
	"apigo/back/controllers"
	"apigo/back/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// --- AUTH/USERS ---
	router.POST("/users", func(c *gin.Context) {
		controllers.CreateUser(c.Writer, c.Request)
	})

	router.POST("/login", func(c *gin.Context) {
		controllers.Authenticate(c.Writer, c.Request)
	})

	router.PUT("/users/:uuid", controllers.UpdateUser)

	router.DELETE("/users/:uuid", controllers.DeleteUser)

	// --- VOTES ---
	router.GET("/votes", func(c *gin.Context) {
		controllers.GetVotes(c.Writer, c.Request)
	})

	router.POST("/votes", func(c *gin.Context) {
		controllers.CreateVote(c.Writer, c.Request)
	})

	router.GET("/votes/:uuid", controllers.GetVotes)

	router.PUT("/votes/:uuid", controllers.UpdateVote)

	router.DELETE("/votes/:uuid", controllers.DeleteVote)

	// Vote route
	router.Use(middleware.JwtAuthentication) // attach JWT auth middleware
	router.Run(os.Getenv("api_port"))
}
