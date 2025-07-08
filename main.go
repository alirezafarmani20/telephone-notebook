package main

import (
	"log"
	"telephone/internal/database"
	"telephone/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error .env file is not loading")
	}
	database.ConnectToDatabase()
	server.CreateServer()
}
