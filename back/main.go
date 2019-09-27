package main

import (
	"apigo/back/controllers"
	"apigo/back/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// --- AUTH ---
	router.POST("/new", func(c *gin.Context) {
		controllers.CreateUser(c.Writer, c.Request)
	})

	router.POST("/login", func(c *gin.Context) {
		controllers.Authenticate(c.Writer, c.Request)
	})

	// --- VOTES ---
	router.GET("/votes", func(c *gin.Context) {
		controllers.GetVotes(c.Writer, c.Request)
	})

	router.POST("/votes", func(c *gin.Context) {
		controllers.CreateVote(c.Writer, c.Request)
	})

	// Vote route
	router.Use(middleware.JwtAuthentication) // attach JWT auth middleware
	router.Run(os.Getenv("api_port"))
}
