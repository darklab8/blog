package markdowner

import (
	"bytes"
	"darklab_blog/blog/settings/logus"
	"html/template"
	"os"

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

func HtmlFromMdFile(path string) ([]byte, error) {
	mds, err := os.ReadFile(path)
	if logus.Log.CheckError(err, "failed to read file") {
		return nil, err
	}

	md := []byte(mds)
	html := MdToHTML(md)
	return html, nil
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
