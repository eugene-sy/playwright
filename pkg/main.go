package main

import (
	"github.com/eugene-sy/playwright/pkg/utils"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"

	"github.com/eugene-sy/playwright/pkg/commands"
	"github.com/eugene-sy/playwright/pkg/logger"
)

const (
	CreateCommand = "create"
	UpdateCommand = "update"
	DeleteCommand = "delete"
)

func main() {
	app := kingpin.New("playwright", "Command line utility for Ansible role management")
	app.Version("0.0.4")
	app.Author("Eugene Sypachev (@eugene-sy)")

	createCmd := app.Command(CreateCommand, "Creates roles")
	createName := createCmd.Arg("name", "Name for role").Required().String()
	updateCmd := app.Command(UpdateCommand, "Updates roles")
	updateName := updateCmd.Arg("name", "Name for role").Required().String()
	deleteCmd := app.Command(DeleteCommand, "Deletes roles")
	deleteName := deleteCmd.Arg("name", "Name for role").Required().String()
	// Folder flags
	withHandlers := app.Flag("handlers", "Add 'handlers' folder").Bool()
	withTemplates := app.Flag("templates", "Add 'templates' folder").Bool()
	withFiles := app.Flag("files", "Add 'files' folder").Bool()
	withVars := app.Flag("vars", "Add 'vars' folder").Bool()
	withDefaults := app.Flag("defaults", "Add 'defaults' folder").Bool()
	withMeta := app.Flag("meta", "Add 'meta' folder").Bool()
	all := app.Flag("all", "Apply action to all folders").Bool()
	noColor := app.Flag("no-color", "Disable color output").Bool()
	parsed := kingpin.MustParse(app.Parse(os.Args[1:]))

	configureLogging(noColor)

	var err error
	var success string
	var cmd commands.ICommand
	commandConfiguration := commands.CommandConfiguration{
		WithHandlers:  *withHandlers,
		WithTemplates: *withTemplates,
		WithFiles:     *withFiles,
		WithVars:      *withVars,
		WithDefaults:  *withDefaults,
		WithMeta:      *withMeta,
		All:           *all,
	}

	switch parsed {
	case CreateCommand:
		cmd = &commands.CreateCommand{CommandConfiguration: commandConfiguration, PlaybookName: *createName}
	case UpdateCommand:
		cmd = &commands.UpdateCommand{CommandConfiguration: commandConfiguration, PlaybookName: *updateName}
	case DeleteCommand:
		cmd = &commands.DeleteCommand{CommandConfiguration: commandConfiguration, PlaybookName: *deleteName}
	default:
		cmd = &commands.NoOpCommand{}
	}
	success, err = cmd.Execute()

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

func configureLogging(noColor *bool) {
	osNoColorEnv := utils.GetEnvBool(utils.SystemNoColor, false)
	playwrightNoColorEnv := utils.GetEnvBool(utils.PlaywrightNoColor, false)
	termEnv := os.Getenv(utils.Term)
	useNoColor := *noColor || playwrightNoColorEnv || osNoColorEnv || termEnv == utils.DumbTerminalEnvVarValue
	logger.Configure(useNoColor)
}
