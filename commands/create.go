package commands

import (
	"fmt"
	"os"

	"github.com/Axblade/playwright/logger"
	"github.com/Axblade/playwright/utils"
)

type CreateCommand struct {
	Command
}

func (command *CreateCommand) Execute() (success string, err error) {
	folders := command.SelectFolders()

	rolesPath, err := command.ReadRolesPath()
	if err != nil {
		return "", err
	}

	return createPlaybookStructure(rolesPath, command.Command.PlaybookName, folders)
}

func createPlaybookStructure(rolesPath string, name string, folders []string) (success string, err error) {
	if string(rolesPath[len(rolesPath)-1]) != "/" {
		rolesPath = utils.Concat(rolesPath, "/")
	}

	playbookPath := utils.Concat(rolesPath, name)

	if utils.FolderExists(playbookPath) {
		return "", fmt.Errorf("Role %s already exists", name)
	}

	if string(playbookPath[len(playbookPath)-1]) != "/" {
		playbookPath = utils.Concat(playbookPath, "/")
	}

	for _, folder := range folders {
		folderPath := utils.Concat(playbookPath, folder)
		os.MkdirAll(folderPath, 0755)
		logger.LogSimple("Created directory: %s", folder)

		if folder != "files" && folder != "templates" {
			filePath := utils.Concat(folderPath, "/main.yml")
			os.Create(filePath)
			logger.LogSimple("Created main.yml for %s", folder)
		} else {
			logger.LogWarning("Skipped creation of main.yml for %s", folder)
		}
	}

	return fmt.Sprintf("Role %s was created successfully", name), nil
}
