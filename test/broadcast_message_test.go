package test

import (
	"net"
	"testing"
	"time"

	"tcp-chat/server"
)

type mockConn struct {
	lastMessage string
}

func (m *mockConn) Read(b []byte) (n int, err error)   { return 0, nil }
func (m *mockConn) Write(b []byte) (n int, err error)  { m.lastMessage = string(b); return len(b), nil }
func (m *mockConn) Close() error                       { return nil }
func (m *mockConn) LocalAddr() net.Addr                { return nil }
func (m *mockConn) RemoteAddr() net.Addr               { return nil }
func (m *mockConn) SetDeadline(t time.Time) error      { return nil }
func (m *mockConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *mockConn) SetWriteDeadline(t time.Time) error { return nil }

func TestBroadcastMessage(t *testing.T) {
	// Clear the clients map and message history
	server.Clients = make(map[net.Conn]string)
	server.MessageHistory = []string{}

	// Create mock connections
	excludeConn := &mockConn{}
	conn1 := &mockConn{}
	conn2 := &mockConn{}
	conn3 := &mockConn{}

	// Add mock connections to clients map
	Clients[excludeConn] = "Excluded"
	Clients[conn1] = "Client1"
	Clients[conn2] = "Client2"
	Clients[conn3] = "Client3"

	// Test message
	testMessage := "Test broadcast message"

	// Call broadcastMessage
	broadcastMessage(testMessage, excludeConn)

	// Check if the message was added to messageHistory
	if len(messageHistory) != 1 || messageHistory[0] != testMessage {
		t.Errorf("Message not added to messageHistory correctly")
	}

	// Check if the message was sent to all clients except excludeConn
	if excludeConn.lastMessage != "" {
		t.Errorf("Message was sent to excluded connection")
	}
	if conn1.lastMessage != testMessage+"\n" {
		t.Errorf("Message not sent to conn1")
	}
	if conn2.lastMessage != testMessage+"\n" {
		t.Errorf("Message not sent to conn2")
	}
	if conn3.lastMessage != testMessage+"\n" {
		t.Errorf("Message not sent to conn3")
	}
}
