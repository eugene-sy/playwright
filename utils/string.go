package utils

import (
	"bytes"
	"strings"
)

func CheckKey(key string, expected string) bool {
	return strings.Contains(key, expected)
}

func Concat(prefix string, suffix string) string {
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(suffix)
	return buffer.String()
}
