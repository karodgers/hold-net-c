package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// MaxClients is the maximum number of clients allowed in the chat
const MaxClients = 10

// MaxUsernameLength is the maximum number of characters allowed in a username
const MaxUsernameLength = 15

// HandleConnection handles communication with a connected client
func HandleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Welcome to TCP-Chat!\n"))

	// Send ASCII art
	SendAsciiArt(conn)

	// Check if the number of clients exceeds the maximum limit
	Mutex.Lock()
	if len(Clients) >= MaxClients {
		conn.Write([]byte("Sorry, the chat room is full. Please try again later.\n"))
		Mutex.Unlock()
		return
	}
	Mutex.Unlock()

	conn.Write([]byte("[ENTER YOUR NAME]: "))

	var name string
	for {
		nameInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		name = strings.TrimSpace(nameInput)

		// Check if the username is empty
		if name == "" {
			conn.Write([]byte("Name cannot be empty.\n[ENTER YOUR NAME]: "))
			continue
		}

		// Check if the username exceeds the max length
		if len(name) > MaxUsernameLength {
			conn.Write([]byte(fmt.Sprintf("Username cannot exceed %d characters.\n[ENTER YOUR NAME]: ", MaxUsernameLength)))
			continue
		}

		// Check if the username is already taken
		Mutex.Lock()
		if isUsernameTaken(name) {
			conn.Write([]byte("Username taken, try again.\n[ENTER YOUR NAME]: "))
			Mutex.Unlock()
			continue
		}
		// If the username is unique, add the client
		Clients[conn] = name
		Mutex.Unlock()
		break
	}

	// Send the message history to the new client
	sendMessageHistory(conn)

	// Broadcast the join message with a timestamp and log it
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	joinMessage := fmt.Sprintf("[%s][%s]: %s has joined our chat...", timestamp, name, name)
	BroadcastMessage(joinMessage, conn)

	// Log the join event
	logToFile(joinMessage)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		message = strings.TrimSpace(message)

		// Handle username change command
		if strings.HasPrefix(message, "/name ") {
			newName := strings.TrimPrefix(message, "/name ")

			// Validate the new username length
			if len(newName) > MaxUsernameLength {
				conn.Write([]byte(fmt.Sprintf("Username cannot exceed %d characters.\n", MaxUsernameLength)))
				continue
			}

			oldName, err := handleNameChange(conn, newName)
			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
				continue
			}

			// Update the local name variable
			name = newName

			// Notify other users of the username change
			timestamp := time.Now().Format("2006-01-02 15:04:05")
			changeMessage := fmt.Sprintf("[%s][%s]: %s changed their name to %s", timestamp, newName, oldName, newName)
			BroadcastMessage(changeMessage, conn)

			// Log the username change
			logToFile(changeMessage)
			continue
		}

		// Skip broadcasting if message is empty
		if message == "" {
			continue
		}

		// Generate timestamp for the message
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		// Create the timestamped message
		timestampedMessage := fmt.Sprintf("[%s][%s]: %s", timestamp, name, message)

		// Broadcast the timestamped message
		BroadcastMessage(timestampedMessage, conn)

		// Log the message sent
		logToFile(timestampedMessage)
	}

	// Remove the client when they disconnect
	Mutex.Lock()
	delete(Clients, conn)
	Mutex.Unlock()

	// Broadcast the leave message with a timestamp and log it
	timestamp = time.Now().Format("2006-01-02 15:04:05")
	leaveMessage := fmt.Sprintf("[%s][%s]: %s has left our chat", timestamp, name, name)
	BroadcastMessage(leaveMessage, conn)

	// Log the leave event
	logToFile(leaveMessage)
}
