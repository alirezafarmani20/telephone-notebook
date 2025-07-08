package main

import (
	"log"
	"telephone/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error .env file is not loading")
	}
	server.CreateServer()
}
