package about

import (
	"darklab_blog/blog/common/compon"
	"darklab_blog/blog/common/urls"
)

type About struct {
	compon.Component
}

func NewAbout() *About {
	return &About{
		compon.NewComponent(
			urls.About,
			AboutT(),
		),
	}
}
