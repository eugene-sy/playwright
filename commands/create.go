package commands

import (
	"fmt"

	"com.github/axblade/playwright/utils"
)

type CreateCommand struct {
	Command
}

func (self *CreateCommand) Execute() (err error) {
	folders := self.SelectFolders()

	rolesPath, err := self.ReadRolesPath()
	if err != nil {
		return err
	}

	createPlaybookStructure(rolesPath, self.Command.Name, folders)

	return nil
}

func createPlaybookStructure(rolesPath string, name string, folders []string) {
	if string(rolesPath[len(rolesPath)-1]) != "/" {
		rolesPath = utils.Concat(rolesPath, "/")
	}

	playbookPath := concat(rolesPath, name)

	if string(playbookPath[len(playbookPath)-1]) != "/" {
		playbookPath = utils.Concat(playbookPath, "/")
	}

	for _, folder := range folders {
		folderPath := utils.Concat(playbookPath, folder)
		os.MkdirAll(folderPath, 0755)

		if folder != "files" && folder != "templates" {
			filePath := utils.Concat(folderPath, "/main.yml")
			os.Create(filePath)
		}
	}
}
