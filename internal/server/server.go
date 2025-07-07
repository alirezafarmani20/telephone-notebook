package server

import "github.com/gofiber/fiber/v2"

func CreateServer() {
	app := fiber.New()

	app.Get("/health" , func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
}