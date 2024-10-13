package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"myapp/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	authClient := proto.NewAuthServiceClient(conn)
	messageClient := proto.NewMessageServiceClient(conn)

	var username, password, choice string
	fmt.Println("Choose: \n	1 to register new account \n	2 to login your account")
	fmt.Scanln(&choice)

	if choice != "1" && choice != "2" {
		fmt.Println("Invalid choice")
		os.Exit(1)
	}
	fmt.Println("Enter username:")
	fmt.Scanln(&username)
	fmt.Println("Enter password:")
	fmt.Scanln(&password)

	if choice == "1" {
		res, err := authClient.Register(context.Background(), &proto.UserRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Fatalf("Failed to register: %v", err)
		}
		fmt.Println(res.Message)
	} else if choice == "2" {
		res, err := authClient.Login(context.Background(), &proto.UserRequest{
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Fatalf("Failed to login: %v", err)
		}
		fmt.Println(res.Message)

		// After login, ask user for next action
		for {
			fmt.Println("Choose: \n	1 to send message \n	2 to read messages \n	0 to logout")
			fmt.Scanln(&choice)

			switch choice {
			case "1":
				var receiverUsername, message string
				fmt.Println("Enter the username of the person you want to send a message to:")
				fmt.Scanln(&receiverUsername)
				fmt.Println("Enter your message:")
				fmt.Scanln(&message)

				msgRes, err := messageClient.SendMessage(context.Background(), &proto.SendMessageRequest{
					SenderUsername:   username,
					ReceiverUsername: receiverUsername,
					Message:          message,
				})
				if err != nil {
					log.Fatalf("Failed to send message: %v", err)
				}
				fmt.Println(msgRes.Message)

			case "2":
				res, err := messageClient.ReadMessages(context.Background(), &proto.ReadMessagesRequest{
					Username: username,
				})
				if err != nil {
					log.Fatalf("Failed to read messages: %v", err)
				}
				for _, notification := range res.Notifications {
					fmt.Printf("From: %s, Message: %s\n", notification.SenderUsername, notification.Message)
				}

			case "0":
				fmt.Println("Logging out...")
				return

			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}
	}
}
