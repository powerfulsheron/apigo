package main

import (
	"github.com/gin-gonic/gin"
	"apigo/back/controllers"
	"apigo/back/middleware"
	"os"
)

func main() {

	router := gin.Default()
	router.POST("/login", func(c *gin.Context){
		controllers.CreateAccount(c.Writer, c.Request)
	})
	router.Use(middleware.JwtAuthentication) //attach JWT auth middleware
	router.Run(os.Getenv("api_port"))
}