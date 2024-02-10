package utils_filepath

import (
	"path/filepath"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

func Join(paths ...utils_types.FilePath) utils_types.FilePath {
	return utils_types.FilePath(filepath.Join(utils.CompL(paths, func(path utils_types.FilePath) string { return string(path) })...))
}

func Dir(path utils_types.FilePath) utils_types.FilePath {
	return utils_types.FilePath(filepath.Dir(string(path)))
}

func Base(path utils_types.FilePath) utils_types.FilePath {
	return utils_types.FilePath(filepath.Base(string(path)))
}
