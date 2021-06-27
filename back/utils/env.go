package utils

import "os"

func GetFromEnvOrDefault(variable string, defaultValue string) string {
	value := os.Getenv(variable)
	if value == "" {
		return defaultValue
	}

	return value
}
