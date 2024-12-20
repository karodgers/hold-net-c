package test

import (
	"net"
	"testing"

	"tcp-chat/server"
)

// Test output when username is taken
func TestIsUsernameTaken_ExistingUsername(t *testing.T) {
	// Save the original Clients map and restore it after the test
	originalClients := server.Clients
	defer func() { server.Clients = originalClients }()

	// Set up a test Clients map
	server.Clients = map[net.Conn]string{
		&mockConnection{}: "Alice",
		&mockConnection{}: "Bob",
		&mockConnection{}: "Charlie",
	}

	// Test with an existing username
	result := server.IsUsernameTaken("Bob")

	if !result {
		t.Errorf("isUsernameTaken(\"Bob\") = %v; want true", result)
	}
}

// Test output when username is not taken
func TestIsUsernameTaken_NotTaken(t *testing.T) {
	// Clear the Clients map
	server.Clients = make(map[net.Conn]string)

	// Add some dummy clients
	server.Clients[&mockConnection{}] = "Atieno"
	server.Clients[&mockConnection{}] = "Bantu"

	// Check for a username that doesn't exist
	result := server.IsUsernameTaken("Caro")

	if result {
		t.Errorf("Expected isUsernameTaken to return false for a non-existent username, but got true")
	}
}

// Test for case-sensitivity
func TestIsUsernameTaken_CaseSensitive(t *testing.T) {
	// Clear the Clients map before the test
	server.Clients = make(map[net.Conn]string)

	// Add a test client
	server.Clients[&mockConnection{}] = "TestUser"

	// Test with the exact same username
	if !server.IsUsernameTaken("TestUser") {
		t.Error("Expected 'TestUser' to be taken, but it wasn't")
	}

	// Test with a different case
	if server.IsUsernameTaken("testuser") {
		t.Error("Expected 'testuser' to be available, but it was taken")
	}

	// Test with a completely different username
	if server.IsUsernameTaken("AnotherUser") {
		t.Error("Expected 'AnotherUser' to be available, but it was taken")
	}
}

// Test output when username is empty
func TestIsUsernameTaken_EmptyUsername(t *testing.T) {
	// Clear the Clients map to ensure a clean state
	server.Clients = make(map[net.Conn]string)

	// Test with an empty username
	result := server.IsUsernameTaken("")

	if result {
		t.Errorf("isUsernameTaken(\"\") = %v; want false", result)
	}
}

// Test output when username contains special characters
func TestIsUsernameTaken_SpecialCharacters(t *testing.T) {
	// Clear the Clients map
	server.Clients = make(map[net.Conn]string)

	// Add a user with special characters
	specialUsername := "user@123!#$%"
	server.Clients[&mockConnection{}] = specialUsername

	// Test if the special username is taken
	if !server.IsUsernameTaken(specialUsername) {
		t.Errorf("Expected username '%s' to be taken, but it wasn't", specialUsername)
	}

	// Test if a different username is not taken
	differentUsername := "regular_user"
	if server.IsUsernameTaken(differentUsername) {
		t.Errorf("Expected username '%s' to not be taken, but it was", differentUsername)
	}
}
