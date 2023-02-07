package commands

import (
	"fmt"
	"os"

	"github.com/eugene-sy/playwright/pkg/logger"
	"github.com/eugene-sy/playwright/pkg/utils"
)

// CreateCommand implements logic for a role creation
// depending on the flags provided it creates
// - tasks folder with main.yml file in it
// - any other folders/files used in playbooks
// main.yml file is created with yaml document separator in the beginning
type CreateCommand struct {
	CommandConfiguration
	PlaybookName string
}

// CreateCommand.Execute creates the requested filesystem tree
func (command *CreateCommand) Execute() (success string, err error) {
	folders := command.SelectFolders()

	rolesPath, err := command.ReadRolesPath()
	if err != nil {
		return "", err
	}

	return createPlaybookStructure(rolesPath, command.PlaybookName, folders)
}

func createPlaybookStructure(rolesPath string, name string, folders []string) (success string, err error) {
	if string(rolesPath[len(rolesPath)-1]) != "/" {
		rolesPath = utils.Concat(rolesPath, "/")
	}

	playbookPath := utils.Concat(rolesPath, name)

	if utils.FolderExists(playbookPath) {
		return "", fmt.Errorf("role %s already exists", name)
	}

	if string(playbookPath[len(playbookPath)-1]) != "/" {
		playbookPath = utils.Concat(playbookPath, "/")
	}

	for _, folder := range folders {
		folderPath := utils.Concat(playbookPath, folder)
		_ = os.MkdirAll(folderPath, 0755)
		logger.LogSimple("Created directory: %s", folder)

		if folder != "files" && folder != "templates" {
			filePath := utils.Concat(folderPath, "/main.yml")
			file, err := os.Create(filePath)

			if err != nil {
				return "", fmt.Errorf("could not create file %s", filePath)
			}

			defer file.Close()

			if _, err := file.WriteString(YamlFilePrefix); err != nil {
				return "", fmt.Errorf("could not write prefix to the file %s", filePath)
			}

			file.Sync()

			logger.LogSimple("Created main.yml for %s", folder)
		} else {
			logger.LogWarning("Skipped creation of main.yml for %s", folder)
		}
	}

	return fmt.Sprintf("Role %s was created successfully", name), nil
}
