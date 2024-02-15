package handlers

import (
	"database/sql"
	"net/http"

	"github.com/TobiasJoseHermann/goApi/api/models"
	"github.com/TobiasJoseHermann/goApi/dbConnect"

	"github.com/gin-gonic/gin"
)




func GetConversations(c *gin.Context) {

	type requestBody struct {
		Email string `json:"email"`
	}

    var reqBody requestBody
    if err := c.ShouldBindJSON(&reqBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    rows, err := dbConnect.DB.Query("SELECT * FROM Conversation WHERE email = @email", sql.Named("email", reqBody.Email))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var conversations []models.Conversation
    for rows.Next() {
        var conv models.Conversation
        if err := rows.Scan(&conv.ConversationID, &conv.Name, &conv.Email); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        conversations = append(conversations, conv)
    }

    c.JSON(http.StatusOK, conversations)
}
