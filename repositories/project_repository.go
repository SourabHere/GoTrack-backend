package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"example.com/db/queries"
	"example.com/domain/entities"
)

type ProjectRepository struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		DB: db,
	}
}

func (projectRep *ProjectRepository) Save(project *entities.Project) (*entities.Project, error) {
	query := queries.InsertProject()

	stmt, err := projectRep.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	project.Created_At = time.Now()
	project.Updated_At = time.Now()

	defer stmt.Close()

	result := stmt.QueryRow(project.ProjectName, project.Project_Desc, project.Project_Category_ID, project.Project_URL, project.Organisation_ID, project.Created_At, project.Updated_At)

	err = result.Scan(&project.ProjectID)

	if err != nil {
		return nil, err
	}

	return project, nil

}

func (projectRep *ProjectRepository) Update(project *entities.Project) error {
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

func (projectRep *ProjectRepository) GetAllProjects() ([]entities.Project, error) {

	query := queries.GetAllProjects()

	rows, err := projectRep.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var projects []entities.Project

	for rows.Next() {
		var project entities.Project

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

func (projectRep *ProjectRepository) GetProjectById(id int64) (*entities.Project, error) {

	query := queries.GetProjectById()

	row := projectRep.DB.QueryRow(query, id)

	var project entities.Project

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

func (projectRep *ProjectRepository) GetProjectCategories() ([]string, error) {
	query := `SELECT category_name FROM Category;`

	rows, err := projectRep.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var projectTypes []string

	for rows.Next() {
		var projectType string

		err := rows.Scan(&projectType)

		if err != nil {
			return nil, err
		}

		projectTypes = append(projectTypes, projectType)
	}

	return projectTypes, nil

}

func (projectRep *ProjectRepository) GetProjectCategoryIDByName(name string) (int, error) {
	query := `SELECT category_id FROM Category WHERE category_name = $1;`

	var id int

	err := projectRep.DB.QueryRow(query, name).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		insertQuery := `INSERT INTO Category (category_name) VALUES ($1) RETURNING category_id;`

		err = projectRep.DB.QueryRow(insertQuery, name).Scan(&id)
		if err != nil {
			return 0, err
		}
	case err != nil:
		return 0, err
	}

	return id, nil
}
