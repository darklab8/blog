package blog

import (
	"bytes"
	"context"
	"darklab_blog/blog/logus"
	"darklab_blog/blog/settings"
	"os"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
	"github.com/yosssi/gohtml"
)

/*
Article navigation page
*/

type Home struct{}

func NewHome() *Home { return &Home{} }

const (
	Build utils_types.FilePath = "build"
)

func (h *Home) Write() {
	buf := bytes.NewBuffer([]byte{})

	component := HomeT()
	component.Render(context.Background(), buf)

	err := os.WriteFile(utils_filepath.Join(settings.ProjectFolder, Build, "index.html").ToString(), gohtml.FormatBytes(buf.Bytes()), os.ModePerm)
	logus.Log.CheckError(err, "failed to write index.html")
}
