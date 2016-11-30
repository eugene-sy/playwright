package utils

import (
	"bufio"
	"bytes"
	"strings"
)

func checkKey(key string, expected string) bool {
	return strings.Contains(key, expected)
}

func concat(prefix string, suffix string) string {
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(suffix)
	return buffer.String()
}
