package commands

import (
	"errors"
	"fmt"
	"os"

	"com.github/axblade/playwright/utils"
)

type DeleteCommand struct {
	Command
}

func (command *DeleteCommand) Execute() (err error) {
	folders := command.SelectFolders()

	rolesPath, err := command.ReadRolesPath()
	if err != nil {
		return err
	}

	return deletePlaybookStructure(rolesPath, command.Command.PlaybookName, folders)
}

func deletePlaybookStructure(rolesPath string, name string, folders []string) (err error) {
	if string(rolesPath[len(rolesPath)-1]) != "/" {
		rolesPath = utils.Concat(rolesPath, "/")
	}

	playbookPath := utils.Concat(rolesPath, name)

	if !utils.FolderExists(playbookPath) {
		return errors.New(fmt.Sprintf("Role %s does not exist", name))
	}

	if string(playbookPath[len(playbookPath)-1]) != "/" {
		playbookPath = utils.Concat(playbookPath, "/")
	}

	os.RemoveAll(playbookPath)

	return nil
}
