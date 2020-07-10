package utils

import (
	"os"
	"strconv"
)

const (
	// SystemNoColor - system-wide no-color env variable
	SystemNoColor = "NO_COLOR"
	// PlaywrightNoColor - application level no-color env variable
	PlaywrightNoColor = "PLAYWRIGHT_NOCOLOR"
	// Term - system-wide terminal env variable
	Term = "TERM"
	// DumbTerminalEnvVarValue - dumb value for terminal env variable
	DumbTerminalEnvVarValue = "dumb"
	// AnsibleConfigVar - ansible config file env variable
	AnsibleConfigVar = "ANSIBLE_CONFIG"
	// AnsibleConfigVar - ansible config file env variable containing path to roles
	AnsibleRolesPath = "ANSIBLE_ROLES_PATH"
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
