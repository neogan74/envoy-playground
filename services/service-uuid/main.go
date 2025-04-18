package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New()

	app.Get("/uuid", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"uuid": uuid.New().String(),
		})
	})

	app.Listen(":8080")
}
