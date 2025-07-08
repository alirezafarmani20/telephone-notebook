package server

import (
	"os"
	"telephone/internal/api"

	"github.com/gofiber/fiber/v2"
)

func CreateServer() {
	app := fiber.New()

	port := os.Getenv("PORT")
	// create test route
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")
	})

	// create routes for user
	app.Get("/users", api.GetUsers)
	app.Get("/users/:id", api.GetUserById)
	app.Post("/users", api.CreateUser)
	app.Put("/users/:id", api.UpdateUser)
	app.Delete("/users/:id", api.DeleteUser)

	app.Listen(port)
}
