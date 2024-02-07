package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"example.com/db/queries"
	"example.com/domain"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		DB: db,
	}
}

func (projectRep *ProjectRepository) Save(project *domain.Project) error {
	query := queries.InsertProject()

	stmt, err := projectRep.DB.Prepare(query)

	if err != nil {
		return err
	}

	project.Created_At = time.Now()
	project.Updated_At = time.Now()

	defer stmt.Close()

	result := stmt.QueryRow(project.ProjectName, project.Project_Desc, project.Project_Category_ID, project.Project_URL, project.Organisation_ID, project.Created_At, project.Updated_At)

	err = result.Scan(&project.ProjectID)

	return err

}

func (projectRep *ProjectRepository) Update(project *domain.Project) error {
	query := queries.UpdateProject()

	stmt, err := projectRep.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	project.Updated_At = time.Now()

	_, err = stmt.Exec(project.ProjectName, project.Project_Desc, project.Project_Category_ID, project.Project_URL, project.Organisation_ID, project.Updated_At, project.ProjectID)

	return err
}

func (projectRep *ProjectRepository) Delete(id int64) error {

	project, err := projectRep.GetProjectById(id)

	if err != nil {
		return err
	}

	query := queries.DeleteProject()

	stmt, err := projectRep.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(project.ProjectID)

	return err

}

func (projectRep *ProjectRepository) GetAllProjects() ([]domain.Project, error) {

	query := queries.GetAllProjects()

	rows, err := projectRep.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var projects []domain.Project

	for rows.Next() {
		var project domain.Project

		err := rows.Scan(
			&project.ProjectID,
			&project.ProjectName,
			&project.Project_Desc,
			&project.Project_Category_ID,
			&project.Project_URL,
			&project.Organisation_ID,
			&project.Created_At,
			&project.Updated_At,
		)

		if err != nil {
			fmt.Print("Error scanning rows: ")
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (projectRep *ProjectRepository) GetProjectById(id int64) (*domain.Project, error) {

	query := queries.GetProjectById()

	row := projectRep.DB.QueryRow(query, id)

	var project domain.Project

	err := row.Scan(
		&project.ProjectID,
		&project.ProjectName,
		&project.Project_Desc,
		&project.Project_Category_ID,
		&project.Project_URL,
		&project.Organisation_ID,
		&project.Created_At,
		&project.Updated_At,
	)

	if err != nil {
		return nil, err
	}

	return &project, nil
}
