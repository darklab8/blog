package blog

import (
	"darklab_blog/blog/about"
	"darklab_blog/blog/articles"
	"darklab_blog/blog/articles/article_detailed/article_git_conventional_commits"
	"darklab_blog/blog/common/builder"
	"darklab_blog/blog/common/urls"
	"darklab_blog/blog/pet_projects"
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
			urls.Articles,
			articles.HomeT(),
		),
		builder.NewComponent(
			urls.ArticleGitConventionalCommits,
			article_git_conventional_commits.ArticleT(),
		),
	)
}