package utils

import "github.com/gofiber/fiber/v2"

func newResponse(c *fiber.Ctx, statusCode int, msg string, err error, data fiber.Map) error {
	if err != nil {
		return c.Status(statusCode).JSON(fiber.Map{
			"message": msg,
			"error":   err.Error(),
			"data":    data,
		})
	} else {
		return c.Status(statusCode).JSON(fiber.Map{
			"message": msg,
			"error":   err,
			"data":    data,
		})
	}
}

func OkResponse(c *fiber.Ctx, msg string, data fiber.Map) error {
	return newResponse(c, fiber.StatusOK, msg, nil, data)
}

func CreatedResponse(c *fiber.Ctx, msg string, data fiber.Map) error {
	return newResponse(c, fiber.StatusCreated, msg, nil, data)
}

func BadRequestResponse(c *fiber.Ctx, msg string, err error) error {
	return newResponse(c, fiber.StatusBadRequest, msg, err, nil)
}

func UnauthorizedResponse(c *fiber.Ctx, msg string, err error) error {
	return newResponse(c, fiber.StatusUnauthorized, msg, err, nil)
}

func ForbiddenResponse(c *fiber.Ctx, msg string, err error) error {
	return newResponse(c, fiber.StatusForbidden, msg, err, nil)
}

func NotFoundResponse(c *fiber.Ctx, msg string, err error) error {
	return newResponse(c, fiber.StatusNotFound, msg, err, nil)
}

func InternalServerErrorResponse(c *fiber.Ctx, msg string, err error) error {
	return newResponse(c, fiber.StatusInternalServerError, msg, err, nil)
}
