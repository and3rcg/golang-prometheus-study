package api

import (
	"fiber-prometheus-demo/internal"
	"fiber-prometheus-demo/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RandomSleepHandler(c *fiber.Ctx) error {
	randomDuration := rand.Intn(10)

	time.Sleep(time.Duration(randomDuration) * time.Second)

	return utils.OkResponse(c, fmt.Sprintf("Waited for %d seconds", randomDuration), fiber.Map{})
}

func PostEchoHandler(c *fiber.Ctx) error {
	type requestData struct {
		Message string `json:"message"`
	}
	var request requestData

	err := c.BodyParser(&request)
	if err != nil {
		return utils.BadRequestResponse(c, "PostEchoHandler: Failed to parse request body", err)
	}

	return utils.OkResponse(c, "Echo message", fiber.Map{
		"message": request.Message,
	})
}

func TriggerOkResponseHandler(c *fiber.Ctx) error {
	internal.Prometheus.APITriggersMetric.WithLabelValues("trigger_ok").Inc()
	return utils.OkResponse(c, "Triggered OK response", fiber.Map{})
}

func TriggerCreatedResponseHandler(c *fiber.Ctx) error {
	internal.Prometheus.APITriggersMetric.WithLabelValues("trigger_created").Inc()
	return utils.CreatedResponse(c, "Triggered Created response", fiber.Map{})
}

func TriggerBadRequestResponseHandler(c *fiber.Ctx) error {
	internal.Prometheus.APITriggersMetric.WithLabelValues("trigger_bad_request").Inc()
	return utils.BadRequestResponse(c, "Triggered Bad Request response", nil)
}

func TriggerUnauthorizedResponseHandler(c *fiber.Ctx) error {
	internal.Prometheus.APITriggersMetric.WithLabelValues("trigger_unauthorized").Inc()
	return utils.UnauthorizedResponse(c, "Triggered Unauthorized response", nil)
}

func TriggerForbiddenResponseHandler(c *fiber.Ctx) error {
	internal.Prometheus.APITriggersMetric.WithLabelValues("trigger_forbidden").Inc()
	return utils.ForbiddenResponse(c, "Triggered Forbidden response", nil)
}

func TriggerNotFoundResponseHandler(c *fiber.Ctx) error {
	internal.Prometheus.APITriggersMetric.WithLabelValues("trigger_not_found").Inc()
	return utils.NotFoundResponse(c, "Triggered Not Found response", nil)
}

func TriggerInternalServerErrorResponseHandler(c *fiber.Ctx) error {
	internal.Prometheus.APITriggersMetric.WithLabelValues("trigger_internal_server_error").Inc()
	return utils.InternalServerErrorResponse(c, "Triggered Internal Server Error response", nil)
}
