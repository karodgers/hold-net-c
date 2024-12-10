package server

import "net"

// BroadcastMessage sends a message to all connected clients, excluding the sender
func BroadcastMessage(message string, excludeConn net.Conn) {
	Mutex.Lock()
	defer Mutex.Unlock()
	MessageHistory = append(MessageHistory, message)

	for conn := range Clients {
		if conn != excludeConn {
			conn.Write([]byte(message + "\n"))
		}
	}
}
