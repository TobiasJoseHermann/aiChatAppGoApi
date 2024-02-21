package db

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() error {
	server := os.Getenv("DB_SERVER")
	userId := os.Getenv("DB_USER_ID")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	connStringGorm := fmt.Sprintf("sqlserver://%s:%s@%s:1433?database=%s&app name=aiChatAppGoApi&encrypt=true&trustServerCertificate=true", userId, password, server, database)

	var err error

	DB, err = gorm.Open(sqlserver.Open(connStringGorm), &gorm.Config{})
	if err != nil {
		DB = nil
		return err
	}

	println("Connected!", DB)
	return nil
}
