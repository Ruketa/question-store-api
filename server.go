package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/question-store-api/db"
	"github.com/question-store-api/router"
)

const defaultPort = "8000"

func initServer() *gin.Engine {
	server := gin.Default()
	router.Setup(server)
	return server
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db.InitDB()
	defer db.CloseDB()

	server := initServer()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(server.Run(":8000"))
}
