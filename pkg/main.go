package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/eugene-sy/playwright/pkg/commands"
	"github.com/eugene-sy/playwright/pkg/logger"
)

func main() {
	app := kingpin.New("playwright", "Command line utility for Ansible role management")
	app.Version("0.0.4")
	app.Author("Eugene Sypachev (@eugene-sy)")
	
	createCmd := app.Command("create", "Creates roles")
	createName := createCmd.Arg("name", "Name for role").Required().String()
	updateCmd := app.Command("update", "Updates roles")
	updateName := updateCmd.Arg("name", "Name for role").Required().String()
	deleteCmd := app.Command("delete", "Deletes roles")
	deleteName := deleteCmd.Arg("name", "Name for role").Required().String()
	// Folder flags
	withHandlers := app.Flag("handlers", "Add 'handlers' folder").Bool()
	withTemplates := app.Flag("templates", "Add 'templates' folder").Bool()
	withFiles := app.Flag("files", "Add 'files' folder").Bool()
	withVars := app.Flag("vars", "Add 'vars' folder").Bool()
	withDefaults := app.Flag("defaults", "Add 'defaults' folder").Bool()
	withMeta := app.Flag("meta", "Add 'meta' folder").Bool()
	all := app.Flag("all", "Apply action to all folders").Bool()
	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))

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
		app.Terminate(func(code int) {
			logger.LogError("Error: %s\n", err)
			os.Exit(code)
		})
		app.FatalIfError(err, "")
	}
}
