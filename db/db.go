package db

import (
	"log"

	"myapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:(Aland&DB)@tcp(127.0.0.1:3306)/chat_service?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	db = db.Debug() // Keep this for development only
	log.Println("Connected to the database successfully")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Message{})

	return db
}
