package main

import "fmt"
import "os"
import "errors"

func main() {
	if path, err := ansibleConfigPath(); err == nil {
		fmt.Println("Ansible configuration file found: ", path);
	} else {
		fmt.Println("Cannot find Ansible configuration file");
	}
}

func ansibleConfigPath() (path string, err error) {
	envPath := os.Getenv("ANSIBLE_CONFIG");
	if envPath != "" {
		return envPath, nil;
	}
	if _, err := os.Stat("./ansible.cfg"); err == nil {
		return "./ansible.cfg", nil;
	}
	if _, err := os.Stat("./.ansible.cfg"); err == nil {
		return "./.ansible.cfg", nil;
	}
	if _, err := os.Stat("/etc/ansible/ansible.cfg"); err == nil {
		return "/etc/ansible/ansible.cfg", nil;
	}
	return "", errors.New("Ansible config not found");
}
