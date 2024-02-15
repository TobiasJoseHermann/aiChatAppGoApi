package handlers

import (
	"database/sql"
	"net/http"

	"github.com/TobiasJoseHermann/goApi/api/models"
	"github.com/TobiasJoseHermann/goApi/dbConnect"
	"github.com/gin-gonic/gin"
)


func ChangeConversation(c *gin.Context) {

	type requestBody struct {
		ConversationID int `json:"id"`
	}

    var reqBody requestBody
    if err := c.ShouldBindJSON(&reqBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    rows, err := dbConnect.DB.Query("SELECT * FROM UserMessage WHERE conversation_id = @id", sql.Named("id", reqBody.ConversationID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var result [] models.UserMessage
    for rows.Next() {
        var msg models.UserMessage
        if err := rows.Scan(&msg.ConversationID, &msg.Text, &msg.ConversationID, &msg.Is_ai_response); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        result = append(result, msg)
    }

    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, result)
}