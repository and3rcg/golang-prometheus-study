package models

import "gorm.io/gorm"

type CardModel struct {
	gorm.Model

	Name      string `json:"name"`
	Attribute string `json:"attribute"`
	Type      string `json:"type"`
	Level     int    `json:"level"`
	Attack    int    `json:"attack"`
	Defense   int    `json:"defense"`
}
