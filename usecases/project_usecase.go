package usecases

import "example.com/domain"

type ProjectUsecase struct {
	ProjectRepo domain.ProjectRepository
}

func NewProjectUsecase(projectRepo domain.ProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{
		ProjectRepo: projectRepo,
	}
}

func (projectUC *ProjectUsecase) CreateProject(project *domain.Project) error {
	return projectUC.ProjectRepo.Save(project)
}

func (projectUC *ProjectUsecase) UpdateProject(project *domain.Project) error {
	return projectUC.ProjectRepo.Update(project)
}

func (projectUC *ProjectUsecase) DeleteProject(id int64) error {
	return projectUC.ProjectRepo.Delete(id)
}

func (projectUC *ProjectUsecase) GetProjects() ([]domain.Project, error) {
	return projectUC.ProjectRepo.GetAllProjects()
}

func (projectUC *ProjectUsecase) GetProjectById(id int64) (*domain.Project, error) {
	return projectUC.ProjectRepo.GetProjectById(id)
}
