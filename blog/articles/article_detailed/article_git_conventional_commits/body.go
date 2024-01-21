package article_git_conventional_commits

import "strings"

func init() {
}

var (
	Headline = "Git Conventional Commits"
	Intro    = strings.ReplaceAll(`We write code not just for machines, but for other humans to read
	(including future us, who forgot the written code in a year).
	Software development is a team effort, and therefore it requires communicating
	what we change and why we change it.`, "\n", "")
)
