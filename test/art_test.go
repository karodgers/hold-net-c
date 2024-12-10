package test

import (
	"net"
	"tcp-chat/server"
	"testing"
)

type mockConnection struct {
	net.Conn
	writeFunc func([]byte) (int, error)
}

func (m *mockConnection) Write(p []byte) (n int, err error) {
	if m.writeFunc != nil {
		return m.writeFunc(p)
	}
	return 0, nil
}

// Test correct sending of ASCII art to connection
func TestSendAsciiArt(t *testing.T) {
	var writtenData []byte
	mockConn := &mockConnection{
		writeFunc: func(p []byte) (n int, err error) {
			writtenData = append(writtenData, p...)
			return len(p), nil
		},
	}
	server.SendAsciiArt(mockConn)

	expectedArt := "         _nnnn_\n" +
		"        dGGGGMMb\n" +
		"       @p~qp~~qMb\n" +
		"       M|@||@) M|\n" +
		"       @,----.JM|\n" +
		"      JS^\\__/  qKL\n" +
		"     dZP        qKRb\n" +
		"    dZP          qKKb\n" +
		"   fZP            SMMb\n" +
		"   HZM            MMMM\n" +
		"   FqM            MMMM\n" +
		" __| \".        |\\dS\"qML\n" +
		" |    `.       | `' \\Zq\n" +
		"_)      \\.___.,|     .'\n" +
		"\\____   )MMMMMP|   .'\n" +
		"     `-'       `--'\n"

	if string(writtenData) != expectedArt {
		t.Errorf("Unexpected ASCII art sent. Got:\n%s\nWant:\n%s", string(writtenData), expectedArt)
	}
}
