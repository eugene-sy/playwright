package commands

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/Axblade/playwright/log"
	"github.com/Axblade/playwright/utils"
)

type Command struct {
	PlaybookName  string
	WithHandlers  bool
	WithTemplates bool
	WithFiles     bool
	WithVars      bool
	WithDefaults  bool
	WithMeta      bool
	All           bool
}

type ICommand interface {
	Execute() (success string, err error)
}

var (
	ANSIBLE_CONFIG_VAR = "ANSIBLE_CONFIG"
	ANSIBLE_CONFIG     = "./ansible.cfg"
	ANSIBLE_CONFIG_DOT = "./.ansible.cfg"
	ANSIBLE_CONFIG_OS  = "/etc/ansible/ansible.cfg"
)

// SelectFolders returns an array of folder names
// names are selected when relevant flag in command structure is set to TRUE
func (command *Command) SelectFolders() []string {
	result := []string{"tasks"}

	if command.WithHandlers {
		result = append(result, "handlers")
	}
	if command.WithTemplates {
		result = append(result, "templates")
	}
	if command.WithFiles {
		result = append(result, "files")
	}
	if command.WithVars {
		result = append(result, "vars")
	}
	if command.WithDefaults {
		result = append(result, "defaults")
	}
	if command.WithMeta {
		result = append(result, "meta")
	}

	return result
}

// ReadRolesPath finds path to roles directory and checks if the directory exists
// - from ansible configuration file if it is set
// - otherwise returns current directory followed by 'roles'
func (command *Command) ReadRolesPath() (rolesPath string, err error) {
	path, err := command.ansibleConfigPath()
	if err != nil {
		return "", errors.New("Cannot find Ansible configuration file")
	}

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", errors.New("Cannot open Ansible configuration file")
	}

	parts := strings.SplitAfter(path, "/")
	prefix := strings.Join(parts[:len(parts)-1], "")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "roles_path") {
			option := scanner.Text()
			rolesPath = strings.TrimSpace(strings.Split(option, "=")[1])
			return utils.Concat(prefix, rolesPath), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", errors.New("Cannot read data from Ansible configuration file")
	}

	log.LogWarning("Roles path was not found in configuration file, using default path.")

	return utils.Concat(prefix, "roles"), nil
}

// ansibleConfigPath checks if path to ansible config set
// as environment variable
// in current folder
// in OS deefault location
// returns location if found, error otherwise
func (command *Command) ansibleConfigPath() (path string, err error) {
	envPath := os.Getenv(ANSIBLE_CONFIG_VAR)

	if envPath != "" {
		return envPath, nil
	}

	if _, err := os.Stat(ANSIBLE_CONFIG); err == nil {
		return ANSIBLE_CONFIG, nil
	}

	if _, err := os.Stat(ANSIBLE_CONFIG_DOT); err == nil {
		return ANSIBLE_CONFIG_DOT, nil
	}

	if _, err := os.Stat(ANSIBLE_CONFIG_OS); err == nil {
		return ANSIBLE_CONFIG_OS, nil
	}

	return "", errors.New("Ansible config not found")
}
