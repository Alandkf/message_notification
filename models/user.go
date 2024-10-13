// models/user.go
package models

import "gorm.io/gorm"

type User struct {
  gorm.Model
  Username string `gorm:"unique"`
  Password string
}

type Message struct {
  gorm.Model
  SenderID   uint
  ReceiverID uint
  Message    string
}