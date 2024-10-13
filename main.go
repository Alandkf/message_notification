package main

import (
	"fmt"
	"log"
	"net"

	"myapp/db"
	"myapp/proto"
	"myapp/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	dbConn := db.InitDB()
	fmt.Println("DB Connection: ", dbConn)
	grpcServer := grpc.NewServer()
	authService := &server.AuthServiceServer{DB: dbConn}
	messageService := &server.MessageServiceServer{DB: dbConn}

	proto.RegisterAuthServiceServer(grpcServer, authService)
	proto.RegisterMessageServiceServer(grpcServer, messageService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Println("gRPC server is running on port 50051")
}
