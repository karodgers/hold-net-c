package server

import (
	"fmt"
	"net"
	"sync"
)

var (
	clients        = make(map[net.Conn]string)
	mutex          = &sync.Mutex{}
	messageHistory []string
)

// StartServer starts the TCP server and listens for incoming connections
func StartServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server started on port %d...\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}
