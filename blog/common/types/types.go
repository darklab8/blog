package types

import "github.com/darklab8/darklab_goutils/goutils/utils/utils_types"

type Theme int64

const ThemeCtxKey = "theme"

const (
	ThemeDark Theme = iota
	ThemeLight
)

const SiteRootCtxKey = "buildpath"

type GlobalParams struct {
	Buildpath         utils_types.FilePath
	Theme             Theme
	SiteRoot          string
	StaticRoot        string
	OppositeThemeRoot string
}
