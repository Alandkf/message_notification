package models

import "gorm.io/gorm"

type Message struct {
    gorm.Model
    SenderID   uint   `gorm:"not null"`  // Foreign key from User
    ReceiverID uint   `gorm:"not null"`  // Foreign key from User
    Message    string `gorm:"not null"`
    IsSeen     bool   `gorm:"default:false"`

    // Associations with the User model
    Sender   User `gorm:"foreignKey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`  // Relationship with User
    Receiver User `gorm:"foreignKey:ReceiverID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relationship with User
}
