package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	"com.github/axblade/playwright/commands"
)

var (
	// Commands and args
	createCmd = kingpin.Command("create", "Creates a playbook").Default()
	name = createCmd.Arg("name", "Name for playbook").Required().String()
	// Folder flags
	withHandlers = kingpin.Flag("handlers", "Add 'handlers' folder").Bool()
	withTemplates = kingpin.Flag("templates", "Add 'templates' folder").Bool()
	withFiles = kingpin.Flag("files", "Add 'files' folder").Bool()
	withVars = kingpin.Flag("vars", "Add 'vars' folder").Bool()
	withDefaults = kingpin.Flag("defaults", "Add 'defaults' folder").Bool()
	withMeta = kingpin.Flag("meta", "Add 'meta' folder").Bool()
)

func main() {
	kingpin.Version("0.0.2")
	parsed := kingpin.Parse()

	switch parsed {
	case "create":
		fmt.Printf("create called\n");
		cmd := &commands.CreateCommand{ commands.Command{*name, *withHandlers, *withTemplates, *withFiles, *withVars, *withDefaults, *withMeta} }
		cmd.Execute()
	default:
		fmt.Errorf("nothing called\n");
	}

	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("Provide name of the playbook as first argument")
		return
	}

	roleName := args[0]

	folders := selectFolders(args)

	path, err := ansibleConfigPath()
	if err != nil {
		fmt.Println("Cannot find Ansible configuration file")
		return
	}

	rolesPath, err := readRolesPath(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Roles path is:", rolesPath)

	createPlaybookStructure(rolesPath, roleName, folders)
}

func selectFolders(args []string) []string {
	result := []string{"tasks"}

	for _, arg := range args {
		if checkKey(arg, "--with-handlers") {
			result = append(result, "handlers")
		}
		if checkKey(arg, "--with-templates") {
			result = append(result, "templates")
		}
		if checkKey(arg, "--with-files") {
			result = append(result, "files")
		}
		if checkKey(arg, "--with-vars") {
			result = append(result, "vars")
		}
		if checkKey(arg, "--with-defaults") {
			result = append(result, "defaults")
		}
		if checkKey(arg, "--with-meta") {
			result = append(result, "meta")
		}
	}

	return result
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
	prefix := strings.Join(parts[:len(parts)-1], "")

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


func createPlaybookStructure(rolesPath string, name string, folders []string) {
	if string(rolesPath[len(rolesPath)-1]) != "/" {
		rolesPath = concat(rolesPath, "/")
	}

	playbookPath := concat(rolesPath, name)

	if string(playbookPath[len(playbookPath)-1]) != "/" {
		playbookPath = concat(playbookPath, "/")
	}

	for _, folder := range folders {
		folderPath := concat(playbookPath, folder)
		os.MkdirAll(folderPath, 0755)

		if folder != "files" && folder != "templates" {
			filePath := concat(folderPath, "/main.yml")
			os.Create(filePath)
		}
	}
}
