package server

import "net"

// SendMessageHistory sends previous messages to a new client
func sendMessageHistory(conn net.Conn) {
	Mutex.Lock()
	defer Mutex.Unlock()
	for _, msg := range MessageHistory {
		conn.Write([]byte(msg + "\n"))
	}
}
