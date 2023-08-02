package model

import "gorm.io/gorm"

// Player represents the model for the "players" table in the database.
type Player struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:15"`
	Country string `json:"country" gorm:"size:2"`
	Score   int    `json:"score"`
}
