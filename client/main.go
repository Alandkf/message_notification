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
  conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    log.Fatalf("Failed to connect: %v", err)
  }
  defer conn.Close()

  client := proto.NewAuthServiceClient(conn)

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
    res, err := client.Register(context.Background(), &proto.UserRequest{
      Username: username,
      Password: password,
    })
    if err != nil {
      log.Fatalf("Failed to register: %v", err)
    }
    fmt.Println(res.Message)
  } else if choice == "2" {
    res, err := client.Login(context.Background(), &proto.UserRequest{
      Username: username,
      Password: password,
    })
    if err != nil {
      log.Fatalf("Failed to login: %v", err)
    }
    fmt.Println(res.Message)
  }
}
