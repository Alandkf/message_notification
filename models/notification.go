package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	SenderUsername   string `gorm:"not null"`
	ReceiverUsername string `gorm:"not null"`
	Message          string `gorm:"not null"`
}
