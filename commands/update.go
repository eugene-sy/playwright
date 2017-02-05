package commands

import (
	"fmt"
	"strings"
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

	updatePlaybookStructure(rolesPath, self.Command.Name, folders)

	return nil
}

func updatePlaybookStructure(rolesPath string, name string, folders []string) (err error) {
	if string(rolesPath[len(rolesPath)-1]) != "/" {
		rolesPath = utils.Concat(rolesPath, "/")
	}

	// ToDo: check if role exists

	playbookPath := utils.Concat(rolesPath, name)

	if string(playbookPath[len(playbookPath)-1]) != "/" {
		playbookPath = utils.Concat(playbookPath, "/")
	}

	for _, folder := range folders {
		folderPath := utils.Concat(playbookPath, folder)
		// ToDo: check if folder exists, otehrwise throw error
		os.MkdirAll(folderPath, 0755)

		if folder != "files" && folder != "templates" {
			filePath := utils.Concat(folderPath, "/main.yml")
			os.Create(filePath)
		}
	}
}
