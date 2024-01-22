package article_git_conventional_commits

import (
	"darklab_blog/blog/common/markdowner"
	"darklab_blog/blog/pet_projects/pet_projects_urls"
	"darklab_blog/blog/settings"
	"strings"
	"time"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
)

type Vars struct {
	StaticRoot string
	AutogitURL string
}

func init() {
	Body = markdowner.ReadMarkdownAndTemplate(
		utils_filepath.Join(utils.GetCurrentFolder(), "git_conv_commits.md"),
		Vars{
			StaticRoot: settings.StaticRoot,
			AutogitURL: pet_projects_urls.Autogit,
		},
	)
}

var (
	Headline = "Git Conventional Commits"
	Intro    = strings.ReplaceAll(`We write code not just for machines, but for other humans to read
	(including future us, who forgot the written code in a year).
	Software development is a team effort, and therefore it requires communicating
	what we change and why we change it.`, "\n", "")
	Date = time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC)
	Body []byte
)
