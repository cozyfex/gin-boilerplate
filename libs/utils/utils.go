package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func IsFileExist(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else {
		return false
	}
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func GetFileName(file string) string {
	_, f := filepath.Split(file)
	return f[:strings.Index(f, ".")]
}
