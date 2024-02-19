package usecases

import "example.com/domain/entities"

type ProjectUsecase struct {
	ProjectRepo entities.ProjectRepository
}

func NewProjectUsecase(projectRepo entities.ProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{
		ProjectRepo: projectRepo,
	}
}

func (projectUC *ProjectUsecase) CreateProject(project *entities.Project) error {
	return projectUC.ProjectRepo.Save(project)
}

func (projectUC *ProjectUsecase) UpdateProject(project *entities.Project) error {
	return projectUC.ProjectRepo.Update(project)
}

func (projectUC *ProjectUsecase) DeleteProject(id int64) error {
	return projectUC.ProjectRepo.Delete(id)
}

func (projectUC *ProjectUsecase) GetProjects() ([]entities.Project, error) {
	return projectUC.ProjectRepo.GetAllProjects()
}

func (projectUC *ProjectUsecase) GetProjectById(id int64) (*entities.Project, error) {
	return projectUC.ProjectRepo.GetProjectById(id)
}
