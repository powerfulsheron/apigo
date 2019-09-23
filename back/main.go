package main

import (
	"github.com/gorilla/mux"
	"apigo/back/controllers"
	"apigo/back/middleware"
	"os"
	"fmt"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.Use(middleware.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("api_port") //Get port from .env file
	if port == "" {
		fmt.Print("Can't find port from env, defaulting to 8080")
		port = "8080" //localhost if no port found
	}
	
	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}