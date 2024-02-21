package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ConversationID int    `json:"ConversationID"`
	Text           string `json:"Text" gorm:"type:varchar(1000)"`
	IsAiResponse   bool   `json:"IsAiResponse" gorm:"type:bit"`
}
