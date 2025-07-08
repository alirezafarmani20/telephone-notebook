package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func CreateServer() {
	app := fiber.New()

	port := os.Getenv("PORT")
	// create test route
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")
	})
	
	app.Listen(port)
}