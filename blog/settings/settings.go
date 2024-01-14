package settings

import (
	"os"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

var ProjectFolder utils_types.FilePath
var SiteRoot = ""
var StaticRoot = ""

func init() {
	if value, ok := os.LookupEnv("SITE_ROOT"); ok {
		SiteRoot = value
	} else {
		SiteRoot = "/"
	}

	StaticRoot = SiteRoot + "static/"

	ProjectFolder = utils_filepath.Dir(utils_filepath.Dir(utils.GetCurrentFolder()))
}
