package database

import (
	"Auth/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {

	log.Println("⚡ main() has started")
	defer fmt.Println("Connection is Succesfull ✅ ")

	var err error
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Couldn't open the database: ", err)
	}
	DB.AutoMigrate(models.User{})
}
