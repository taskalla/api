package env

import (
	"os"
	"strconv"
)

func String(name, fallback string) string {
	result := os.Getenv(name)
	if result == "" {
		return fallback
	}

	return result
}

func Int(name string, fallback int) int {
	resultString := os.Getenv(name)
	if resultString == "" {
		return fallback
	}

	result, err := strconv.Atoi(resultString)
	if err != nil {
		return fallback
	}

	return result
}
