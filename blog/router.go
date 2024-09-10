package blog

import (
	"github.com/darklab8/blog/blog/about"
	"github.com/darklab8/blog/blog/articles"
	"github.com/darklab8/blog/blog/common/builder"
	"github.com/darklab8/blog/blog/common/urls"
	"github.com/darklab8/blog/blog/examples"
	"github.com/darklab8/blog/blog/favourite"
	"github.com/darklab8/blog/blog/freelancer_menu"
	"github.com/darklab8/blog/blog/pet_projects"
	"github.com/darklab8/blog/blog/pet_projects/ideas"

	"github.com/darklab8/go-utils/goutils/utils/utils_types"
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
			urls.CommunityFreelancer,
			freelancer_menu.FreelancerMenuT(),
		),
		builder.NewComponent(
			"examples.html",
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
	for _, article := range articles.Articles {
		Builder.RegComps(
			builder.NewComponent(
				article.Pagepath,
				articles.ArticleT(article),
			),
		)
	}
}
