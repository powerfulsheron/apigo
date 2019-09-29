package main

import (
	"apigo/back/controllers"
	"apigo/back/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middleware.IPFirewall()) // attach IPBlock middleware
	router.Use(middleware.JwtAuthentication) // attach JWT auth middleware

	// --- AUTH ---
	router.POST("/users", func(c *gin.Context) {
		controllers.CreateUser(c.Writer, c.Request)
	})

	router.POST("/login", func(c *gin.Context) {
		controllers.Authenticate(c.Writer, c.Request)
	})

	router.PUT("/users/:uuid", controllers.UpdateUser)

	router.DELETE("/users/:uuid", controllers.DeleteUser)

	// --- VOTES ---
	router.GET("/votes", controllers.GetVotes)

	router.POST("/votes", controllers.CreateVote)

	router.GET("/votes/:uuid", controllers.GetVote)

	router.PUT("/votes/:uuid", controllers.UpdateVote)

	router.DELETE("/votes/:uuid", controllers.DeleteVote)

	router.POST("/votes/:uuid", controllers.Vote)

	// Vote route
	router.Use(middleware.JwtAuthentication) // attach JWT auth middleware
	router.Run(os.Getenv("api_port"))
}
