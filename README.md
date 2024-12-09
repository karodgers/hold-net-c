# Net-Cat

This is a simple, multi-client TCP chat server written in Go. The server allows multiple clients to connect and chat with each other in real time. The server enforces a maximum client limit, username restrictions, and handles user commands like changing usernames.

## Features

- **Multi-client chat**: Multiple users can connect to the server and send messages in real-time.
- **Username management**: Clients can set a username (with a character length limit) and change it during the chat session.
- **Client connection limit**: The server allows up to a specified number of clients.
- **Message history**: When a user joins, they can view the previous chat history.
- **Username collision prevention**: If a username is already taken, users are prompted to choose a different one.
- **Timestamped messages**: Every message is timestamped to provide clarity on when each message was sent.

## Prerequisites

Before you can run the chat server, ensure you have the following installed:

- Go 1.18+ (for compiling and running the server)
- A terminal or command line interface

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/krodgers/net-cat.git
    cd net-cat
    ```

2. **Install the required dependencies** (if any, although this project uses standard Go libraries, no external dependencies are needed).

    ```bash
    go mod tidy
    ```

## Running the Server

To start the server, run the following command:

```bash
go run main.go
```

The server will start on the default port :8989. You can modify the port by specifying the port to run the server. 
```bash
go run main.go <port>
```
## Usage

Once connected to the server, you will see the following prompts:

    Enter your name: Type your username (must be between 1 and MaxUsernameLength characters).
    Chat messages: After joining, you can type your messages and send them to all other connected users.
    Change your name: To change your username, type /name newname where newname is the new name you'd like to use.
    Join messages: When a user joins, a message like [timestamp][username]: username has joined our chat... is broadcasted to all connected clients.
    Leave messages: When a user disconnects, a message like [timestamp][username]: username has left our chat is broadcasted to all other clients.

## Features Overview

    Max Clients: The server allows a maximum of 10 clients (configurable with the MaxClients constant). If the server is full, a new client will be notified that the chat room is full and asked to try again later.

    Username Rules: The maximum length of a username is restricted to 15 characters (configurable with the MaxUsernameLength constant). If a user tries to use an existing or too long username, they will be prompted to pick a different one.

    Message History: When a new client connects, they receive the history of previous messages sent to the chat room.

## Code Structure

    main.go: The entry point of the program. It initializes the server and listens for incoming connections.
    server.go: Contains the handleConnection function, which handles the communication with each client. It manages client connections, broadcasting messages, and user commands.

## License

This project is open-source and available under the MIT License. See the LICENSE file for more details.