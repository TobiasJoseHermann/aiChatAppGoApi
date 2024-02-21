package handlers

import (
	"net/http"

	"github.com/TobiasJoseHermann/goApi/api/models"
	"github.com/TobiasJoseHermann/goApi/db"

	"github.com/gin-gonic/gin"
)

func GetMessages(c *gin.Context) {
	conversationID := c.Param("ConversationID")

	var messages []models.Message
	if err := db.DB.Where("conversation_id = ?", conversationID).Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}
