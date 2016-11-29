package commands

import (
	"fmt"
)

type CreateCommand struct {
	Command
}

func (self *CreateCommand) Execute() (err error) {
	fmt.Println("Execute of CreateCommand is called")
	return nil
}
