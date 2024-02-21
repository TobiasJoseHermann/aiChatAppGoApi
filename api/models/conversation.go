package models

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	Name           string `json:"Name" gorm:"type:varchar(100)"`
	Email          string `json:"Email" gorm:"type:varchar(100)"`
	Messages    []Message `gorm:"foreignKey:ConversationID;references:ID"`
}