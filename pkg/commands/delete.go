package commands

import (
	"fmt"
	"os"

	"github.com/eugene-sy/playwright/pkg/logger"
	"github.com/eugene-sy/playwright/pkg/utils"
)

// DeleteCommand implements logic for a role update
// depending on the flags provided it
// - removes the role entirely
// - removes the parts of the role, except the tasks folder and main.yml in it
type DeleteCommand struct {
	Command
}

// DeleteCommand.Execute runs the requested filesystem tree updates, deleting requested parts of it
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

	if len(folders) != 1 {
		for _, folder := range folders {
			if folder == "tasks" {
				continue
			}
			folderPath := utils.Concat(playbookPath, folder)
			os.RemoveAll(folderPath)
			logger.LogSimple("Removed %s of the role %s", folder, name)
		}
	} else {
		os.RemoveAll(playbookPath)
		logger.LogSimple("Removed role %s with all contents", name)
	}

	return fmt.Sprintf("Role %s was deleted successfully", name), nil
}
