package server

import "net"

// SendMessageHistory sends previous messages to a new client
func sendMessageHistory(conn net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, msg := range messageHistory {
		conn.Write([]byte(msg + "\n"))
	}
}
