package articles

import (
	"darklab_blog/blog/common/compon"
	"darklab_blog/blog/common/urls"
)

/*
Article navigation page
*/

type Home struct {
	compon.Component
}

func NewHome() *Home {
	return &Home{
		compon.NewComponent(
			urls.Home,
			HomeT(),
		),
	}
}
