package pet_projects

import (
	"darklab_blog/blog/common/compon"
	"darklab_blog/blog/common/urls"
)

type PetProjects struct {
	compon.Component
}

func NewPetProjects() *PetProjects {
	return &PetProjects{
		compon.NewComponent(
			urls.PetProjects,
			PetProjectsT(),
		),
	}
}
