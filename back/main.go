package main

import (
	"github.com/gin-gonic/gin"
	"apigo/back/controllers"
	"apigo/back/middleware"
	"os"
	"fmt"
)

func main() {

	router := gin.Default()
	// Login route
	router.POST("/login", controllers.Authenticate)
	// New user route
	router.POST("/users", controllers.CreateAccount)
	
	router.Use(middleware.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("api_port") //Get port from .env file
	if port == "" {
		fmt.Print("Can't find port from env, defaulting to 8080")
		port = "8080" //localhost if no port found
	}

	router.Run(port)
	
}