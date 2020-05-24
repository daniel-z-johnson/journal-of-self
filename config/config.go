package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Database *DatabaseSQL `json:"database"`
}

type DatabaseSQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// LoadConfig - give it a config file location and it will
// load that config file and return a Configuration struct
// or if there was an error an error will be returned
func LoadConfig(fileName string) (*Config, error) {
	var config Config
	f1, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f1.Close()

	decoder := json.NewDecoder(f1)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
