package server

import (
    "context"
    "log"
    "myapp/models"
    "myapp/proto"
    "gorm.io/gorm"
)

type MessageServiceServer struct {
    DB *gorm.DB
    proto.UnimplementedMessageServiceServer
}

// SendMessage Implementation
func (s *MessageServiceServer) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.MessageResponse, error) {
    var sender, receiver models.User
    if err := s.DB.Where("username = ?", req.SenderUsername).First(&sender).Error; err != nil {
        return &proto.MessageResponse{Message: "Sender not found"}, err
    }
    if err := s.DB.Where("username = ?", req.ReceiverUsername).First(&receiver).Error; err != nil {
        return &proto.MessageResponse{Message: "Receiver not found"}, err
    }

    message := models.Message{
        SenderID:   sender.ID,
        ReceiverID: receiver.ID,
        Message:    req.Message,
    }

    if err := s.DB.Create(&message).Error; err != nil {
        return &proto.MessageResponse{Message: "Failed to send message"}, err
    }

    log.Printf("Message sent from %s to %s", req.SenderUsername, req.ReceiverUsername)
    return &proto.MessageResponse{Message: "Message sent successfully"}, nil
}

// ReadMessages Implementation
func (s *MessageServiceServer) ReadMessages(ctx context.Context, req *proto.ReadMessagesRequest) (*proto.ReadMessagesResponse, error) {
    var user models.User
    if err := s.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
        return nil, err
    }

    var chatWith models.User
    if err := s.DB.Where("username = ?", req.ChatWith).First(&chatWith).Error; err != nil {
        return nil, err
    }

    var messages []models.Message
    if err := s.DB.Where("receiver_id = ? AND sender_id = ?", user.ID, chatWith.ID).Or("receiver_id = ? AND sender_id = ?", chatWith.ID, user.ID).Find(&messages).Error; err != nil {
        return nil, err
    }

    var responseMessages []*proto.Notification
    for _, message := range messages {
        responseMessages = append(responseMessages, &proto.Notification{
            Id:             uint32(message.ID),
            SenderUsername: message.Sender.Username,
            Message:        message.Message,
            IsSeen:         message.IsSeen,
        })
    }

    // Mark all messages from chatWith as "seen"
    s.DB.Model(&models.Message{}).Where("sender_id = ? AND receiver_id = ? AND is_seen = ?", chatWith.ID, user.ID, false).Update("is_seen", true)

    return &proto.ReadMessagesResponse{Messages: responseMessages}, nil
}

// ListContacts - Retrieves contacts with unread messages
func (s *MessageServiceServer) ListContacts(ctx context.Context, req *proto.ContactListRequest) (*proto.ContactListResponse, error) {
    var user models.User
    if err := s.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
        return nil, err
    }

    var unreadContacts []struct {
        Username      string
        UnreadMessage int
    }

    // Query to get contacts with unread messages count
    if err := s.DB.Table("users").Select("users.username, COUNT(messages.id) as unread_message").
        Joins("JOIN messages ON messages.sender_id = users.id").
        Where("messages.receiver_id = ? AND messages.is_seen = false", user.ID).
        Group("users.username").Scan(&unreadContacts).Error; err != nil {
        return nil, err
    }

    var responseContacts []*proto.Contact
    for _, contact := range unreadContacts {
        responseContacts = append(responseContacts, &proto.Contact{
            Username:     contact.Username,
            UnreadMessages: int32(contact.UnreadMessage),
        })
    }

    return &proto.ContactListResponse{Contacts: responseContacts}, nil
}
