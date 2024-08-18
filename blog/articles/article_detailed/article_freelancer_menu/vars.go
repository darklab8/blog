package article_freelancer_menu

import "github.com/darklab8/blog/blog/pet_projects"

type Vars struct {
	StaticRoot              string
	SiteRoot                string
	ArticleDiscoLinux       string
	ArticleAllShortestPaths string

	AnchorToolDarkstat string
	AnchorToolDarkbot  string
	AnchorToolDarklint string
	AnchorToolConfigs  string

	Darkstat *pet_projects.PetProject
	Darkbot  *pet_projects.PetProject
	Darklint *pet_projects.PetProject
	Configs  *pet_projects.PetProject
}
