package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/TobiasJoseHermann/goApi/dbConnect"
	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {

	// Message

	type requestBody struct {
		Message string `json:"message"`
		ConversationID int `json:"conversation_id"`
	}

	var reqBody requestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRows, err := dbConnect.DB.Query("INSERT INTO UserMessage (text,conversation_id, is_ai_response) VALUES (@text, @conversation_id, @is_ai_response)", sql.Named("text", reqBody.Message),sql.Named("conversation_id", reqBody.ConversationID), sql.Named("is_ai_response", false))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(userRows)

	// Response

	u, err := url.Parse("https://chat-flask-app-2sy27a6joa-uc.a.run.app/palm2")
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("user_input", reqBody.Message)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	type Response struct {
		Content string `json:"content"`
	}

	var respContent Response
	err = json.Unmarshal(body, &respContent)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(respContent.Content)

	// Store Response

	aiRows, err := dbConnect.DB.Query("INSERT INTO UserMessage (text,conversation_id, is_ai_response) VALUES (@text, @conversation_id, @is_ai_response)", sql.Named("text", respContent.Content),sql.Named("conversation_id", reqBody.ConversationID), sql.Named("is_ai_response", true))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(aiRows)

	c.JSON(http.StatusOK, "Message sent")
}