package server

import "net"

// BroadcastMessage sends a message to all connected clients, excluding the sender
func broadcastMessage(message string, excludeConn net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	messageHistory = append(messageHistory, message)

	for conn := range Clients {
		if conn != excludeConn {
			conn.Write([]byte(message + "\n"))
		}
	}
}
