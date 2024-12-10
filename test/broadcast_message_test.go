package test

import (
	"net"
	"testing"
	"time"

	"tcp-chat/server"
)

type MockConn struct {
	ReadData    []string
	WriteData   []string
	LastMessage string
}

func (m *MockConn) Read(b []byte) (n int, err error)   { return 0, nil }
func (m *MockConn) Write(b []byte) (n int, err error)  { m.LastMessage = string(b); return len(b), nil }
func (m *MockConn) Close() error                       { return nil }
func (m *MockConn) LocalAddr() net.Addr                { return nil }
func (m *MockConn) RemoteAddr() net.Addr               { return nil }
func (m *MockConn) SetDeadline(t time.Time) error      { return nil }
func (m *MockConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *MockConn) SetWriteDeadline(t time.Time) error { return nil }

func TestBroadcastMessage(t *testing.T) {
	// Clear the clients map and message history
	server.Clients = make(map[net.Conn]string)
	server.MessageHistory = []string{}

	// Create mock connections
	excludeConn := &MockConn{}
	conn1 := &MockConn{}
	conn2 := &MockConn{}
	conn3 := &MockConn{}

	// Add mock connections to clients map
	server.Clients[excludeConn] = "Excluded"
	server.Clients[conn1] = "Client1"
	server.Clients[conn2] = "Client2"
	server.Clients[conn3] = "Client3"

	// Test message
	testMessage := "Test broadcast message"

	// Call broadcastMessage
	server.BroadcastMessage(testMessage, excludeConn)

	// Check if the message was added to messageHistory
	if len(server.MessageHistory) != 1 || server.MessageHistory[0] != testMessage {
		t.Errorf("Message not added to messageHistory correctly")
	}

	// Check if the message was sent to all clients except excludeConn
	if excludeConn.LastMessage != "" {
		t.Errorf("Message was sent to excluded connection")
	}
	if conn1.LastMessage != testMessage+"\n" {
		t.Errorf("Message not sent to conn1")
	}
	if conn2.LastMessage != testMessage+"\n" {
		t.Errorf("Message not sent to conn2")
	}
	if conn3.LastMessage != testMessage+"\n" {
		t.Errorf("Message not sent to conn3")
	}
}
