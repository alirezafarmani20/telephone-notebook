package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TimeLogger(ctx *fiber.Ctx) error {
	hour, minut, sectnd := time.Now().Clock()
	log.Println("time is %d:%d:%d", hour, minut, sectnd)
	return ctx.Next()
}