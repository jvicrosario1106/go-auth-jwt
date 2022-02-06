package database

import (
	"log"

	"github.com/jvicrosario1106/jwt-golang/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to Connect to database")
	}

	log.Print("Successfully Connect to Database")

	DB = db
	db.AutoMigrate(&models.User{})

}
