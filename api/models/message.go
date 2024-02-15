package models

type UserMessage struct {
	UserMessageID int `json:"user_message_id"`
	Text string `json:"text"`
	ConversationID int `json:"conversation_id"`
	Is_ai_response bool `json:"is_ai_response"`
}
