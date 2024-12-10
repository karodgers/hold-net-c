package test

import (
	"net"
	"testing"

	"tcp-chat/server"
)

type mockConnection struct {
	net.Conn
	WriteFunc func([]byte) (int, error)
}

func (m *mockConnection) Write(p []byte) (n int, err error) {
	if m.WriteFunc != nil {
		return m.WriteFunc(p)
	}
	return 0, nil
}

// Test correct sending of ASCII art to connection
func TestSendAsciiArt(t *testing.T) {
	var writtenData []byte
	mockConn := &mockConnection{
		WriteFunc: func(p []byte) (n int, err error) {
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

func TestSendAsciiArt_NetworkError(t *testing.T) {
	mockConn := &mockConnection{}
	mockConn.WriteFunc = func(p []byte) (n int, err error) {
		return 0, net.ErrClosed
	}

	server.SendAsciiArt(mockConn)
	// If the function completes without panicking, it's handling errors gracefully
}
