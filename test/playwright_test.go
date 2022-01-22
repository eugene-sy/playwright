package test

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"testing"
)

const (
	createCommand      = "create"
	updateCommand      = "update"
	deleteCommand      = "delete"
	handlersParam      = "handlers"
	templateParam      = "templates"
	filesParam         = "files"
	varsParam          = "vars"
	defaultsParam      = "defaults"
	metaParam          = "meta"
	testFolder         = "/tmp/testdir"
	configFile         = testFolder + "/ansible.cfg"
	config             = ""
	relativeBinaryPath = "../playwright"
)

var (
	binary   = ""
	commands = []string{
		createCommand, updateCommand, deleteCommand,
	}
	params = []string{
		handlersParam, templateParam, filesParam, varsParam, defaultsParam, metaParam,
	}
)

func TestMain(m *testing.M) {
	createTestProjectStructure()
	locateBinary()
	m.Run()
	removeTestProjectStructure()
}

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
		t.Run(fmt.Sprintf("TestInvocationWithoutParameters--%s", command), func(t *testing.T) {
			cmd := exec.Command(binary, command)
			out, err := cmd.CombinedOutput()

			if err == nil {
				outStr := string(out)
				t.Error("'playwright ", command, "' was expected to fail, but it succeeded with output: ", outStr)
			}

			errStr := err.Error()
			if _, matchErr := regexp.MatchString("required argument 'name' not provided", errStr); matchErr != nil {
				t.Error("'playwright ", command, "' was expected to fail with suggestion to add 'name' parameter, but output was: ", errStr)
			}
		})
	}
}

func TestCreateWithDefaultOptions(t *testing.T) {
	roleName := randomRoleName()
	cmd := exec.Command(binary, createCommand, roleName)
	cmd.Dir = testFolder
	out, err := cmd.CombinedOutput()

	if err != nil {
		t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", err.Error())
	}

	outStr := string(out)
	if _, matchErr := regexp.MatchString("created successfully", outStr); matchErr != nil {
		t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", outStr)
	}

	tasksFile := testFolder + "/roles/" + roleName + "/tasks/main.yml"
	if !fileExists(tasksFile) {
		t.Errorf("File was not created: %s", tasksFile)
	}
}

func TestCreateWithParams(t *testing.T) {
	for _, param := range params {
		t.Run(fmt.Sprintf("TestUpdateWithParams--%s", param), func(t *testing.T) {
			roleName := randomRoleName()
			createParam := fmt.Sprintf("--%s", param)
			cmd := exec.Command(binary, createCommand, roleName, createParam)
			cmd.Dir = testFolder
			out, err := cmd.CombinedOutput()

			if err != nil {
				t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", err.Error())
			}

			outStr := string(out)
			println(outStr)
			if _, matchErr := regexp.MatchString("created successfully", outStr); matchErr != nil {
				t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", outStr)
			}

			tasksFile := fmt.Sprintf("%s/roles/%s/tasks/main.yml", testFolder, roleName)
			if !fileExists(tasksFile) {
				t.Errorf("File was not created: %s", tasksFile)
			}

			var paramFile string
			if param != templateParam && param != filesParam {
				paramFile = fmt.Sprintf("%s/roles/%s/%s/main.yml", testFolder, roleName, param)
				if !fileExists(tasksFile) {
					t.Errorf("Create command flag used: [%s]. File was not created: %s", param, paramFile)
				}
			} else {
				paramFile = fmt.Sprintf("%s/roles/%s/%s", testFolder, roleName, param)
				if !isDirectory(paramFile) {
					t.Errorf("Create command flag used: [%s]. Directory was not created: %s", param, paramFile)
				}
			}
		})
	}
}

