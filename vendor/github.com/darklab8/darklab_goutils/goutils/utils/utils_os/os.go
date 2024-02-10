package utils_os

import (
	"os"
	"path/filepath"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

func GetDirsSimply(path utils_types.FilePath) []utils_types.FilePath {
	dirs := []utils_types.FilePath{}
	files, err := os.ReadDir(string(path))

	if err != nil {
		panic(err)
	}

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			dirs = append(dirs, utils_types.FilePath(filepath.Join(string(path), fileInfo.Name())))
		}
	}

	return dirs
}

func GetRecursiveDirs(path utils_types.FilePath) []utils_types.FilePath {
	dirs := GetDirsSimply(path)
	copied_dirs := []utils_types.FilePath{}
	copied_dirs = append(copied_dirs, dirs...)
	for _, dir := range copied_dirs {
		nested_dirs := GetRecursiveDirs(dir)
		dirs = append(dirs, nested_dirs...)
	}
	return dirs
}
