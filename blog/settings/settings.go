package settings

import (
	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

var ProjectFolder utils_types.FilePath

func init() {
	ProjectFolder = utils_filepath.Dir(utils_filepath.Dir(utils.GetCurrentFolder()))
}
