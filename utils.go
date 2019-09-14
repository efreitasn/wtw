package main

import (
	"os"
	"path"
)

func configUnitToAPIUnit(configUnit string) string {
	switch configUnit {
	case TempCelsius:
		return "metric"
	case TempFahrenheit:
		return "imperial"
	case TempKelvin:
		fallthrough
	default:
		return ""
	}
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return path.Join(homeDir, ".wtw"), nil
}
