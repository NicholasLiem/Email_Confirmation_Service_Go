package main

import (
	"github.com/NicholasLiem/Email_Confirmation_Service_Go/database"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.SetupDB()
}
