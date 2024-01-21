package builder

import (
	"darklab_blog/blog/settings"
	"os"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_cp"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_os"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

const BuildpathDefault utils_types.FilePath = "build"

type Builder struct {
	components []*Component
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) RegComps(components ...*Component) {
	b.components = append(b.components, components...)
}

func (b *Builder) build(buildpath utils_types.FilePath) {
	os.RemoveAll(buildpath.ToString())
	os.MkdirAll(buildpath.ToString(), os.ModePerm)

	for _, comp := range b.components {
		comp.Write(buildpath)
	}

	folders := utils_os.GetRecursiveDirs(settings.ProjectFolder)
	for _, folder := range folders {
		if utils_filepath.Base(folder) == "static" {
			utils_cp.Dir(folder.ToString(), utils_filepath.Join(settings.ProjectFolder, utils_types.FilePath(buildpath.ToString()), "static").ToString())
		}
	}
}

func (b *Builder) BuildAll() {
	b.build(BuildpathDefault)
}
