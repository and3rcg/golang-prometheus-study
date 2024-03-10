package api

import "github.com/gofiber/fiber/v2"

func StartAPIRoutes(app *fiber.App) {
	cardsAPI := app.Group("/api/v1")

	cardsAPI.Post("/cards", AddCardHandler)
	cardsAPI.Post("/cards/multiple", AddMultipleCardsHandler)
	cardsAPI.Get("/cards", ListCardsHandler)
	cardsAPI.Get("/cards/:card_id", GetCardHandler)
	cardsAPI.Put("/cards/:card_id", UpdateCardHandler)
	cardsAPI.Delete("/cards/:card_id", DeleteCardHandler)

	metricsTestingAPI := app.Group("/api/v1/metrics_testing")

	metricsTestingAPI.Get("/random_sleep", RandomSleepHandler)
	metricsTestingAPI.Post("/echo", PostEchoHandler)
	metricsTestingAPI.Get("/trigger/ok", TriggerOkResponseHandler)
	metricsTestingAPI.Get("trigger/created", TriggerCreatedResponseHandler)
	metricsTestingAPI.Get("trigger/bad_request", TriggerBadRequestResponseHandler)
	metricsTestingAPI.Get("trigger/unauthorized", TriggerUnauthorizedResponseHandler)
	metricsTestingAPI.Get("trigger/forbidden", TriggerForbiddenResponseHandler)
	metricsTestingAPI.Get("trigger/not_found", TriggerNotFoundResponseHandler)
	metricsTestingAPI.Get("trigger/internal_server_error", TriggerInternalServerErrorResponseHandler)
}
