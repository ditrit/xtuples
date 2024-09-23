package common

import (
	"os"
)

func ReadFromFile(filePath string) (string, error) {
	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(configBytes), nil
}
