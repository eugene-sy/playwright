package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/Axblade/playwright/commands"
	"github.com/Axblade/playwright/log"
)

var (
	// Commands and args
	createCmd  = kingpin.Command("create", "Creates roles")
	createName = createCmd.Arg("name", "Name for role").Required().String()
	updateCmd  = kingpin.Command("update", "Updates roles")
	updateName = updateCmd.Arg("name", "Name for role").Required().String()
	deleteCmd  = kingpin.Command("delete", "Deletes roles")
	deleteName = deleteCmd.Arg("name", "Name for role").Required().String()
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
	kingpin.Version("0.0.3")
	parsed := kingpin.Parse()

	var err error
	var success string

	switch parsed {
	case "create":
		cmd := &commands.CreateCommand{commands.Command{*createName, *withHandlers, *withTemplates, *withFiles, *withVars, *withDefaults, *withMeta, *all}}
		success, err = cmd.Execute()
	case "update":
		cmd := &commands.UpdateCommand{commands.Command{*updateName, *withHandlers, *withTemplates, *withFiles, *withVars, *withDefaults, *withMeta, *all}}
		success, err = cmd.Execute()
	case "delete":
		cmd := &commands.DeleteCommand{commands.Command{*deleteName, *withHandlers, *withTemplates, *withFiles, *withVars, *withDefaults, *withMeta, *all}}
		success, err = cmd.Execute()
	default:
		err = fmt.Errorf("Nothing was called, check --help command.\n")
	}

	if err == nil {
		log.LogSuccess(success)
	} else {
		log.LogError("Error: %s\n", err)
	}
}
