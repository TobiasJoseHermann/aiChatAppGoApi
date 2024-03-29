package handlers

import (
	"net/http"

	"github.com/TobiasJoseHermann/goApi/api/models"
	"github.com/TobiasJoseHermann/goApi/db"

	"github.com/gin-gonic/gin"
)

func GetConversations(c *gin.Context) {
	email := c.Param("Email")

	var conversations []models.Conversation
	if err := db.DB.Where("email = ?", email).Find(&conversations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, conversations)
}
