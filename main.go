package main

import (
	"github.com/TobiasJoseHermann/goApi/api/handlers"
	"github.com/TobiasJoseHermann/goApi/api/middleware"
	"github.com/TobiasJoseHermann/goApi/dbConnect"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {

	envErr := godotenv.Load(".env")
	if envErr != nil{
		panic(envErr.Error())
	}

	err := dbConnect.ConnectToDatabase()
	if err != nil {
		panic(err.Error())
	}
	defer dbConnect.DB.Close()

	app, err := middleware.InitFirebase()
	if err != nil {
		panic(err.Error())
	}



	router := gin.Default()

	config := cors.DefaultConfig()
    config.AllowAllOrigins = true  // Permite todas las origenes
    config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}  // Permite todos los métodos HTTP
    config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "authtoken"}  // Permite todos los encabezados
	

    router.Use(cors.New(config))
	
	router.POST("/addConversationAPI", middleware.CheckToken(app),handlers.AddConversation)
	router.POST("/conversationsAPI", middleware.CheckToken(app),handlers.GetConversations)
	router.POST("/changeConversationAPI", middleware.CheckToken(app),handlers.ChangeConversation)
	router.POST("/sendMessageAPI", middleware.CheckToken(app),handlers.SendMessage)
	router.DELETE("/deleteConversationAPI", middleware.CheckToken(app),handlers.DeleteConversation)
	

	
	router.Run("localhost:8080")
}
