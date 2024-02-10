package blog

import (
	"darklab_blog/blog/about"
	"darklab_blog/blog/articles"
	"darklab_blog/blog/articles/article_detailed"
	"darklab_blog/blog/common/builder"
	"darklab_blog/blog/common/urls"
	"darklab_blog/blog/examples"
	"darklab_blog/blog/favourite"
	"darklab_blog/blog/pet_projects"
	"darklab_blog/blog/pet_projects/ideas"

	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

var Builder *builder.Builder

func init() {
	Builder = builder.NewBuilder()
	Builder.RegComps(
		builder.NewComponent(
			urls.About,
			about.AboutT(),
		),
		builder.NewComponent(
			urls.PetProjects,
			pet_projects.PetProjectsT(),
		),
		builder.NewComponent(
			urls.PetProjectsIdeas,
			ideas.IdeasT(),
		),
		builder.NewComponent(
			urls.Favourite,
			favourite.PageT(),
		),
		builder.NewComponent(
			urls.Articles,
			articles.HomeT(),
		),
		builder.NewComponent(
			"examples",
			examples.Example(),
		),
	)
	for _, idea := range ideas.Ideas {
		Builder.RegComps(
			builder.NewComponent(
				utils_types.FilePath(idea.Url),
				ideas.IdeaTemplate(idea),
			),
		)
	}
	for _, article := range article_detailed.Articles {
		Builder.RegComps(
			builder.NewComponent(
				article.Pagepath,
				article_detailed.ArticleT(article),
			),
		)
	}
}
