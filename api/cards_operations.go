package api

import (
	"fiber-prometheus-demo/internal"
	"fiber-prometheus-demo/models"
)

// AddCardOperation adds a card into the database
func AddCardOperation(cardObj models.CardModel) error {
	result := internal.Database.Create(&cardObj)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// AddMultipleCardsOperation adds an array of cards into the database
func AddMultipleCardsOperation(cardList []models.CardModel) error {
	result := internal.Database.Create(&cardList)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ListCardsOperation retrieves all cards in the database
func ListCardsOperation() ([]models.CardModel, error) {
	var cards []models.CardModel
	result := internal.Database.Find(&cards)
	if result.Error != nil {
		return nil, result.Error
	}

	return cards, nil
}

// GetCardOperation retrieves a single card from the database
func GetCardOperation(id int) (*models.CardModel, error) {
	var cardObj models.CardModel

	result := internal.Database.First(&cardObj, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cardObj, nil
}

// UpdateCardOperation updates an existing card in the database
// I have to use GetCardOperation to see if the card exists and return a possible Not Found error
func UpdateCardOperation(id int, cardObj models.CardModel) error {
	_, err := GetCardOperation(id)
	if err != nil {
		return err
	}

	result := internal.Database.Where("id = ?", id).Updates(&cardObj)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeleteCardOperation deletes an existing card from the database
// I have to use GetCardOperation to see if the card exists and return a possible Not Found error
func DeleteCardOperation(id int) error {
	_, err := GetCardOperation(id)
	if err != nil {
		return err
	}

	result := internal.Database.Delete(&models.CardModel{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
