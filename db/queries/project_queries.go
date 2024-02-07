package queries

func GetAllProjects() string {
	return `SELECT * FROM Projects;`
}

func GetProjectById() string {
	return `SELECT * FROM Projects WHERE project_id = $1;`

}

func DeleteProject() string {
	return `DELETE FROM Projects WHERE project_id = $1;`
}

func UpdateProject() string {
	return `UPDATE Projects SET 
	project_name = $1, 
	project_desc = $2, 
	project_category_id = $3, 
	project_url = $4, 
	organisation_id = $5, 
	updated_at = $6 
	WHERE project_id = $7;`
}

func InsertProject() string {
	return `INSERT INTO Projects 
	(project_name, project_desc, project_category_id, project_url, organisation_id, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING project_id;`
}
