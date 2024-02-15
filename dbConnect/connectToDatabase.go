package dbConnect

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func ConnectToDatabase() (error) {
    server := os.Getenv("DB_SERVER")
    userId := os.Getenv("DB_USER_ID")
    password := os.Getenv("DB_PASSWORD")
    database := os.Getenv("DB_DATABASE")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", server, userId, password, database)

	var err error
    DB, err = sql.Open("sqlserver", connString)
    if err != nil {
		DB = nil
        return err
    }
    err = DB.Ping()
    if err != nil {
		DB = nil
        return err
    }
	println("Connected!", DB)
    return nil
}