package article_static_typed_logging

import (
	"darklab_blog/blog/common/markdowner"
	"darklab_blog/blog/common/types"
	"darklab_blog/blog/pet_projects/pet_projects_urls"
	"strings"
	"time"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
)

type Vars struct {
	StaticRoot string
	GoTypelog  string
	PyTypelog  string
}

func GetBody(gp types.GlobalParams) string {
	return string(markdowner.ReadMarkdownAndTemplate(
		utils_filepath.Join(utils.GetCurrentFolder(), "typelog.md"),
		Vars{
			StaticRoot: gp.StaticRoot,
			GoTypelog:  pet_projects_urls.GoTypelog,
			PyTypelog:  pet_projects_urls.PyTypelog,
		},
	))
}

var (
	Headline = "Typelog - type safe structured logging"
	Intro    = strings.ReplaceAll(`With modern logging systems able to parse JSON out of the box, we need defining easily jsonable logs.
Known solutions do not do it consistently and in a type safe way. Typelog comes to rescue.`, "\n", "")
	Date = time.Date(2024, time.January, 28, 0, 0, 0, 0, time.UTC)
)
