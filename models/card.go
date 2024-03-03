package models

type CardModel struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Condition string `json:"condition"`
	Amount    int    `json:"amount"`
	PackName  string `json:"pack_name"`
}
