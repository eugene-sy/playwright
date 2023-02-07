package commands

import (
	"fmt"
	"testing"
)

func TestNoOpCommand_Execute(t *testing.T) {
	cmd := NoOpCommand{}
	result, err := cmd.Execute()

	if result != "" {
		msg := fmt.Sprintf("NoOpCommand result was not empty. Result: '%s'", result)
		t.Error(msg)
	}
	if err == nil {
		t.Error("NoOpCommand did not return error")
	}
}
