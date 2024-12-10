package server

import (
	"fmt"
	"net"
	"strings"
)

// handleNameChange handles the logic for changing a user's name
func handleNameChange(conn net.Conn, newName string) (string, error) {
	// Trim spaces around the name
	newName = strings.TrimSpace(newName)

	// Check if the new name is valid and unique
	mutex.Lock()
	defer mutex.Unlock()

	if newName == "" {
		return "", fmt.Errorf("name cannot be empty")
	}
	if isUsernameTaken(newName) {
		return "", fmt.Errorf("username taken! Try again")
	}

	// Get the old name to notify others
	oldName := Clients[conn]

	// Update the username
	Clients[conn] = newName

	// Return the old and new names for broadcasting purposes
	return oldName, nil
}
