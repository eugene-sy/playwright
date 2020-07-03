package test

import (
	"fmt"
	"os"
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

func TestInvocationWithoutParameters(t *testing.T) {
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

func TestCreateWithDefaultOptions(t *testing.T) {
	roleName := "test"
	cmd := exec.Command(binary, "create", roleName)
	cmd.Dir = testFolder
	out, err := cmd.CombinedOutput()

	if err != nil {
		errStr := err.Error()
		t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", errStr)
	}

	outStr := string(out)
	if _, matchErr := regexp.MatchString("Role test was created successfully", outStr); matchErr != nil {
		t.Error("'playwright ", cmd, "' was expected to fail with suggestion to add 'name' parameter, but output was: ", outStr)
	}
}

func TestUpdate(t *testing.T) {
	roleName := "test"
	cmd := exec.Command(binary, "create", roleName)
	cmd.Dir = testFolder
	out, err := cmd.CombinedOutput()

	if err != nil {
		errStr := err.Error()
		t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", errStr)
	}

	cmd = exec.Command(binary, "update", roleName, "--vars")
	cmd.Dir = testFolder
	out, err = cmd.CombinedOutput()

	outStr := string(out)
	if _, matchErr := regexp.MatchString("Role test was updated successfully", outStr); matchErr != nil {
		t.Error("'playwright ", cmd, "' was expected to fail with suggestion to add 'name' parameter, but output was: ", outStr)
	}
}

func TestDelete(t *testing.T) {
	roleName := "test"
	cmd := exec.Command(binary, "create", roleName)
	cmd.Dir = testFolder
	out, err := cmd.CombinedOutput()

	if err != nil {
		errStr := err.Error()
		t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", errStr)
	}

	cmd = exec.Command(binary, "delete", roleName)
	cmd.Dir = testFolder
	out, err = cmd.CombinedOutput()

	outStr := string(out)
	if _, matchErr := regexp.MatchString("Role test was deleted successfully", outStr); matchErr != nil {
		t.Error("'playwright ", cmd, "' was expected to fail with suggestion to add 'name' parameter, but output was: ", outStr)
	}
}

const (
	testFolder = "/tmp/testdir"
	configFile = testFolder + "/ansible.cfg"
	config     = ""
)

func TestMain(m *testing.M) {
	createTestProjectStructure()
	code := m.Run()
	removeTestProjectStructure()
	os.Exit(code)
}

func createTestProjectStructure() {
	os.MkdirAll(testFolder, 0755)
	file, err := os.Create(configFile)
	if err != nil {
		fmt.Errorf("Could not create file %s", configFile)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(config); err != nil {
		fmt.Errorf("Could not write prefix to the file %s", configFile)
		return
	}

	_ = file.Sync()
	fmt.Println("Created test project folder")

	cd, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Could not find current dir: %s", err)
	}
	binary = cd + "/" + binary
}

func removeTestProjectStructure() {
	fmt.Println("Cleaning up test project folder")
	os.RemoveAll(testFolder)
}