func TestUpdateWithParams(t *testing.T) {
	for _, param := range params {
		t.Run(fmt.Sprintf("TestUpdateWithParams--%s", param), func(t *testing.T) {
			roleName := randomRoleName()
			cmd := exec.Command(binary, createCommand, roleName)
			cmd.Dir = testFolder
			_, err := cmd.CombinedOutput()

			if err != nil {
				t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", err.Error())
			}
			updateParam := fmt.Sprintf("--%s", param)
			cmd = exec.Command(binary, updateCommand, roleName, updateParam)
			cmd.Dir = testFolder
			out, _ := cmd.CombinedOutput()

			outStr := string(out)
			if _, matchErr := regexp.MatchString("updated successfully", outStr); matchErr != nil {
				t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", outStr)
			}

			tasksFile := fmt.Sprintf("%s/roles/%s/tasks/main.yml", testFolder, roleName)
			if !fileExists(tasksFile) {
				t.Errorf("Update command flag used: [%s]. File was not created: %s", param, tasksFile)
			}

			var paramFile string
			if param != templateParam && param != filesParam {
				paramFile = fmt.Sprintf("%s/roles/%s/%s/main.yml", testFolder, roleName, param)
				if !fileExists(tasksFile) {
					t.Errorf("Update command flag used: [%s]. File was not created: %s", param, paramFile)
				}
			} else {
				paramFile = fmt.Sprintf("%s/roles/%s/%s", testFolder, roleName, param)
				if !isDirectory(paramFile) {
					t.Errorf("Update command flag used: [%s]. Directory was not created: %s", param, paramFile)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	roleName := randomRoleName()
	cmd := exec.Command(binary, createCommand, roleName)
	cmd.Dir = testFolder
	_, err := cmd.CombinedOutput()

	if err != nil {
		t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", err.Error())
	}

	cmd = exec.Command(binary, deleteCommand, roleName)
	cmd.Dir = testFolder
	out, _ := cmd.CombinedOutput()

	outStr := string(out)
	if _, matchErr := regexp.MatchString("was deleted successfully", outStr); matchErr != nil {
		t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", outStr)
	}

	roleFolder := fmt.Sprintf("%s/roles/%s", testFolder, roleName)
	if fileExists(roleFolder) {
		t.Errorf("File was not deleted: %s", roleFolder)
	}
}

//func TestDeleteWithParams(t *testing.T) {
//	for _, param := range params {
//		t.Run(fmt.Sprintf("TestDeleteWithParams--%s", param), func(t *testing.T) {
//			roleName := randomRoleName()
//			deleteParam := fmt.Sprintf("--%s", param)
//			cmd := exec.Command(binary, createCommand, roleName, deleteParam)
//			cmd.Dir = testFolder
//			out, err := cmd.CombinedOutput()
//
//			outStr := string(out)
//			fmt.Println(outStr)
//			if err != nil {
//				t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", err.Error())
//			}
//
//			cmd = exec.Command(binary, deleteCommand, roleName, deleteParam)
//			cmd.Dir = testFolder
//			out, err = cmd.CombinedOutput()
//
//			outStr = string(out)
//			fmt.Println(outStr)
//			if _, matchErr := regexp.MatchString("was deleted successfully", outStr); matchErr != nil {
//				t.Error("'playwright ", cmd, "' was expected to succeed, but it failed with output: ", outStr)
//			}
//
//			tasksFile := fmt.Sprintf("%s/roles/%s/tasks/main.yml", testFolder, roleName)
//			if !fileExists(tasksFile) {
//				t.Errorf("Delete command flag used: [%s]. File was removed: %s", param, tasksFile)
//			}
//
//			var paramFile string
//			if param != templateParam && param != filesParam {
//				paramFile = fmt.Sprintf("%s/roles/%s/%s/main.yml", testFolder, roleName, param)
//				if fileExists(tasksFile) {
//					t.Errorf("Delete command flag used: [%s]. File was not removed: %s", param, paramFile)
//				}
//			}
//			paramFile = fmt.Sprintf("%s/roles/%s/%s", testFolder, roleName, param)
//			if !isDirectory(paramFile) {
//				t.Errorf("Delete command flag used: [%s]. Directory was not removed: %s", param, paramFile)
//			}
//		})
//	}
//}

func createTestProjectStructure() {
	err := os.MkdirAll(testFolder, 0755)
	if err != nil {
		_ = fmt.Errorf("Could not create file %s", configFile)
		return
	}

	var file *os.File
	file, err = os.Create(configFile)
	if err != nil {
		_ = fmt.Errorf("Could not create file %s", configFile)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(config); err != nil {
		_ = fmt.Errorf("Could not write prefix to the file %s", configFile)
		return
	}

	_ = file.Sync()
}

func locateBinary() {
	cd, err := os.Getwd()
	if err != nil {
		_ = fmt.Errorf("Could not find current dir: %s", err)
	}
	binary = fmt.Sprintf("%s/%s", cd, relativeBinaryPath)
}

func removeTestProjectStructure() {
	os.RemoveAll(testFolder)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}

func isDirectory(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func randomRoleName() string {
	return fmt.Sprintf("test-%d", rand.Intn(10000))
}
