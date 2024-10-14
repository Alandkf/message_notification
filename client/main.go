package main

import (
    "bufio"
    "context"
    "fmt"
    "log"
    "os"
    "strings"
    "myapp/proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect to gRPC server: %v", err)
    }
    defer conn.Close()

    authClient := proto.NewAuthServiceClient(conn)
    messageClient := proto.NewMessageServiceClient(conn)
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("1. Login\n2. Register\n0. Server off")
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        if choice == "0" {
            fmt.Println("Server off.")
            break
        } else if choice == "1" {
            login(authClient, messageClient, reader)
        } else if choice == "2" {
            register(authClient, reader)
        } else {
            fmt.Println("Invalid option. Try again.")
        }
    }
}

func login(authClient proto.AuthServiceClient, messageClient proto.MessageServiceClient, reader *bufio.Reader) {
    fmt.Print("Enter username: ")
    username, _ := reader.ReadString('\n')
    username = strings.TrimSpace(username)

    fmt.Print("Enter password: ")
    password, _ := reader.ReadString('\n')
    password = strings.TrimSpace(password)

    resp, err := authClient.Login(context.Background(), &proto.UserRequest{Username: username, Password: password})
    if err != nil {
        fmt.Println("Login failed. Try again.")
        return
    }

    fmt.Println(resp.Message)
    userSession(messageClient, username, reader)
}

func register(authClient proto.AuthServiceClient, reader *bufio.Reader) {
    fmt.Print("Enter username: ")
    username, _ := reader.ReadString('\n')
    username = strings.TrimSpace(username)

    fmt.Print("Enter password: ")
    password, _ := reader.ReadString('\n')
    password = strings.TrimSpace(password)

    resp, err := authClient.Register(context.Background(), &proto.UserRequest{Username: username, Password: password})
    if err != nil {
        fmt.Println("Registration failed. Try again.")
        return
    }

    fmt.Println(resp.Message)
}

func userSession(messageClient proto.MessageServiceClient, username string, reader *bufio.Reader) {
    for {
        fmt.Println("1. Messages\n2. Logout\n0. Server off")
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        if choice == "0" {
            fmt.Println("Server off.")
            break
        } else if choice == "1" {
            showContacts(messageClient, username, reader)
        } else if choice == "2" {
            fmt.Println("Logged out.")
            break
        } else {
            fmt.Println("Invalid option. Try again.")
        }
    }
}

func showContacts(messageClient proto.MessageServiceClient, username string, reader *bufio.Reader) {
    contactsResp, err := messageClient.ListContacts(context.Background(), &proto.ContactListRequest{Username: username})
    if err != nil {
        fmt.Println("Failed to retrieve contacts.")
        return
    }

    fmt.Println("Contacts with unread messages:")
    for _, contact := range contactsResp.Contacts {
        fmt.Printf("%s (%d unread)\n", contact.Username, contact.UnreadMessages)
    }

    fmt.Print("Enter username to chat with: ")
    chatWith, _ := reader.ReadString('\n')
    chatWith = strings.TrimSpace(chatWith)

    showMessages(messageClient, username, chatWith, reader)
}

func showMessages(messageClient proto.MessageServiceClient, username, chatWith string, reader *bufio.Reader) {
    messagesResp, err := messageClient.ReadMessages(context.Background(), &proto.ReadMessagesRequest{
        Username: username,
        ChatWith: chatWith,
    })
    if err != nil {
        fmt.Println("Failed to retrieve messages.")
        return
    }

    fmt.Println("Messages with", chatWith)
    for _, message := range messagesResp.Messages {
        fmt.Printf("%s: %s (Seen: %v)\n", message.SenderUsername, message.Message, message.IsSeen)
    }

    fmt.Print("Send a message: ")
    msg, _ := reader.ReadString('\n')
    msg = strings.TrimSpace(msg)

    _, err = messageClient.SendMessage(context.Background(), &proto.SendMessageRequest{
        SenderUsername:  username,
        ReceiverUsername: chatWith,
        Message:         msg,
    })
    if err != nil {
        fmt.Println("Failed to send message.")
    }
}
