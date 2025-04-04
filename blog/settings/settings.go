package settings

import (
	"github.com/darklab8/go-utils/utils/enverant"
	"github.com/darklab8/go-utils/utils/utils_filepath"
	"github.com/darklab8/go-utils/utils/utils_os"
	"github.com/darklab8/go-utils/utils/utils_settings"
	"github.com/darklab8/go-utils/utils/utils_types"
)

type BlogEnvVars struct {
	Enver *enverant.Enverant
	utils_settings.UtilsEnvs

	SiteRoot      string
	StaticRoot    string
	StaticPrefix  string
	ProjectFolder utils_types.FilePath
}

var Env BlogEnvVars
var Enverants []*enverant.Enverant

func init() {
	env := enverant.NewEnverant(
		enverant.WithPrefix("BLOG_"),
		enverant.WithDescription("BLOG set of env vars for static site generator"),
	)

	Env = BlogEnvVars{
		Enver:         env,
		UtilsEnvs:     utils_settings.GetEnvs(),
		SiteRoot:      env.GetStr("SITE_ROOT", enverant.OrStr("/"), enverant.WithDesc("useful if wishing serving from github pages sub urls. Makes sure correct link addresses")),
		StaticPrefix:  "static/",
		ProjectFolder: utils_filepath.Dir(utils_filepath.Dir(utils_os.GetCurrentFolder())),
	}
	Env.StaticRoot = Env.SiteRoot + Env.StaticPrefix
	Enverants = append(Enverants, Env.Enver, Env.UtilsEnvs.Enver)
}
