package database

import (
	"github.com/NicholasLiem/Email_Confirmation_Service_Go/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./database/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&schema.User{})
	if err != nil {
		return nil
	}

	return db
}
