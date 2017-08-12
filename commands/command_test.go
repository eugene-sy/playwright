package commands

import (
	"testing"
	"reflect"
	"fmt"
)

func TestAvailableRolesPathWithSingleInput(t *testing.T) {
	first := "/path/to/roles"
	input := "roles_path=" + first

	result := availableRolesPath(input)

	expected := []string{ first }
	if !reflect.DeepEqual(result, expected) {
		message := fmt.Sprintf("Function returned wrong array, %v, expected: %v", result, expected)
		t.Error(message)
	}
}

func TestAvailableRolesPathWithMultipleInput(t *testing.T) {
	first := "/path/to/roles"
	second := "/other/path/to/roles"
	input := "roles_path=" + first + ":" + second

	result := availableRolesPath(input)

	expected := []string{ first, second }
	if !reflect.DeepEqual(result, expected) {
		message := fmt.Sprintf("Function returned wrong array, %v, expected: %v", result, expected)
		t.Error(message)
	}
}
