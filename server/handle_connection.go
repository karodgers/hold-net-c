package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// HandleConnection handles communication with a connected client
func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("Welcome to TCP-Chat!\n"))

	// Send ASCII art
	sendAsciiArt(conn)

	conn.Write([]byte("[ENTER YOUR NAME]: "))

	var name string
	for {
		nameInput, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		name = strings.TrimSpace(nameInput)
		if name == "" {
			conn.Write([]byte("Name cannot be empty.\n[ENTER YOUR NAME]: "))
			continue
		}
		// Check if username already exists
		mutex.Lock()
		if isUsernameTaken(name) {
			conn.Write([]byte("Username taken! Try again.\n[ENTER YOUR NAME]: "))
			mutex.Unlock()
			continue
		}
		// If username is unique, add the client
		clients[conn] = name
		mutex.Unlock()
		break
	}
	sendMessageHistory(conn)

	// Broadcast join message with timestamp and log it
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	joinMessage := fmt.Sprintf("[%s][%s]: %s has joined our chat...", timestamp, name, name)
	broadcastMessage(joinMessage, conn)

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
			oldName, err := handleNameChange(conn, newName)
			if err != nil {
				conn.Write([]byte(err.Error() + "\n"))
				continue
			}

			// Notify other users of the username change
			timestamp := time.Now().Format("2006-01-02 15:04:05")
			changeMessage := fmt.Sprintf("[%s][%s]: %s changed their name to %s", timestamp, oldName, oldName, newName)
			broadcastMessage(changeMessage, conn)

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
		broadcastMessage(timestampedMessage, conn)

		// Log the message sent
		logToFile(timestampedMessage)
	}

	// Remove the client when they disconnect
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()

	// Broadcast leave message with timestamp and log it
	timestamp = time.Now().Format("2006-01-02 15:04:05")
	leaveMessage := fmt.Sprintf("[%s][%s]: %s has left our chat", timestamp, name, name)
	broadcastMessage(leaveMessage, conn)

	// Log the leave event
	logToFile(leaveMessage)
}
