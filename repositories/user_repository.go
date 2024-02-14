package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"example.com/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (userRep *UserRepository) Save(u *domain.User) error {
	query := `INSERT INTO Users (first_name,last_name,email,password,designation_id,user_uuid) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id;;`

	stmt, err := userRep.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result := stmt.QueryRow(
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.Designation_ID,
		u.UserUUID,
	)

	if err != nil {
		return err
	}

	err = result.Scan(&u.UserID)

	return err
}

func (userRep *UserRepository) GetAllUsers() ([]domain.User, error) {

	query := `SELECT * FROM Users;`

	rows, err := userRep.DB.Query(query)

	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, errors.New("no rows returned from the database query")
	}

	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Designation_ID, &user.DateOfJoining, &user.UserUUID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (userRep *UserRepository) GetUserById(uuid string) (*domain.User, error) {

	query := `SELECT * FROM Users WHERE user_uuid = $1;`

	row := userRep.DB.QueryRow(query, uuid)

	var user domain.User

	err := row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Designation_ID,
		&user.DateOfJoining,
		&user.UserUUID,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (userRep *UserRepository) Update(u *domain.User) error {

	query := `UPDATE Users SET first_name = $1, last_name = $2, email = $3, password = $4, designation_id = $5 WHERE user_uuid = $6;`

	stmt, err := userRep.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		u.FirstName,
		u.LastName,
		u.Email,
		u.Password,
		u.Designation_ID,
		u.UserUUID,
	)

	return err
}

func (userRep *UserRepository) Delete(id string) error {

	query := `DELETE FROM Users WHERE user_uuid = $1;`

	stmt, err := userRep.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	fmt.Println(id)

	_, err = stmt.Exec(id)

	return err
}

func (userRep *UserRepository) GetProjectsByUserIdForOrganisation(userId int, organisationId string) ([]domain.Project, error) {

	query := `SELECT p.project_id, p.project_name, p.project_desc, p.project_category_id, p.project_url, p.organisation_id, p.created_at, p.updated_at 
	FROM Projects p INNER JOIN UserProjects as up ON p.project_id = up.project_id 
	WHERE up.user_id = $1 AND p.organisation_id = $2;`

	rows, err := userRep.DB.Query(query, userId, organisationId)

	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, errors.New("no rows returned from the database query")
	}

	defer rows.Close()

	var projects []domain.Project

	for rows.Next() {
		var project domain.Project
		err := rows.Scan(&project.ProjectID,
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
		projects = append(projects, project)
	}

	return projects, nil
}
