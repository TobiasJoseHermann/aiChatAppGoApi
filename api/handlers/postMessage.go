package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/TobiasJoseHermann/goApi/api/models"
	"github.com/TobiasJoseHermann/goApi/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AiResponse struct {
	Content string `json:"content"`
}

func PostMessage(c *gin.Context) {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		var message models.Message
		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return err
		}
		fmt.Println("Message:")
		fmt.Println(message)

		if result := tx.Create(&message); result.Error != nil {
			return result.Error
		}

		// Response
		u, err := url.Parse("https://chat-flask-app-2sy27a6joa-uc.a.run.app/palm2")
		if err != nil {
			return err
		}

		q := u.Query()
		q.Set("user_input", message.Text)
		u.RawQuery = q.Encode()

		resp, err := http.Get(u.String())
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var aiResponse AiResponse
		err = json.Unmarshal(body, &aiResponse)
		if err != nil {
			return err
		}

		log.Println(aiResponse.Content)

		// Store Response
		aiMessage := models.Message{Text: aiResponse.Content, ConversationID: message.ConversationID, IsAiResponse: true}

		if result := tx.Create(&aiMessage); result.Error != nil {
			return result.Error
		}

		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Message sent")
}
