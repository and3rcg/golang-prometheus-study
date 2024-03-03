package api

import "github.com/gofiber/fiber/v2"

func AddCardHandler(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "card added",
	})
}
