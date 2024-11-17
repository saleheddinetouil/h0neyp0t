package main

import (
	"encoding/json"
	"os"
)

// ServiceConfig defines the configuration for a single service.
type ServiceConfig struct {
	Name   string `json:"name"`
	Port   int    `json:"port"`
	Banner string `json:"banner"` // For services like SSH
}


// Config struct represents the overall honeypot configuration.
type Config struct {
	Services []ServiceConfig `json:"services"`
    LogFile string `json:"logfile"`
}

// LoadConfig loads the configuration from the specified JSON file.
func LoadConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()


	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
