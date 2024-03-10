package api

import (
	"fiber-prometheus-demo/models"
	"fiber-prometheus-demo/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// TODO: implement the validator: https://docs.gofiber.io/guide/validation

func AddCardHandler(c *fiber.Ctx) error {
	var request models.CardModel
	err := c.BodyParser(&request)
	if err != nil {
		return utils.BadRequestResponse(c, "AddCardHandler: Failed to parse request body", err)
	}

	err = AddCardOperation(request)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "AddCardHandler: Failed to add card", err)
	}

	return utils.CreatedResponse(c, "Card added successfully", fiber.Map{})
}

func AddMultipleCardsHandler(c *fiber.Ctx) error {
	var request []models.CardModel
	err := c.BodyParser(&request)
	if err != nil {
		return utils.BadRequestResponse(c, "AddMultipleCardsHandler: Failed to parse request body", err)
	}

	err = AddMultipleCardsOperation(request)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "AddMultipleCardsHandler: Failed to add cards", err)
	}

	return utils.CreatedResponse(c, "Cards added successfully", fiber.Map{})
}

func ListCardsHandler(c *fiber.Ctx) error {
	cardList, err := ListCardsOperation()
	if err != nil {
		return utils.InternalServerErrorResponse(c, "ListCardsHandler: Failed to list cards", err)
	}

	return utils.OkResponse(c, "Card list retrieved successfully", fiber.Map{
		"cards": cardList,
	})
}

func GetCardHandler(c *fiber.Ctx) error {
	cardID, err := strconv.Atoi(c.Params("card_id", "0"))
	if err != nil {
		return utils.BadRequestResponse(c, "GetCardHandler: Failed to parse card id", err)
	}

	cardObj, err := GetCardOperation(cardID)
	if err == gorm.ErrRecordNotFound {
		return utils.NotFoundResponse(c, "GetCardHandler: Card not found", err)
	} else if err != nil {
		return utils.InternalServerErrorResponse(c, "GetCardHandler: Failed to get card", err)
	}

	return utils.OkResponse(c, "Card retrieved successfully", fiber.Map{
		"card": cardObj,
	})
}

func UpdateCardHandler(c *fiber.Ctx) error {
	var request models.CardModel
	err := c.BodyParser(&request)
	if err != nil {
		return utils.BadRequestResponse(c, "UpdateCardHandler: Failed to parse request body", err)
	}

	cardID, err := strconv.Atoi(c.Params("card_id", "0"))
	if err != nil {
		return utils.InternalServerErrorResponse(c, "UpdateCardHandler: Failed to parse card id", err)
	}

	err = UpdateCardOperation(cardID, request)
	if err == gorm.ErrRecordNotFound {
		return utils.NotFoundResponse(c, "UpdateCardHandler: Card not found", err)
	} else if err != nil {
		return utils.InternalServerErrorResponse(c, "UpdateCardHandler: Failed to update card", err)
	}

	return utils.OkResponse(c, "Card updated successfully", fiber.Map{})
}

func DeleteCardHandler(c *fiber.Ctx) error {
	cardID, err := strconv.Atoi(c.Params("card_id", "0"))
	if err != nil {
		return utils.InternalServerErrorResponse(c, "UpdateCardHandler: Failed to parse card id", err)
	}

	err = DeleteCardOperation(cardID)
	if err == gorm.ErrRecordNotFound {
		return utils.NotFoundResponse(c, "DeleteCardHandler: Card not found", err)
	} else if err != nil {
		return utils.InternalServerErrorResponse(c, "DeleteCardHandler: Failed to delete card", err)
	}

	return utils.OkResponse(c, "Card deleted successfully", fiber.Map{})
}
