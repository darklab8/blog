package articles

import (
	"darklab_blog/blog/common/compon"
	"darklab_blog/blog/common/urls"
)

/*
Article navigation page
*/

type Articles struct {
	compon.Component
}

func NewArticles() *Articles {
	return &Articles{
		compon.NewComponent(
			urls.Articles,
			HomeT(),
		),
	}
}
