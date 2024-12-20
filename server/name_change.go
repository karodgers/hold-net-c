package server

import (
	"fmt"
	"net"
	"strings"
)

// HandleNameChange handles the logic for changing a user's name
func HandleNameChange(conn net.Conn, newName string) (string, error) {
	// Trim spaces around the name
	newName = strings.TrimSpace(newName)

	// Check if the new name is valid and unique
	Mutex.Lock()
	defer Mutex.Unlock()

	if newName == "" {
		return "", fmt.Errorf("name cannot be empty")
	}
	if IsUsernameTaken(newName) {
		return "", fmt.Errorf("username taken! Try again")
	}

	// Get the old name to notify others
	oldName := Clients[conn]

	// Update the username
	Clients[conn] = newName

	// Return the old and new names for broadcasting purposes
	return oldName, nil
}
