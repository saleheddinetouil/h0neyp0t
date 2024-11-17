package main

import (
	"fmt"
	"net"
	"time"
    "bufio"
)


// emulateService emulates a single service on the specified port.
func emulateService(config ServiceConfig, logChan chan LogEntry) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logChan <- LogEntry{Timestamp: time.Now(), Service: config.Name, Event: fmt.Sprintf("Failed to listen on port %d: %s", config.Port, err)}
		return
	}
	defer listener.Close()



	logChan <- LogEntry{Timestamp: time.Now(), Service: config.Name, Event: fmt.Sprintf("Listening on port %d", config.Port)}



	for {
		conn, err := listener.Accept()


		if err != nil {
			logChan <- LogEntry{Timestamp: time.Now(), Service: config.Name, Event: fmt.Sprintf("Failed to accept connection: %s", err)}


			continue
		}


		go handleConnection(conn, config, logChan)
	}
}

// handleConnection handles a single client connection.
func handleConnection(conn net.Conn, config ServiceConfig, logChan chan LogEntry) {
	defer conn.Close()


	remoteAddr := conn.RemoteAddr().String()
	logChan <- LogEntry{Timestamp: time.Now(), Service: config.Name, Event: fmt.Sprintf("Connection from %s", remoteAddr), RemoteAddr: remoteAddr}


	if config.Banner != "" {
		_, err := conn.Write([]byte(config.Banner + "\r\n"))



		if err != nil {
			logChan <- LogEntry{Timestamp: time.Now(), Service: config.Name, Event: fmt.Sprintf("Error writing banner to %s : %s", remoteAddr, err), RemoteAddr: remoteAddr}
		}
	}



	for {
		scanner := bufio.NewScanner(conn)



		if ok := scanner.Scan(); !ok { // Check if there was an error during scanning

			break

		}
		command := scanner.Text()
		logChan <- LogEntry{Timestamp: time.Now(), Service: config.Name, Event: fmt.Sprintf("Command received from %s: %s", remoteAddr, command), RemoteAddr: remoteAddr, Command: command}

	}
    logChan <- LogEntry{Timestamp: time.Now(), Service: config.Name, Event: fmt.Sprintf("Connection from %s closed", remoteAddr), RemoteAddr: remoteAddr}



}
