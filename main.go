package main

import (
	"github.com/TobiasJoseHermann/goApi/api/handlers"
	"github.com/TobiasJoseHermann/goApi/api/middleware"
	"github.com/TobiasJoseHermann/goApi/api/models"
	"github.com/TobiasJoseHermann/goApi/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		panic(envErr.Error())
	}

	err := db.ConnectToDatabase()
	if err != nil {
		panic(err.Error())
	}

	app, err := middleware.InitFirebase()
	if err != nil {
		panic(err.Error())
	}

	db.DB.AutoMigrate(&models.Conversation{})
	db.DB.AutoMigrate(&models.Message{})

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "authtoken"}

	router.Use(cors.New(config))

	router.GET("/getConversations/:Email", middleware.CheckToken(app), handlers.GetConversations)
	router.GET("/getMessages/:ConversationID", middleware.CheckToken(app), handlers.GetMessages)

	router.POST("/postConversation", middleware.CheckToken(app), handlers.PostConversation)
	router.POST("/postMessage", middleware.CheckToken(app), handlers.PostMessage)

	router.DELETE("/deleteConversation/:ConversationID", middleware.CheckToken(app), handlers.DeleteConversation)

	router.Run("localhost:8080")
}
