package models

type Conversation struct {
	ConversationID int `json:"conversation_id"`
	Name string `json:"name"`
	Email string `json:"email"`
}