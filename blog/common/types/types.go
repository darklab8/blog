package types

import (
	"context"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

type Theme int64

const (
	ThemeDark Theme = iota
	ThemeLight
)

const GlobalParamsCtxKey = "global_params"

type GlobalParams struct {
	Buildpath         utils_types.FilePath
	Theme             Theme
	SiteRoot          string
	StaticRoot        string
	OppositeThemeRoot string
	Pagepath          string
}

func GetCtx(ctx context.Context) GlobalParams {
	return ctx.Value(GlobalParamsCtxKey).(GlobalParams)
}