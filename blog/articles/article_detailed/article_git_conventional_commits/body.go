package article_git_conventional_commits

import (
	"darklab_blog/blog/common/compon"
	"darklab_blog/blog/common/urls"
	"strings"
	"time"
)

func init() {
}

var (
	Headline = "Git Conventional Commits"
	Intro    = strings.ReplaceAll(`We write code not just for machines, but for other humans to read
	(including future us, who forgot the written code in a year).
	Software development is a team effort, and therefore it requires communicating
	what we change and why we change it.`, "\n", "")
	Date = time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC)
)

type Article struct {
	compon.Component
}

func NewArticle() *Article {
	return &Article{
		compon.NewComponent(
			urls.ArticleGitConventionalCommits,
			ArticleT(),
		),
	}
}
