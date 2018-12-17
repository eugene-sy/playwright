package commands

import (
	"fmt"
	"os"

	"github.com/eugene-sy/playwright/pkg/logger"
	"github.com/eugene-sy/playwright/pkg/utils"
)

type UpdateCommand struct {
	Command
}

func (command *UpdateCommand) Execute() (success string, err error) {
	folders := command.SelectFolders()

	rolesPath, err := command.ReadRolesPath()

	if err != nil {
		return "", err
	}

	return updatePlaybookStructure(rolesPath, command.Command.PlaybookName, folders)
}

func updatePlaybookStructure(rolesPath string, name string, folders []string) (success string, err error) {
	if string(rolesPath[len(rolesPath)-1]) != "/" {
		rolesPath = utils.Concat(rolesPath, "/")
	}

	playbookPath := utils.Concat(rolesPath, name)

	if !utils.FolderExists(playbookPath) {
		return "", fmt.Errorf("Role %s does not exists", name)
	}

	if string(playbookPath[len(playbookPath)-1]) != "/" {
		playbookPath = utils.Concat(playbookPath, "/")
	}

	for _, folder := range folders {
		if folder != "tasks" {
			folderPath := utils.Concat(playbookPath, folder)

			if utils.FolderExists(folderPath) {
				return "", fmt.Errorf("Folder %s already exists for role %s", folder, name)
			}

			os.MkdirAll(folderPath, 0755)
			logger.LogSimple("Created directory: %s", folder)

			if folder != "files" && folder != "templates" {
				filePath := utils.Concat(folderPath, "/main.yml")
				file, err := os.Create(filePath)

				if err != nil {
					return "", fmt.Errorf("Could not create file %s", filePath)
				}

				defer file.Close()

				if _, err := file.WriteString(YamlFilePrefix); err != nil {
					return "", fmt.Errorf("Could not write prefix to the file %s", filePath)
				}

				file.Sync()

				logger.LogSimple("Created main.yml for %s", folder)
			} else {
				logger.LogWarning("Skipped creation of main.yml for %s", folder)
			}
		}
	}

	return fmt.Sprintf("Role %s was updated successfully", name), nil
}