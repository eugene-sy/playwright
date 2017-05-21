package commands

import (
	"fmt"
	"os"

	"github.com/Axblade/playwright/utils"
)

type DeleteCommand struct {
	Command
}

func (command *DeleteCommand) Execute() (success string, err error) {
	folders := command.SelectFolders()

	rolesPath, err := command.ReadRolesPath()
	if err != nil {
		return "", err
	}

	return deletePlaybookStructure(rolesPath, command.Command.PlaybookName, folders)
}

func deletePlaybookStructure(rolesPath string, name string, folders []string) (success string, err error) {
	if string(rolesPath[len(rolesPath)-1]) != "/" {
		rolesPath = utils.Concat(rolesPath, "/")
	}

	playbookPath := utils.Concat(rolesPath, name)

	if !utils.FolderExists(playbookPath) {
		return "", fmt.Errorf("Role %s does not exist", name)
	}

	if string(playbookPath[len(playbookPath)-1]) != "/" {
		playbookPath = utils.Concat(playbookPath, "/")
	}

	os.RemoveAll(playbookPath)
	log.LogSimple("Removed role %s with all contents", name)

	return fmt.Sprintf("Role %s was deleted successfully", name), nil
}
