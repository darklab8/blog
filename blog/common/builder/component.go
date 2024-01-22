package builder

import (
	"bytes"
	"context"
	"darklab_blog/blog/settings"
	"darklab_blog/blog/settings/logus"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
	"github.com/yosssi/gohtml"
)

type Component struct {
	relpath    utils_types.FilePath
	templ_comp templ.Component
}

func NewComponent(
	relpath utils_types.FilePath,
	templ_comp templ.Component,
) *Component {
	return &Component{
		relpath:    relpath,
		templ_comp: templ_comp,
	}
}

func (h *Component) Write(buildpath utils_types.FilePath) {
	buf := bytes.NewBuffer([]byte{})

	h.templ_comp.Render(context.Background(), buf)

	abs_buildpath := utils_filepath.Join(settings.ProjectFolder, buildpath, h.relpath)
	haveParentFoldersCreated(abs_buildpath)

	err := os.WriteFile(abs_buildpath.ToString(), gohtml.FormatBytes(buf.Bytes()), os.ModePerm)
	logus.Log.CheckFatal(err, "failed to write *.html file")
}

func haveParentFoldersCreated(buildpath utils_types.FilePath) {
	path := buildpath.ToString()
	folder_path := filepath.Dir(path)
	err := os.MkdirAll(folder_path, os.ModePerm)
	logus.Log.CheckError(err,
		"haveParentFoldersCreated finished",
		logus_core.Any("folderpath", folder_path),
		logus_core.Any("path", path),
	)
}