package server

import (
  "context"
  "log"
  "fmt"
  "myapp/models"  // Import models package
  "myapp/proto"
  "gorm.io/gorm"
)

type AuthServiceServer struct {
  DB *gorm.DB
  proto.UnimplementedAuthServiceServer
}

func (s *AuthServiceServer) Register(ctx context.Context, req *proto.UserRequest) (*proto.AuthResponse, error) {
  user := models.User{Username: req.Username, Password: req.Password}

  if err := s.DB.Create(&user).Error; err != nil {
    return &proto.AuthResponse{Message: "Failed to register"}, err
  }

  log.Printf("New user registered: %s", req.Username)
  return &proto.AuthResponse{Message: fmt.Sprintf("User %s registered", req.Username)}, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *proto.UserRequest) (*proto.AuthResponse, error) {
  var user models.User
  if err := s.DB.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
    return &proto.AuthResponse{Message: "Invalid credentials"}, err
  }

  log.Printf("User logged in: %s", req.Username)
  return &proto.AuthResponse{Message: fmt.Sprintf("User %s logged in", req.Username)}, nil
}
