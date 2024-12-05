package utils

import (
	"log"
	"time"
)

// GetFormattedTime returns the current time in a readable format
func GetFormattedTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// HandleError logs an error if not nil
func HandleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
