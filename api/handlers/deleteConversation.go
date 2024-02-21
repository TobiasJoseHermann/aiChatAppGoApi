package handlers

import (
	"net/http"

	"github.com/TobiasJoseHermann/goApi/api/models"
	"github.com/TobiasJoseHermann/goApi/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteConversation(c *gin.Context) {
	conversationID := c.Param("ConversationID")

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// Delete messages
		if err := tx.Where("conversation_id = ?", conversationID).Delete(&models.Message{}).Error; err != nil {
			return err
		}

		// Delete conversation
		if err := tx.Where("id = ?", conversationID).Delete(&models.Conversation{}).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Conversation and associated messages deleted")
}
