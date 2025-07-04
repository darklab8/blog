package articles

import (
	"context"
	"sort"
	"time"

	"github.com/darklab8/blog/blog/archive"
	"github.com/darklab8/blog/blog/articles/article_detailed/article_20231211_git_conventional_commits"
	"github.com/darklab8/blog/blog/articles/article_detailed/article_20240128_static_typed_logging"
	"github.com/darklab8/blog/blog/articles/article_detailed/article_20240228_lts_software"
	"github.com/darklab8/blog/blog/articles/article_detailed/article_20240614_freelancer_setup_at_linux"
	"github.com/darklab8/blog/blog/articles/article_detailed/article_20240619_shortest_paths"
	"github.com/darklab8/blog/blog/articles/article_detailed/article_20240908_choosing_pet_projects"
	"github.com/darklab8/blog/blog/articles/article_detailed/article_20250506_visual_debugger_in_vscode"
	"github.com/darklab8/blog/blog/articles/article_detailed/article_20250609_grafana"

	"github.com/darklab8/blog/blog/common"
	"github.com/darklab8/blog/blog/common/types"
	"github.com/darklab8/blog/blog/common/urls"
	"github.com/darklab8/blog/blog/pet_projects/pet_projects_urls"

	_ "embed"

	"github.com/darklab8/go-utils/utils/utils_filepath"
	"github.com/darklab8/go-utils/utils/utils_os"
)

var artcieles_root = utils_filepath.Join(utils_os.GetCurrentFolder(), "article_detailed")

//go:embed article_detailed/article_20240619_shortest_paths/trades/floyd.go
var floyd_main_code string

//go:embed article_detailed/article_20240619_shortest_paths/trades/floyd_test.go
var floyd_test_code string

//go:embed article_detailed/article_20240619_shortest_paths/trades/heap.go
var shortest_paths_heap string

//go:embed article_detailed/article_20240619_shortest_paths/trades/johnson.go
var shortest_paths_johnson_code string

//go:embed article_detailed/article_20240619_shortest_paths/trades/johnson_test.go
var shortest_paths_Johnson_test string

//go:embed article_detailed/article_20240619_shortest_paths/trades/dijkstra_apsp.go
var dijkstra_apsp_code string

//go:embed article_detailed/article_20240619_shortest_paths/trades/dijkstra_apsp_test.go
var dijkstra_apsp_test string

//go:embed article_detailed/article_20240619_shortest_paths/trades/graph.go
var graph_shared_code string

var ArticleDiscoLinux *Article = NewArticle(
	"Freelancer Discovery setup with Lutris, Wine at Linux",
	"article/article_freelancer_setup_at_linux.html",
	utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "disco", "article_start_disco.md"),
	time.Date(2024, time.July, 14, 20, 0, 0, 0, time.UTC),
	WithDescription(`Using Lutris, custom Wine, Wine Tricks to install all custom dependencies
	for launching Freelancer Disovery space simulator at Linux`),
	WithVars(func(ctx context.Context) any {
		return article_20240614_freelancer_setup_at_linux.Vars{
			StaticRoot: types.GetCtx(ctx).StaticRoot,
			SiteRoot:   types.GetCtx(ctx).SiteRoot,
		}
	}),
	WithTitlePicture(TitlePicture{
		Path: utils_filepath.Join("article_20240614_freelancer_setup_at_linux", "installer_picture.png"),
	}),
	WithMoreMarkdowns(
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_intro.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_dependencies.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_lutris.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_wine.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_freelancer.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_winetricks.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_dxvk.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "disco", "article_setup_discovery.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_post_installation.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_dll_override.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "disco", "article_post_installation_disco.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "disco", "article_launch_discovery.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_extra_info_d3d8.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_extra_info_shared.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_extra_info_mirate_proton_8.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_acknowledgements.md"),
	),
	WitHidden(),
)

