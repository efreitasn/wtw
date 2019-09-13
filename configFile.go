package main

import (
	"encoding/json"
	"os"
	"path"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return path.Join(homeDir, ".wtw"), nil
}

// Config represents the config data related to accessing the OpenWeather API.
type Config struct {
	APIKey string `json:"apiKey"`
	CityID string `json:"cityID"`
}

// GetConfig returns the config data stored in the config file.
func GetConfig() (*Config, error) {
	filePath, err := getConfigFilePath()

	if err != nil {
		return nil, err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	jsonDec := json.NewDecoder(f)
	var c Config

	err = jsonDec.Decode(&c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Write writes the config data to the config file.
func (c *Config) Write() error {
	filePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	jsonEnc := json.NewEncoder(f)
	err = jsonEnc.Encode(c)

	if err != nil {
		return err
	}

	return nil
}
