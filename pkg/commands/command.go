package commands

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/eugene-sy/playwright/pkg/logger"
	"github.com/eugene-sy/playwright/pkg/utils"
)

// Command represents user action with parameters
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

// ICommand -- command interface providing unified execution mechanism
type ICommand interface {
	Execute() (success string, err error)
}

const (
	AnsibleConfigVar = utils.AnsibleConfigVar
	AnsibleConfig    = "./ansible.cfg"
	AnsibleConfigDot = "./.ansible.cfg"
	AnsibleConfigOs  = "/etc/ansible/ansible.cfg"
	AnsibleRolesPath = utils.AnsibleRolesPath
	YamlFilePrefix   = "---\n"
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
// - from ANSIBLE_ROLES_PATH variable
// - from ansible configuration file if it is set
// - otherwise returns current directory followed by 'roles'
func (command *Command) ReadRolesPath() (rolesPath string, err error) {
	envRolesPath := os.Getenv(AnsibleRolesPath)

	if envRolesPath != "" {
		return envRolesPath, nil
	}

	return command.readRolesPathFromConfig()
}

// readRolesPathFromConfig - reads roles path from ansible config file
func (command *Command) readRolesPathFromConfig() (rolesPath string, err error) {
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

	defaultPath := utils.Concat(prefix, "roles")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		option := scanner.Text()
		if strings.Contains(option, "roles_path") {
			paths := availableRolesPath(option)
			logger.LogSimple("%v", paths)

			if len(paths) == 0 {
				return defaultPath, nil
			}

			if len(paths) > 1 {
				return makeUserSelectPath(paths)
			}

			return utils.Concat(prefix, paths[0]), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", errors.New("Cannot read data from Ansible configuration file")
	}

	logger.LogWarning("Roles path was not found in configuration file, using default path.")

	return defaultPath, nil
}

// ansibleConfigPath checks if path to ansible config set
// as environment variable
// in current folder
// in OS default location
// returns location if found, error otherwise
func (command *Command) ansibleConfigPath() (path string, err error) {
	envPath := os.Getenv(AnsibleConfigVar)

	if envPath != "" {
		return envPath, nil
	}

	if _, err := os.Stat(AnsibleConfig); err == nil {
		return AnsibleConfig, nil
	}

	if _, err := os.Stat(AnsibleConfigDot); err == nil {
		return AnsibleConfigDot, nil
	}

	if _, err := os.Stat(AnsibleConfigOs); err == nil {
		return AnsibleConfigOs, nil
	}

	return "", errors.New("Ansible config not found")
}

// availableRolesPath parses roles_path string into a set of roles paths
// roles_path='' is parsed into empty array
// roles_path=/something is parsed into array with one element '/something'
// roles_path=/something:/something-else is parsed into array of strings delimited by a ':'
func availableRolesPath(rolesPaths string) []string {
	options := strings.TrimSpace(strings.Split(rolesPaths, "=")[1])

	if len(options) == 0 {
		return []string{}
	}

	return strings.Split(options, ":")
}

// printMultipleRolesPathMessage prints message that multiple roles path found
// also prints path options
func printMultipleRolesPathMessage(rolesPaths []string) {
	logger.LogSimple("Configuration file contains multiple role paths: \n\n")

	for index, entry := range rolesPaths {
		logger.LogSimple("%d. %s", index+1, entry)
	}

	logger.LogSimple("\nPlease, select path where you want role to be created.")
}

// makeUserSelectPath asks user to select path in the array
// returns selected path
func makeUserSelectPath(rolesPaths []string) (path string, err error) {
	printMultipleRolesPathMessage(rolesPaths)

	reader := bufio.NewReader(os.Stdin)
	waitingForInput := true

	var index int
	for waitingForInput {
		logger.LogSimple("Enter path number: ")
		text, err := reader.ReadString('\n')

		if err != nil {
			return "", err
		}

		text = strings.Replace(text, "\n", "", -1)
		index, err = strconv.Atoi(text)
		if err != nil {
			logger.LogError("Input cannot be parsed, please, try again")
		} else if index < 1 || index > len(rolesPaths) {
			logger.LogError("You must enter a number from list, please, try again")
		} else {
			waitingForInput = false
		}
	}

	selected := rolesPaths[index-1]
	logger.LogSimple("Selected: %s", selected)

	return selected, nil
}
