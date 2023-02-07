package commands

import "fmt"

// NoOpCommand represents invocation with unknown or incomplete command
type NoOpCommand struct{}

// NoOpCommand.Execute always fails suggesting usage of the help flag
func (command *NoOpCommand) Execute() (success string, err error) {
	return "", fmt.Errorf("no command was called, use --help to get list of available options")
}
