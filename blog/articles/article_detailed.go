package articles

import (
	"context"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type EmptyVars struct{}

type TitlePicture struct {
	Path utils_types.FilePath
}

type Article struct {
	Title        string
	TitlePicture *TitlePicture
	Pagepath     utils_types.FilePath
	Description  string
	Date         time.Time
	extra        templ.Component
	vars         func(ctx context.Context) any
	md_path      utils_types.FilePath
}

func NewArticle(
	Title string,
	Pagepath utils_types.FilePath,
	md_path utils_types.FilePath,
	Date time.Time,
	opts ...ArticleOption,
) *Article {
	a := &Article{
		Title:    Title,
		md_path:  md_path,
		Pagepath: Pagepath,
		Date:     Date,
		vars:     func(ctx context.Context) any { return EmptyVars{} },
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

type ArticleOption func(b *Article)

func WithVars(vars func(ctx context.Context) any) ArticleOption {
	return func(b *Article) {
		b.vars = vars
	}
}

func WithDescription(Description string) ArticleOption {
	return func(b *Article) {
		b.Description = strings.ReplaceAll(Description, "\n", "")
	}
}

func WithTitlePicture(picture TitlePicture) ArticleOption {
	return func(b *Article) {
		b.TitlePicture = &picture
	}
}