var ArticleFreelancerVanillaLinux *Article = NewArticle(
	"Freelancer Vanilla setup with Lutris, Wine at Linux",
	"article/article_freelancer_vanilla_at_linux.html",
	utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "vanilla", "article_start_vanilla.md"),
	time.Date(2024, time.August, 19, 20, 0, 0, 0, time.UTC),
	WithDescription(`Using Lutris, custom Wine, Wine Tricks to install all custom dependencies
	for launching Freelancer Vanilla space simulator at Linux`),
	WithVars(func(ctx context.Context) any {
		return article_20240614_freelancer_setup_at_linux.Vars{
			StaticRoot: types.GetCtx(ctx).StaticRoot,
			SiteRoot:   types.GetCtx(ctx).SiteRoot,
		}
	}),
	WithTitlePicture(TitlePicture{
		Path: utils_filepath.Join("freelancer_vanilla", "install01.jpg"),
	}),
	WithMoreMarkdowns(
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_intro.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_dependencies.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_lutris.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_wine.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_freelancer.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_winetricks.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "vanilla", "article_setup_combopatch.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_post_installation.md"),

		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_dxvk.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_dll_override.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "vanilla", "article_launch_vanilla.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_extra_info_d3d8.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_extra_info_shared.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_acknowledgements.md"),
	),
	WitHidden(),
)

var ArticleFreelancerHDLinux *Article = NewArticle(
	"Freelancer HD Edition setup with Lutris, Wine at Linux",
	"article/article_freelancer_hd_edition_at_linux.html",
	utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "hd_edition", "article_start_hd.md"),
	time.Date(2024, time.August, 19, 21, 0, 0, 0, time.UTC),
	WithDescription(`Using Lutris, custom Wine, Wine Tricks to install all custom dependencies
	for launching Freelancer HD Edition space simulator at Linux`),
	WithVars(func(ctx context.Context) any {
		return article_20240614_freelancer_setup_at_linux.Vars{
			StaticRoot: types.GetCtx(ctx).StaticRoot,
			SiteRoot:   types.GetCtx(ctx).SiteRoot,
		}
	}),
	WithTitlePicture(TitlePicture{
		Path: utils_filepath.Join("freelancer_hd_edition", "installer.png"),
	}),
	WithMoreMarkdowns(
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_intro.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_dependencies.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_lutris.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_wine.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_freelancer.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_setup_winetricks.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "hd_edition", "article_setup_hd_edition.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_post_installation.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "hd_edition", "article_launch_hd.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_extra_info_shared.md"),
		utils_filepath.Join(artcieles_root, "article_20240614_freelancer_setup_at_linux", "article_acknowledgements.md"),
	),
	WitHidden(),
)

var ArticleAllShortestPaths = NewArticle(
	"All Shortest Paths in Graph with Golang",
	"article/article_shortest_paths.html",
	utils_filepath.Join(artcieles_root, "article_20240619_shortest_paths", "article.md"),
	time.Date(2024, time.June, 19, 20, 0, 0, 0, time.UTC),
	WithDescription(`Finding All Shortest Paths with Floyd and Johnson in Golang.
		Comparison, profiling and optimization with parallelization.
		Calculating distances for trading routes in a space simulator game.`),
	WithVars(func(ctx context.Context) any {
		return article_20240619_shortest_paths.Vars{
			StaticRoot:       types.GetCtx(ctx).StaticRoot,
			DarkstatUrl:      types.GetCtx(ctx).SiteRoot + "pet_projects.html#fl-darkstat",
			FloydMain:        floyd_main_code,
			FloydTest:        floyd_test_code,
			HeapCode:         shortest_paths_heap,
			JohnsonCode:      shortest_paths_johnson_code,
			JohnsonTest:      shortest_paths_Johnson_test,
			DijkstraApsp:     dijkstra_apsp_code,
			DijkstraApspTest: dijkstra_apsp_test,
			GraphCode:        graph_shared_code,
		}
	}),
	WithTitlePicture(TitlePicture{
		Path: utils_filepath.Join("shortest_paths", "constellations.jpg"),
	}),
	WithOgImage(&common.OgImage{
		Url: utils_filepath.Join("shortest_paths", "og_image.jpg").ToString(),
	}),
)

