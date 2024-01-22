package article_git_conventional_commits

import (
	"darklab_blog/blog/common/markdowner"
	"darklab_blog/blog/pet_projects/pet_projects_urls"
	"darklab_blog/blog/settings"
	"darklab_blog/blog/settings/logus"
	"os"
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
	var err error
	data1, err := os.ReadFile(string(utils_filepath.Join(utils.GetCurrentFolder(), "git_conv_commits.md").ToString()))
	logus.Log.CheckFatal(err, "failed to read git_conv_commits.md")

	data2, err := markdowner.TemplateHtml(string(data1), Vars{
		StaticRoot: settings.StaticRoot,
		AutogitURL: pet_projects_urls.Autogit,
	})
	logus.Log.CheckFatal(err, "failed to template git_conv_commits.md")

	data3 := markdowner.MdToHTML([]byte(data2))

	Body = data3
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
