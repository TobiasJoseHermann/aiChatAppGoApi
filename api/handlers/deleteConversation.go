package handlers

import (
	"database/sql"
	"net/http"

	"github.com/TobiasJoseHermann/goApi/dbConnect"

	"github.com/gin-gonic/gin"
)

func DeleteConversation(c *gin.Context) {
	type requestBody struct {
		ConversationID int `json:"id"`
	}

	var reqBody requestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := dbConnect.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = tx.Exec("DELETE FROM UserMessage WHERE conversation_id = @conversation_id", sql.Named("conversation_id", reqBody.ConversationID))
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = tx.Exec("DELETE FROM Conversation WHERE conversation_id = @conversation_id", sql.Named("conversation_id", reqBody.ConversationID))
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Conversation and associated messages deleted")
}