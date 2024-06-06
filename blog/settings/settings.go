package settings

import (
	"os"

	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

var StaticPrefix = "static/"

func GetSiteRoot() string {
	var siteRoot string
	if value, ok := os.LookupEnv("SITE_ROOT"); ok {
		siteRoot = value
	} else {
		siteRoot = "/"
	}
	return siteRoot
}

func GetStaticRoot() string {
	return GetSiteRoot() + StaticPrefix
}

var ProjectFolder utils_types.FilePath

func init() {

	ProjectFolder = utils_filepath.Dir(utils_filepath.Dir(utils.GetCurrentFolder()))
}
