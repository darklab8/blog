package builder

import (
	"bytes"
	"context"
	"os"
	"path/filepath"

	"github.com/darklab8/blog/blog/common/types"
	"github.com/darklab8/blog/blog/settings"
	"github.com/darklab8/blog/blog/settings/logus"
	"github.com/darklab8/go-typelog/typelog"

	"github.com/a-h/templ"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
	"github.com/yosssi/gohtml"
)

type Component struct {
	pagepath   utils_types.FilePath
	templ_comp templ.Component
}

func NewComponent(
	pagepath utils_types.FilePath,
	templ_comp templ.Component,
) *Component {
	return &Component{
		pagepath:   pagepath,
		templ_comp: templ_comp,
	}
}

func (h *Component) Write(gp types.GlobalParams) {
	buf := bytes.NewBuffer([]byte{})

	gp.Pagepath = string(h.pagepath)

	h.templ_comp.Render(context.WithValue(context.Background(), types.GlobalParamsCtxKey, gp), buf)

	abs_buildpath := utils_filepath.Join(settings.ProjectFolder, gp.Buildpath, h.pagepath)
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
		typelog.String("folderpath", folder_path),
		typelog.String("path", path),
	)
}
