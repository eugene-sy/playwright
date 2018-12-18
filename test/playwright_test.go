package test

import (
	"os/exec"
	"regexp"
	"testing"
)

var (
	binary = "../playwright"
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
