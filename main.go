package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const version = "0.1.0"

func main() {
	fiberConfig := fiber.Config{
		AppName: fmt.Sprintf("Fiber DevOps showcase v%s", version),
	}
	app := fiber.New(fiberConfig)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":4000")
}
