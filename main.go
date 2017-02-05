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
	name = createCmd.Arg("name", "Name for playbook").Required().String()e
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
		cmd := &commands.CreateCommand{ commands.Command{*name, *withHandlers, *withTemplates, *withFiles, *withVars, *withDefaults, *withMeta} }
		cmd.Execute()
	case "update":
		cmd := &commands.UpdateCommand{ commands.Update{*name, *withHandlers, *withTemplates, *withFiles, *withVars, *withDefaults, *withMeta} }
		cmd.Execute()
	default:
		fmt.Errorf("nothing called\n");
	}
}
