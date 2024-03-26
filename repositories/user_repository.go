package repositories

import (
	"database/sql"
	"errors"

	"example.com/domain/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (userRep *UserRepository) Save(u *entities.User) error {
	query := `INSERT INTO Users (first_name,last_name,email,designation_id,user_uuid) VALUES ($1, $2, $3, $4, $5) RETURNING user_id;`

	stmt, err := userRep.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result := stmt.QueryRow(
		u.FirstName,
		u.LastName,
		u.Email,
		u.Designation_ID,
		u.UserUUID,
	)

	err = result.Scan(&u.UserID)

	return err
}

func (userRep *UserRepository) GetAllUsers() ([]entities.User, error) {

	query := `SELECT * FROM Users;`

	rows, err := userRep.DB.Query(query)

	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, errors.New("no rows returned from the database query")
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Designation_ID, &user.DateOfJoining, &user.UserUUID, &user.Location)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (userRep *UserRepository) GetUserByUUID(uuid string) (*entities.User, error) {

	query := `SELECT * FROM Users WHERE user_uuid = $1;`

	row := userRep.DB.QueryRow(query, uuid)

	var user entities.User

	err := row.Scan(
		&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Designation_ID,
		&user.DateOfJoining,
		&user.UserUUID,
		&user.Location,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (userRep *UserRepository) Update(u *entities.User) error {

	query := `UPDATE Users SET first_name = $1, last_name = $2, email = $3, designation_id = $4 WHERE user_uuid = $5;`

	stmt, err := userRep.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		u.FirstName,
		u.LastName,
		u.Email,
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

	_, err = stmt.Exec(id)

	return err
}

func (userRep *UserRepository) GetProjectsByUserIdForOrganisation(userId int, organisationId string) ([]entities.Project, error) {

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

	var projects []entities.Project

	for rows.Next() {
		var project entities.Project
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

func (userUC *UserRepository) GetUserOrganisationByUUID(userUUID string) ([]entities.Organisation, error) {
	query := `SELECT * FROM Organisations WHERE organisation_id IN (SELECT organisation_id FROM UserOrganisations WHERE user_id = (SELECT user_id FROM Users WHERE user_uuid = $1));`

	rows, err := userUC.DB.Query(query, userUUID)

	if err != nil {
		return nil, err
	}

	if rows == nil {
		return nil, errors.New("no rows returned from the database query")
	}

	defer rows.Close()

	var organisations []entities.Organisation

	for rows.Next() {
		var organisation entities.Organisation
		err := rows.Scan(&organisation.Organisation_ID,
			&organisation.Organisation_Name,
			&organisation.Organisation_Type,
			&organisation.Organisation_URL,
			&organisation.Organisation_Logo,
			&organisation.Organisation_Location,
			&organisation.Created_At,
		)

		if err != nil {
			return nil, err
		}

		organisations = append(organisations, organisation)
	}

	return organisations, nil
}

func (userUC *UserRepository) GetUserIdByUUID(userUUID string) (int, error) {
	query := `SELECT user_id FROM Users WHERE user_uuid = $1;`

	row := userUC.DB.QueryRow(query, userUUID)

	var userID int

	err := row.Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (userUC *UserRepository) GetDesignationByID(DesignationID int64) (string, error) {
	query := `SELECT designation_Name FROM UserDesignation WHERE UD_ID = $1;`

	row := userUC.DB.QueryRow(query, DesignationID)

	var designation_name string

	err := row.Scan(&designation_name)

	if err != nil {
		return "-1", err
	}

	return designation_name, nil
}
