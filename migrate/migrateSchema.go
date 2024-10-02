package main

import (
	"log"
	"t-card/database"
	"t-card/models"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}
	database.ConnectDatabase()
}

func main() {
	database.DB.AutoMigrate(&models.User{})
}
