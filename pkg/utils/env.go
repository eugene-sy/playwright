package utils

import (
	"os"
	"strconv"
)

// GetEnvBool - returns a value of environment variable if it is set or default value otherwise
func GetEnvBool(key string, defaultValue bool) bool {
	s := os.Getenv(key)
	value, err := strconv.ParseBool(s)
	if err != nil {
		return defaultValue
	}
	return value
}
