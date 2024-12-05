package server

import (
	"fmt"
	"log"
	"os"
	"time"
)

// logToFile saves the log message to a log file with a timestamp.
func logToFile(message string) {
	// Open the log file (create if doesn't exist, append if exists)
	file, err := os.OpenFile("server_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	// Create a new logger that writes to the log file
	logger := log.New(file, "", 0)

	// Get the current time for the timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Write the log message with the timestamp
	logger.Printf("[%s] %s\n", timestamp, message)
}
