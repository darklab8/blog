package markdowner

import (
	"bytes"
	"darklab_blog/blog/settings/logus"
	"os"
	"text/template"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func TemplateHtml(html string, input any) (string, error) {
	tmpl, err := template.New("test").Parse(html)
	if logus.Log.CheckError(err, "failed to template") {
		return "", err
	}
	var buf bytes.Buffer = bytes.Buffer{}
	err = tmpl.Execute(&buf, input)
	if logus.Log.CheckError(err, "failed to execute templ") {
		return "", err
	}
	result := buf.String()
	return result, nil
}

func ReadMarkdownAndTemplate(path utils_types.FilePath, values any) []byte {
	var err error
	data1, err := os.ReadFile(path.ToString())
	logus.Log.CheckFatal(err, "failed to read git_conv_commits.md")

	data2, err := TemplateHtml(string(data1), values)
	logus.Log.CheckFatal(err, "failed to template git_conv_commits.md")

	data3 := MdToHTML([]byte(data2))

	return data3
}

func MdToHtml(
	markdown_path utils_types.FilePath,
	values any,
) string {
	return string(ReadMarkdownAndTemplate(
		markdown_path,
		values,
	))
}
