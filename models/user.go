package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
    // This is optional, but you can add a slice of Messages to represent the relationship
    SentMessages     []Message `gorm:"foreignKey:SenderID"`
    ReceivedMessages []Message `gorm:"foreignKey:ReceiverID"`
}
