package article_detailed

import (
	"context"
	"darklab_blog/blog/articles/article_detailed/article_git_conventional_commits"
	"darklab_blog/blog/articles/article_detailed/article_lts_software"
	"darklab_blog/blog/articles/article_detailed/article_static_typed_logging"
	"darklab_blog/blog/common/types"
	"darklab_blog/blog/common/urls"
	"darklab_blog/blog/pet_projects/pet_projects_urls"
	"sort"
	"time"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
)

var Articles []*Article = []*Article{
	NewArticle(
		"Git conventional commits - communicating with git",
		urls.ArticleGitConventionalCommits,
		utils_filepath.Join(utils.GetCurrentFolder(), "article_git_conventional_commits", "git_conv_commits.md"),
		time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
		WithVars(func(ctx context.Context) any {
			return article_git_conventional_commits.Vars{
				StaticRoot: types.GetCtx(ctx).StaticRoot,
				AutogitURL: pet_projects_urls.Autogit,
			}
		}),
		WithDescription(`About usage of Git Conventional Commits, linters and auto changelog generating from your git commits.
		How to communicate easier with your end users through git and releases.`),
	),
	NewArticle(
		"Typelog - type safe structured logging",
		urls.ArticleTypelog,
		utils_filepath.Join(utils.GetCurrentFolder(), "article_static_typed_logging", "typelog.md"),
		time.Date(2024, time.January, 28, 0, 0, 0, 0, time.UTC),
		WithVars(func(ctx context.Context) any {
			return article_static_typed_logging.Vars{
				StaticRoot: types.GetCtx(ctx).StaticRoot,
				GoTypelog:  pet_projects_urls.GoTypelog,
				PyTypelog:  pet_projects_urls.PyTypelog,
			}
		}),
		WithDescription(`
			With modern logging systems able to parse JSON out of the box, we need defining easily jsonable logs.
			Known solutions do not do it consistently and in a type safe way. Typelog comes to rescue.
		`),
	),
	NewArticle(
		"Long term maintained software",
		"article/long_term_maintained_software.html",
		utils_filepath.Join(utils.GetCurrentFolder(), "article_lts_software", "article.md"),
		time.Date(2024, time.February, 9, 0, 0, 0, 0, time.UTC),
		WithDescription(`Some thoughts about how to have long term maintanance software with minimal toll to maintain and keep up to date`),
		WithVars(func(ctx context.Context) any {
			return article_lts_software.Vars{
				StaticRoot: types.GetCtx(ctx).StaticRoot,
			}
		}),
	),
	/*
		TODO articles
			- write: raising ECS cluster (especially powered by EC2)
			- refactor python docker containers article ( https://github.com/darklab8/darklab_article_docker_python )
			- refactor "personal docs as git-crypt"
				- https://github.com/darklab8/darklab_article_documentation_as_a_code
				- https://darklab8.github.io/darklab_article_documentation_as_a_code/
			- refactor "parallel pytest"
				- (needs code refactor. Very Ugly. But small amoutn of code)
				- https://github.com/darklab8/darklab_article_parallel_pytest
			- write: python documentation as a code

			- decide what to do: previous git conventional commits article
				- ❌ https://github.com/darklab8/darklab_article_autogit
				- We have already better one. Perhaps no point to refactor
	*/
}

func init() {
	sort.Slice(Articles[:], func(i, j int) bool {
		return Articles[i].Date.After(Articles[j].Date)
	})
}
