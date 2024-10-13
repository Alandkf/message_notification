package db

import (
  "log"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "myapp/models" // Import the models package
)

func InitDB() *gorm.DB {
  dsn := "root:(Aland&DB)@tcp(127.0.0.1:3306)/login_register_service?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatalf("Failed to connect to the database: %v", err)
  }

  db = db.Debug()
  log.Println("Connected to the database successfully")

  // Auto-migrate the User model
  db.AutoMigrate(&models.User{})
  log.Println("Auto-migrated the User model")
  db.AutoMigrate(&models.Notification{})
  log.Println("Auto-migrated the Notification model")

  return db
}
