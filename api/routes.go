package api

import "github.com/gofiber/fiber/v2"

func APIRoutes(app *fiber.App) {
	groupAPI := app.Group("/api/v1")

	groupAPI.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
