package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	"com.github/axblade/playwright/commands"
)

var (
	// Commands and args
	createCmd  = kingpin.Command("create", "Creates a playbook")
	createName = createCmd.Arg("name", "Name for playbook").Required().String()
	updateCmd  = kingpin.Command("update", "Updates a playbook")
	updateName = updateCmd.Arg("name", "Name for playbook").Required().String()
	deleteCmd  = kingpin.Command("delete", "Deletes a playbook or parts of it")
	deleteName = deleteCmd.Arg("name", "Name for playbook").Required().String()
	// Folder flags
	withHandlers  = kingpin.Flag("handlers", "Add 'handlers' folder").Bool()
	withTemplates = kingpin.Flag("templates", "Add 'templates' folder").Bool()
	withFiles     = kingpin.Flag("files", "Add 'files' folder").Bool()
	withVars      = kingpin.Flag("vars", "Add 'vars' folder").Bool()
	withDefaults  = kingpin.Flag("defaults", "Add 'defaults' folder").Bool()
	withMeta      = kingpin.Flag("meta", "Add 'meta' folder").Bool()
	all           = kingpin.Flag("all", "Apply action to all folders").Bool()
)

func main() {
	kingpin.Version("0.0.2")
	parsed := kingpin.Parse()

	var err error

	switch parsed {
	case "create":
		cmd := &commands.CreateCommand{commands.Command{*createName, *withHandlers, *withTemplates, *withFiles, *withVars, *withDefaults, *withMeta, *all}}
		err = cmd.Execute()
	case "update":
		cmd := &commands.UpdateCommand{commands.Command{*updateName, *withHandlers, *withTemplates, *withFiles, *withVars, *withDefaults, *withMeta, *all}}
		err = cmd.Execute()
	case "delete":
		cmd := &commands.DeleteCommand{commands.Command{*deleteName, *withHandlers, *withTemplates, *withFiles, *withVars, *withDefaults, *withMeta, *all}}
		err = cmd.Execute()
	default:
		fmt.Errorf("nothing called\n")
	}

	if err == nil {
		fmt.Println("Command executed successfully")
	} else {
		fmt.Printf("Error: %s\n", err)
	}
}
