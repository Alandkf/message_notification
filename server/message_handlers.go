package server

import (
	"context"
	// "fmt"
	"log"
	"myapp/models"
	"myapp/proto"
	"gorm.io/gorm"
)

type MessageServiceServer struct {
	DB *gorm.DB
	proto.UnimplementedMessageServiceServer
}

func (s *MessageServiceServer) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*proto.MessageResponse, error) {
	notification := models.Notification{
		SenderUsername:   req.SenderUsername,
		ReceiverUsername: req.ReceiverUsername,
		Message:          req.Message,
	}

	if err := s.DB.Create(&notification).Error; err != nil {
		return &proto.MessageResponse{Message: "Failed to send message"}, err
	}

	log.Printf("Message sent from %s to %s", req.SenderUsername, req.ReceiverUsername)
	return &proto.MessageResponse{Message: "Message sent successfully"}, nil
}

func (s *MessageServiceServer) ReadMessages(ctx context.Context, req *proto.ReadMessagesRequest) (*proto.ReadMessagesResponse, error) {
	var notifications []models.Notification
	if err := s.DB.Where("receiver_username = ?", req.Username).Find(&notifications).Error; err != nil {
		return nil, err
	}

	var responseNotifications []*proto.Notification
	for _, notification := range notifications {
		responseNotifications = append(responseNotifications, &proto.Notification{
			Id:              uint32(notification.ID),
			SenderUsername:  notification.SenderUsername,
			Message:         notification.Message,
		})
	}

	return &proto.ReadMessagesResponse{Notifications: responseNotifications}, nil
}
