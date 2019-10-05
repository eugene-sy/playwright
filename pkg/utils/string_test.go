package utils

import "testing"

func TestConcat(t *testing.T) {
	prefix := "prefix"
	suffix := "suffix"

	result := Concat(prefix, suffix)
	if result != "prefixsuffix" {
		t.Error("Concat of ", prefix, " and ", suffix, "should be prefixsuffix, but is ", result)
	}
}
