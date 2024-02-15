package handlers

import (
	"database/sql"
	"net/http"

	"github.com/TobiasJoseHermann/goApi/dbConnect"

	"github.com/gin-gonic/gin"
)


func AddConversation(c *gin.Context) {

	type requestBody struct {
		Name string `json:"name"`
		Email string `json:"email"`
	}

    var reqBody requestBody
    if err := c.ShouldBindJSON(&reqBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    rows, err := dbConnect.DB.Query("INSERT INTO Conversation (name,email) VALUES (@name, @email)", sql.Named("name", reqBody.Name),sql.Named("email", reqBody.Email))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	println(rows)
    c.JSON(http.StatusOK, "Conversation added")
}