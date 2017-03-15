package utils

import "os"

// FolderExists - checks if file exists and is a folder
func FolderExists(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}
