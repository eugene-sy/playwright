package commands

import (
	"os"

	"com.github/axblade/playwright/utils"
)

type Command struct {
	PlaybookName string
	WithHandlers bool
	WithTemplates bool
	WithFiles bool
	WithVars bool
	WithDefaults bool
	WithMeta bool
}

type ICommand interface {
	Execute() (err error)
}

func (self *Command) SelectFolders() []string {
	result := []string{"tasks"}

	if self.WithHandlers {
		result = append(result, "handlers")
	}
	if self.WithTemplates {
		result = append(result, "templates")
	}
	if self.WithFiles {
		result = append(result, "files")
	}
	if self.WithVars {
		result = append(result, "vars")
	}
	if self.WithDefaults {
		result = append(result, "defaults")
	}
	if self.WithMeta {
		result = append(result, "meta")
	}

	return result
}

func (self *Command) ReadRolesPath() (rolesPath string, err error) {
	path, err := self.AnsibleConfigPath()
	if err != nil {
		return "", error.New("Cannot find Ansible configuration file")
	}

	file, err := os.Open(path)
	if err != nil {
		return "", errors.New("Cannot open Ansible configuration file")
	}
	defer file.Close()

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

	return utils.Concat(prefix, "roles"), nil
}

func (self *Command) AnsibleConfigPath() (path string, err error) {
	envPath := os.Getenv("ANSIBLE_CONFIG")

	if envPath != "" {
		return envPath, nil
	}

	if _, err := os.Stat("./ansible.cfg"); err == nil {
		return "./ansible.cfg", nil
	}

	if _, err := os.Stat("./.ansible.cfg"); err == nil {
		return "./.ansible.cfg", nil
	}

	if _, err := os.Stat("/etc/ansible/ansible.cfg"); err == nil {
		return "/etc/ansible/ansible.cfg", nil
	}

	return "", errors.New("Ansible config not found")
}
