package main

import "fmt"
import "os"
import "errors"
import "bufio"
import "strings"
import "bytes"

func main() {
	path, err := ansibleConfigPath()
	if err != nil {
		fmt.Println("Cannot find Ansible configuration file")
	}

	rolesPath, err := readRolesPath(path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Roles path is:", rolesPath)
}

func ansibleConfigPath() (path string, err error) {
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

func readRolesPath(path string) (rolesPath string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", errors.New("Cannot open Ansible configuration file")
	}
	defer file.Close()

	parts := strings.SplitAfter(path, "/")
	fmt.Println(parts)
	prefix := strings.Join(parts[:len(parts)-1], "")
	fmt.Println(prefix)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "roles_path") {
			option := scanner.Text()
			rolesPath = strings.TrimSpace(strings.Split(option, "=")[1])
			return concat(prefix, rolesPath), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", errors.New("Cannot read data from Ansible configuration file")
	}

	return concat(prefix, "roles"), nil
}

func concat(prefix string, suffix string) string {
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(suffix)
	return buffer.String()
}
