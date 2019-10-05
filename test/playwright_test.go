package test

import (
	"os/exec"
	"regexp"
	"testing"
)

var (
	binary   = "../playwright"
	commands = []string{
		"create", "update", "delete",
	}
)

func TestVersionCall(t *testing.T) {
	cmd := exec.Command(binary, "--version")
	out, err := cmd.CombinedOutput()

	if err != nil {
		t.Error("'playwright --version' failed with ", err)
	}

	outStr := string(out)

	if _, err := regexp.MatchString("^\\d+\\.\\d+\\.\\d+$", outStr); err != nil {
		t.Error("'playwright --version' output is not a correct semver: ", outStr)
	}
}

func TestHelpCall(t *testing.T) {
	cmd := exec.Command(binary, "help")
	out, err := cmd.CombinedOutput()
	outStr := string(out)

	checkHelpOutput(t, outStr, err)
}

func TestHelpOptionCall(t *testing.T) {
	cmd := exec.Command(binary, "--help")
	out, err := cmd.CombinedOutput()
	outStr := string(out)

	checkHelpOutput(t, outStr, err)
}

func checkHelpOutput(t *testing.T, outStr string, err error) {
	if err != nil {
		t.Error("'playwright help' failed with ", err)
	}

	if _, err := regexp.MatchString("^usage: playwright [<flags>] <command> [<args> ...]", outStr); err != nil {
		t.Error("'playwright help' output is not a correct semver: ", outStr)
	}
}

func TestParameterlessInvocation(t *testing.T) {
	for _, command := range commands {
		cmd := exec.Command(binary, command)
		out, err := cmd.CombinedOutput()

		if err == nil {
			outStr := string(out)
			t.Error("'playwright ", command, "' was expected to fail, but it succeded with output: ", outStr)
		}

		errStr := err.Error()
		if _, matchErr := regexp.MatchString("required argument 'name' not provided", errStr); matchErr != nil {

			t.Error("'playwright ", command, "' was expected to fail with suggestion to add 'name' parameter, but output was: ", errStr)
		}
	}
}
