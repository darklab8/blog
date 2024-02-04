package article_detailed

import (
	"context"
	"darklab_blog/blog/articles/article_detailed/article_git_conventional_commits"
	"darklab_blog/blog/articles/article_detailed/article_static_typed_logging"
	"darklab_blog/blog/common/types"
	"darklab_blog/blog/common/urls"
	"darklab_blog/blog/pet_projects/pet_projects_urls"
	"sort"
	"strings"
	"time"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
)

var Articles []*Article = []*Article{
	NewArticle(
		"Git conventional commit",
		urls.ArticleGitConventionalCommits,
		utils_filepath.Join(utils.GetCurrentFolder(), "article_git_conventional_commits", "git_conv_commits.md"),
		time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
		WithVars(func(ctx context.Context) any {
			return article_git_conventional_commits.Vars{
				StaticRoot: types.GetCtx(ctx).StaticRoot,
				AutogitURL: pet_projects_urls.Autogit,
			}
		}),
		WithDescription(strings.ReplaceAll(`We write code not just for machines, but for other humans to read
		(including future us, who forgot the written code in a year).
		Software development is a team effort, and therefore it requires communicating
		what we change and why we change it.`, "\n", "")),
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
		WithDescription(strings.ReplaceAll(`With modern logging systems able to parse JSON out of the box, we need defining easily jsonable logs.
		Known solutions do not do it consistently and in a type safe way. Typelog comes to rescue.`, "\n", "")),
	),
}

func init() {
	sort.Slice(Articles[:], func(i, j int) bool {
		return Articles[i].Date.After(Articles[j].Date)
	})
}
