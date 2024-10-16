Messaging Application - gRPC Service
Overview
This project is a basic messaging application built using Go, gRPC, and MySQL. Users can register, log in, send messages, and interact with contacts. The application implements client-server communication through gRPC, providing a seamless experience for managing authentication and messaging services.

Features
User Registration & Login:

Users can register with a unique username and password.
Upon successful registration, users are directed to log in.
Login credentials are verified by the server to authenticate users.
Messaging Service:

Once logged in, users can access their chat session.
The application displays a list of contacts with unread messages.
Users can select a contact to view their chat history and send messages in real time.
Session Management:

Logged-in users can choose between:
Viewing and sending messages.
Logging out.
Shutting down the server.
The server handles multiple user sessions and allows for clean disconnection when needed.
Project Flow
The application follows this basic flow:

Start the Application:

The client connects to the gRPC server.
Main Menu:

The menu offers three options: Login, Register, and Server Shutdown.
The server can be shut down from this menu by selecting option 0.
User Registration:

Users choosing to register are prompted to enter their username and password.
Registration success is indicated by a message, and the user is prompted to log in.
User Login:

After entering valid credentials, users are taken to their session where they can:
View unread messages.
Logout.
Shut down the server.
Messaging Interface:

Users can view their contacts and select a conversation.
Once a conversation is selected, the user can view the chat history and send new messages.
Logout/Shutdown:

Users can log out and return to the main menu.
The server can be shut down cleanly through the menu.

Folder Structure
├── proto/               # Contains .proto file(s) for defining gRPC services.
├── client/              # Contains the client logic for connecting to the gRPC server.
│   ├── main.go          # Entry point for the client-side application.
├── server/              # Contains the server logic for managing gRPC services.
│   ├── main.go          # Entry point for the server-side application.
├── models/              # Contains database models and connection setup.
│   ├── user.go          # User model.
│   ├── message.go       # Message model.
├── README.md            # Project documentation.

Prerequisites
Go (1.19 or higher)
gRPC (gRPC-Go libraries)
MySQL (for database storage)
Protobuf Compiler (for generating Go code from .proto files)