var Articles []*Article = []*Article{
	NewArticle(
		"Git conventional commits - communicating with git",
		urls.ArticleGitConventionalCommits,
		utils_filepath.Join(artcieles_root, "article_20231211_git_conventional_commits", "git_conv_commits.md"),
		time.Date(2023, time.December, 11, 0, 0, 0, 0, time.UTC),
		WithVars(func(ctx context.Context) any {
			return article_20231211_git_conventional_commits.Vars{
				StaticRoot: types.GetCtx(ctx).StaticRoot,
				AutogitURL: pet_projects_urls.Autogit,
			}
		}),
		WithDescription(`About usage of Git Conventional Commits, linters and auto changelog generating from your git commits.
		How to communicate easier with your end users through git and releases.`),
		WithTitlePicture(TitlePicture{
			Path: utils_filepath.Join("article_commits", "autogit_title_pic.jpg"),
		}),
	),
	NewArticle(
		"Typelog - type safe structured logging",
		urls.ArticleTypelog,
		utils_filepath.Join(artcieles_root, "article_20240128_static_typed_logging", "typelog.md"),
		time.Date(2024, time.January, 28, 0, 0, 0, 0, time.UTC),
		WithVars(func(ctx context.Context) any {
			return article_20240128_static_typed_logging.Vars{
				StaticRoot:     types.GetCtx(ctx).StaticRoot,
				GoTypelog:      pet_projects_urls.GoTypelog,
				PyTypelog:      pet_projects_urls.PyTypelog,
				LinkTypeSafety: common.TemplToStr(archive.LinkT(archive.LinkTypeSafety, "what is type safety, check here"), ctx),
			}
		}),
		WithDescription(`
			With modern logging systems able to parse JSON out of the box, we need defining easily jsonable logs.
			Known solutions do not do it consistently and in a type safe way. Typelog comes to rescue.
		`),
		WithTitlePicture(TitlePicture{
			Path: utils_filepath.Join("typelog", "typelog_title_pic.jpg"),
		}),
	),
	NewArticle(
		"Long term maintained software",
		"article/long_term_maintained_software.html",
		utils_filepath.Join(artcieles_root, "article_20240228_lts_software", "article.md"),
		time.Date(2024, time.February, 9, 0, 0, 0, 0, time.UTC),
		WithDescription(`Some thoughts about how to have long term maintanance software with minimal toll to maintain and keep up to date`),
		WithVars(func(ctx context.Context) any {
			return article_20240228_lts_software.Vars{
				StaticRoot:              types.GetCtx(ctx).StaticRoot,
				LinkColdBloodedSoftware: common.TemplToStr(archive.LinkT(archive.LinkTypeSafety, "cold blooded software"), ctx),
			}
		}),
		WithTitlePicture(TitlePicture{
			Path: utils_filepath.Join("cold_blood", "cold_blood_title_pic.jpg"),
		}),
	),
	ArticleAllShortestPaths,
	ArticleDiscoLinux,
	ArticleFreelancerVanillaLinux,
	ArticleFreelancerHDLinux,
	NewArticle(
		"Choosing software engineering pet projects",
		"choosing_pet_projects.html",
		utils_filepath.Join(artcieles_root, "article_20240908_choosing_pet_projects", "article.md"),
		time.Date(2024, time.September, 8, 20, 0, 0, 0, time.UTC),
		WithDescription(`Some words on how to pick long term pet projects for a programmer`),
		WithVars(func(ctx context.Context) any {
			return article_20240908_choosing_pet_projects.Vars{
				StaticRoot: types.GetCtx(ctx).StaticRoot,
				SiteRoot:   types.GetCtx(ctx).SiteRoot,
			}
		}),
		WithTitlePicture(TitlePicture{
			Path:           utils_filepath.Join("choosing_pet_projects", "logo.png"),
			Attribution:    "picture by Kelly Sikkema",
			AttributionUrl: "https://unsplash.com/photos/white-paper-on-black-table-4TBOXap8qg4",
		}),
	),
	NewArticle(
		"Vscode debugger recipes for python and docker",
		"article_visual_debugger_in_vscode.html",
		utils_filepath.Join(artcieles_root, "article_20250506_visual_debugger_in_vscode", "debugger.md"),
		time.Date(2025, time.April, 6, 20, 0, 0, 0, time.UTC),
		WithDescription(`How to start working with visual debuger in vscode for python.
Including with connecting to already running docker. Written for people wishing to be quickly productive with vscode and for known acquintances wishing to migrate from pycharm :]`),
		WithVars(func(ctx context.Context) any {
			return article_20250506_visual_debugger_in_vscode.Vars{
				StaticRoot: types.GetCtx(ctx).StaticRoot,
				SiteRoot:   types.GetCtx(ctx).SiteRoot,
			}
		}),
		WithTitlePicture(TitlePicture{
			Path: utils_filepath.Join("visual_debugger_in_vscode", "logo.png"),
		}),
	),
	NewArticle(
		"Grafana monitoring with Docker. Part 1 - Logs with Loki",
		"article_grafana_loki.html",
		utils_filepath.Join(artcieles_root, "article_20250609_grafana", "grafana_part1_loki.md"),
		time.Date(2025, time.June, 9, 10, 0, 0, 0, time.UTC),
		WithDescription(`Configuring Grafana monitoring for homelab or small load companies for Docker.
Configuring with Opentofu(Terraform) or Docker-compose. For any backend language capable to emit logs in JSON.`),
		WithVars(func(ctx context.Context) any {
			return article_20250609_grafana.Vars{
				StaticRoot: types.GetCtx(ctx).StaticRoot,
				SiteRoot:   types.GetCtx(ctx).SiteRoot,

				MainCompose:   article_20250609_grafana.MainCompose,
				MainTerraform: article_20250609_grafana.MainTerraform,

				GrafanaConfig:     article_20250609_grafana.GrafanaConfig,
				GrafanaDockerfile: article_20250609_grafana.GrafanaDockerfile,

				LokiConfig:     article_20250609_grafana.LokiConfig,
				LokiDockerfile: article_20250609_grafana.LokiDockerfile,

				AlloyLogsConfig:     article_20250609_grafana.AlloyLogsConfig,
				AlloyLogsDockerfile: article_20250609_grafana.AlloyLogsDockerfile,

				AppLogsDashboard: article_20250609_grafana.AppLogsDashboard,

				HetznerFirewall: article_20250609_grafana.HetznerFirewall,
			}
		}),
		WithTitlePicture(TitlePicture{
			Path: utils_filepath.Join("grafana_loki", "loki_drilldown1.png"),
		}),
	),
	/*
		TODO articles
			- write article about refactoring legacy code based on your AWS Step functions experience?
			- write: raising ECS cluster (especially powered by EC2)
			- write article about making pet projects
				- https://discord.com/channels/838802001861017662/842969729336344614/1204779829258354739
				- https://discord.com/channels/267624335836053506/470889390588035082/1204454719435710554
			- refactor python docker containers article ( https://github.com/darklab8/darklab_article_docker_python )
			- refactor "personal docs as git-crypt"
				- https://github.com/darklab8/darklab_article_documentation_as_a_code
				- https://darklab8.github.io/darklab_article_documentation_as_a_code/
			- refactor "parallel pytest"
				- (needs code refactor. Very Ugly. But small amoutn of code)
				- https://github.com/darklab8/darklab_article_parallel_pytest
			- write: python documentation as a code
			- write article about cached dockerized CI:
				1) Calculate md5 hash from dependencies file lock onto libraries. Whatever your language is. Lets call the value as `builder_base_hash`
				2) Pull image image `{{ builder_base_hash }}` if it exists in docker registry, if not then build up to stage --builder. Save result to docker registry under tag ` {{ builder_base_hash }}`
				(Thus we implemented speed up twice for CI), as we are were able to cache half of longest CI in a way that its CI jobs can run at different runners, as we use remote for persistence
				3) run full building of an image, til the code capable to run unit tests and push to docker registry under tag ` build_${{ github.run_id }}`
				4) at unit test stage: pull the image ` build_${{ github.run_id }}` and run unit tests and other tests
				5) if it passed them, than save the image as `service_name_{{ github.run_number }}` as fit for deployment :slight_smile:, also mark it as `latest` and etc whatever tags u need
				Optionally build Github Actions. For md5 calculations. For golang, For Dagger CI? using Cue-lang? Experiment what to use.
			- Write an article regarding AWS ECS/EKS terraform/docker image deployment procedure
			- Write an article about importantce of unit testing
	*/
}

func init() {
	sort.Slice(Articles[:], func(i, j int) bool {
		return Articles[i].Date.After(Articles[j].Date)
	})
}
