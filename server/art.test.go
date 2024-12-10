package server

import (
	"net"
	"testing"
)

func TestSendAsciiArt(t *testing.T) {
}

type mockConnection struct {
	net.Conn
}

func (m *mockConnection) Write(p []byte) (n int, err error) {
	return 0, nil
}
