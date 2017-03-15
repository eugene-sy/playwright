package commands

import (
	"errors"
	"fmt"
	"os"

	"com.github/axblade/playwright/utils"
)

type UpdateCommand struct {
	Command
}

func (self *UpdateCommand) Execute() (err error) {
	folders := self.SelectFolders()

	rolesPath, err := self.ReadRolesPath()

	if err != nil {
		return err
	}

	return updatePlaybookStructure(rolesPath, self.Command.PlaybookName, folders)
}

func updatePlaybookStructure(rolesPath string, name string, folders []string) (err error) {
	if string(rolesPath[len(rolesPath)-1]) != "/" {
		rolesPath = utils.Concat(rolesPath, "/")
	}

	playbookPath := utils.Concat(rolesPath, name)

	if !utils.FolderExists(playbookPath) {
		return errors.New(fmt.Sprintf("Role %s does not exists", name))
	}

	if string(playbookPath[len(playbookPath)-1]) != "/" {
		playbookPath = utils.Concat(playbookPath, "/")
	}

	for _, folder := range folders {
		if folder != "tasks" {
			folderPath := utils.Concat(playbookPath, folder)

			if utils.FolderExists(folderPath) {
				return errors.New(fmt.Sprintf("Folder %s already exists for role %s", folder, name))
			}

			os.MkdirAll(folderPath, 0755)

			if folder != "files" && folder != "templates" {
				filePath := utils.Concat(folderPath, "/main.yml")
				os.Create(filePath)
			}
		}
	}

	return nil
}
