package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBConnectionString string `json:"db_connection_string"`
}

func LoadConfig(filename string) (Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	return config, err
}
