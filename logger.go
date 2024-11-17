package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// LogEntry represents a single log entry.
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
	Event     string    `json:"event"`
	RemoteAddr string `json:"remote_addr"`
    Command string `json:"command"`
}

// logEvent logs a single event to the log file.
func logEvent(config *Config, entry LogEntry) {
	logFile, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)



	if err != nil {
		log.Println("Error opening log file:", err)
		return
	}


	defer logFile.Close()



	jsonData, err := json.Marshal(entry)



	if err != nil {
		log.Println("Error marshaling log entry:", err)

		return
	}
	fmt.Fprintln(logFile, string(jsonData))


}



func StartLogger(config *Config, logChan chan LogEntry) {


    for entry := range logChan {


            logEvent(config, entry)

    }
}
