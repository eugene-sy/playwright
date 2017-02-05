package commands

import (
	"os"

	"com.github/axblade/playwright/utils"
)

type DeleteCommand struct {
	Command
}

func (self *DeleteCommand) Execute() (err error) {
	folders := self.SelectFolders()

	rolesPath, err := self.ReadRolesPath()
	if err != nil {
		return err
	}

	deletePlaybookStructure(rolesPath, self.Command.PlaybookName, folders)

	return nil
}

func deletePlaybookStructure(rolesPath string, name string, folders []string) {
	if string(rolesPath[len(rolesPath)-1]) != "/" {
		rolesPath = utils.Concat(rolesPath, "/")
	}

	// ToDo: throw erorr is playbook exists

	playbookPath := utils.Concat(rolesPath, name)

	if string(playbookPath[len(playbookPath)-1]) != "/" {
		playbookPath = utils.Concat(playbookPath, "/")
	}

	for _, folder := range folders {
		folderPath := utils.Concat(playbookPath, folder)

		os.RemoveAll(folderPath)
	}
}
