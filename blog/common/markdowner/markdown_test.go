package markdowner

import (
	"darklab_blog/blog/settings"
	"darklab_blog/blog/settings/logus"
	"os"
	"testing"
	"text/template"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
)

func TestMarkdown(t *testing.T) {
	html, err := HtmlFromMdFile(string(utils_filepath.Join(settings.ProjectFolder,
		"blog/articles/article_detailed/article_git_conventional_commits/git_conv_commits.md",
	)))
	logus.Log.CheckFatal(err, "failed to read")

	os.WriteFile("check.html", html, os.ModePerm)
}

func TestInject(t *testing.T) {
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
