package middleware

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// var app *firebase.App

func InitFirebase() (*firebase.App, error){
	opt := option.WithCredentialsFile("./api/middleware/credentials.json")

	// var err error
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		// log.Panic("error initializing app:", err)
		return app, err
	}
	log.Println("Firebase app initialized", app)
	return app, nil
}

func CheckToken(app *firebase.App) gin.HandlerFunc { return func(c *gin.Context) { 

		authToken := c.Request.Header.Get("authtoken")
		if authToken == "" {
			c.JSON(401, gin.H{"error": "Authorization token is required"})
			c.Abort()
		}


		ctx := context.Background()
		client, err := app.Auth(ctx)
		if err != nil {
			log.Printf("error getting Auth client: %v\n", err)
			c.Abort()
		}

		token, err := client.VerifyIDToken(ctx, authToken)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			c.Abort()
		}

		log.Printf("Verified ID token: %v\n", token)

		// Guarda el token decodificado en el contexto de Gin
		// c.Set("token", decodedToken)

		c.Next()
	}
}