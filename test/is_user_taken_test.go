package test

import (
	"net"
	"tcp-chat/server"
	"testing"
)

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
