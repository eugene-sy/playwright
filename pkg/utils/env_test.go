package utils

import (
	"os"
	"testing"
)

func TestGetenvBool(t *testing.T) {
	os.Setenv("DUMMY_BOOL", "true")
	value := GetEnvBool("DUMMY_BOOL", false)
	if !value {
		t.Errorf("GetEnvBool returned %t but true was expected", value)
	}
}

func TestGetenvBoolDefault(t *testing.T) {
	os.Unsetenv("DUMMY_BOOL")
	value := GetEnvBool("DUMMY_BOOL", false)
	if value {
		t.Errorf("GetEnvBool returned %t but false was expected", value)
	}
}

func TestGetenvBoolDefaultOnError(t *testing.T) {
	os.Setenv("DUMMY_BOOL", "not-boolean")
	value := GetEnvBool("DUMMY_BOOL", true)
	if !value {
		t.Errorf("GetEnvBool returned %t but true was expected", value)
	}
}
