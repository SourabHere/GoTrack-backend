package entities

import (
	"time"
)

type Project struct {
	ProjectID           int       `json:"projectid"`
	ProjectName         string    `json:"projectName"`
	Project_Desc        string    `json:"projectDesc"`
	Project_Category_ID int       `json:"projectCategoryID"`
	Project_URL         *string   `json:"projectURL"`
	Organisation_ID     int       `json:"organisationID"`
	Created_At          time.Time `json:"createdAt"`
	Updated_At          time.Time `json:"updatedAt"`
}

type ProjectRepository interface {
	Save(project *Project) (*Project, error)
	Update(project *Project) error
	Delete(id int64) error
	GetAllProjects() ([]Project, error)
	GetProjectById(id int64) (*Project, error)
	GetProjectCategories() ([]string, error)
	GetProjectCategoryIDByName(name string) (int, error)
}

type ProjectUsecase interface {
	CreateProject(project *Project) (*Project, error)
	UpdateProject(project *Project) error
	DeleteProject(id int64) error
	GetProjects() ([]Project, error)
	GetProjectById(id int64) (*Project, error)
	GetProjectCategories() ([]string, error)
	GetProjectCategoryIDByName(name string) (int, error)
}
