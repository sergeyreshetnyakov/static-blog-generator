package utils

import (
	"path/filepath"
	"strings"
)

func ReplacePathRootDir(path string, newRoot string) string {
	splited := strings.Split(filepath.Dir(path), string(filepath.Separator))[1:]
	pathWithoutRoot := filepath.Join(splited...)
	pathWithNewRoot := filepath.Join(newRoot, pathWithoutRoot)
	return pathWithNewRoot
}
