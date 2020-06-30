package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/eugene-sy/playwright/pkg/commands"
	"github.com/eugene-sy/playwright/pkg/logger"
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
	kingpin.Version("0.0.4")
	parsed := kingpin.Parse()

	var err error
	var success string

	switch parsed {
	case "create":
		cmd := &commands.CreateCommand{Command: commands.Command{PlaybookName: *createName, WithHandlers: *withHandlers, WithTemplates: *withTemplates, WithFiles: *withFiles, WithVars: *withVars, WithDefaults: *withDefaults, WithMeta: *withMeta, All: *all}}
		success, err = cmd.Execute()
	case "update":
		cmd := &commands.UpdateCommand{Command: commands.Command{PlaybookName: *updateName, WithHandlers: *withHandlers, WithTemplates: *withTemplates, WithFiles: *withFiles, WithVars: *withVars, WithDefaults: *withDefaults, WithMeta: *withMeta, All: *all}}
		success, err = cmd.Execute()
	case "delete":
		cmd := &commands.DeleteCommand{Command: commands.Command{PlaybookName: *deleteName, WithHandlers: *withHandlers, WithTemplates: *withTemplates, WithFiles: *withFiles, WithVars: *withVars, WithDefaults: *withDefaults, WithMeta: *withMeta, All: *all}}
		success, err = cmd.Execute()
	default:
		err = fmt.Errorf("Nothing was called, check --help command.\n")
	}

	if err == nil {
		logger.LogSuccess(success)
	} else {
		logger.LogError("Error: %s\n", err)
	}
}
