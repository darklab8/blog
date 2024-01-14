package compon

import (
	"bytes"
	"context"
	"darklab_blog/blog/settings"
	"darklab_blog/blog/settings/logus"
	"os"

	"github.com/a-h/templ"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
	"github.com/yosssi/gohtml"
)

type Component struct {
	buildpath  utils_types.FilePath
	templ_comp templ.Component
}

func NewComponent(
	buildpath utils_types.FilePath,
	templ_comp templ.Component,
) Component {
	return Component{
		buildpath:  buildpath,
		templ_comp: templ_comp,
	}
}

const (
	Build utils_types.FilePath = "build"
)

func (h *Component) Write() {
	buf := bytes.NewBuffer([]byte{})

	h.templ_comp.Render(context.Background(), buf)

	err := os.WriteFile(utils_filepath.Join(settings.ProjectFolder, Build, h.buildpath).ToString(), gohtml.FormatBytes(buf.Bytes()), os.ModePerm)
	logus.Log.CheckError(err, "failed to write index.html")
}
