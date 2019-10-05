package utils

import (
	"bytes"
)

// Concat - concatenates prefix and suffix strings, is used to build paths to files/roles
func Concat(prefix string, suffix string) string {
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(suffix)
	return buffer.String()
}
