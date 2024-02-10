package common

import (
	"bytes"
	"context"

	"github.com/a-h/templ"
)

func TemplToStr(templ_comp templ.Component, ctx context.Context) string {
	buf := bytes.NewBuffer([]byte{})

	templ_comp.Render(ctx, buf)

	// gohtml.FormatBytes(buf.Bytes())
	return buf.String()
}
