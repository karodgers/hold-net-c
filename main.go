package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"tcp-chat/server"
)

func main() {
	port := 8989 // Default port

	// Check if exactly one argument is provided for the port
	if len(os.Args) == 2 {
		p, err := strconv.Atoi(os.Args[1]) // Try to convert the first argument to an integer
		if err != nil {
			fmt.Println("[USAGE]: ./TCPChat $port")
			return
		}
		port = p
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	log.Printf("Listening on the port :%d\n", port)
	server.StartServer(port)
}
