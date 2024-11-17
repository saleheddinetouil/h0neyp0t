package main

import (
	"log"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}


	logChan := make(chan LogEntry)

    go StartLogger(config, logChan)

	for _, service := range config.Services {
		go emulateService(service, logChan)
	}



	select {} // Keep the main goroutine running
}
