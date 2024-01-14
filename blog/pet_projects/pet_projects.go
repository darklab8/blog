package pet_projects

import "darklab_blog/blog/common/compon"

type PetProjects struct {
	compon.Component
}

func NewPetProjects() *PetProjects {
	return &PetProjects{
		compon.NewComponent(
			"pet_projects.html",
			PetProjectsT(),
		),
	}
}
