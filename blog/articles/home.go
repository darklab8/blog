package articles

import "darklab_blog/blog/common/compon"

/*
Article navigation page
*/

type Home struct {
	compon.Component
}

func NewHome() *Home {
	return &Home{
		compon.NewComponent(
			"index.html",
			HomeT(),
		),
	}
}
