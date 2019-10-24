package router

import (
	"apigo/back/controllers"
	"apigo/back/middleware"

	"github.com/gin-gonic/gin"
)

// VoteRouter define a test router
func VoteRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.IPFirewall())      // attach IPBlock middleware
	router.Use(middleware.JwtAuthentication) // attach JWT auth middleware

	// --- AUTH ---
	router.POST("/users", func(c *gin.Context) {
		controllers.CreateUser(c.Writer, c.Request)
	})
	router.POST("/login", func(c *gin.Context) {
		controllers.Authenticate(c.Writer, c.Request)
	})

	// --- USERS ---
	router.PUT("/users/:uuid", controllers.UpdateUser)
	router.DELETE("/users/:uuid", controllers.DeleteUser)

	// --- VOTES ---
	router.GET("/votes", controllers.GetVotes)
	router.POST("/votes", controllers.CreateVote)
	router.GET("/votes/:uuid", controllers.GetVote)
	router.PUT("/votes/:uuid", controllers.UpdateVote)
	router.DELETE("/votes/:uuid", controllers.DeleteVote)
	router.POST("/votes/:uuid", controllers.Vote)

	return router
}
