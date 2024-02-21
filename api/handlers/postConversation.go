package handlers

import (
	"net/http"

	"github.com/TobiasJoseHermann/goApi/api/models"
	"github.com/TobiasJoseHermann/goApi/db"
	"github.com/gin-gonic/gin"
)

func PostConversation(c *gin.Context) {
	var conversation models.Conversation
	if err := c.ShouldBindJSON(&conversation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&conversation)
	c.JSON(http.StatusOK, conversation)
}
