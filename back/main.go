package main

import (
	"apigo/back/router"
	"os"
)

func main() {
	router := router.VoteRouter()
	router.Run(os.Getenv("api_port"))
}
