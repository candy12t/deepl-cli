package test

import (
	"os"
	"path/filepath"
)

func ProjectDirPath() string {
	currentDirPath, _ := os.Getwd()
	return filepath.Join(currentDirPath, "..", "..")
}
