package api

import "github.com/gofiber/fiber/v2"

func StartAPIRoutes(app *fiber.App) {
	groupAPI := app.Group("/api/v1")

	groupAPI.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	groupAPI.Post("/cards", AddCardHandler)
	groupAPI.Post("/cards/multiple", AddMultipleCardsHandler)
	groupAPI.Get("/cards", ListCardsHandler)
	groupAPI.Get("/cards/:card_id", GetCardHandler)
	groupAPI.Put("/cards/:card_id", UpdateCardHandler)
	groupAPI.Delete("/cards/:card_id", DeleteCardHandler)
}
