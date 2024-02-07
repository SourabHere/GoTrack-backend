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

func (userRep *UserRepository) GetUserById(id string) (*domain.User, error) {

	query := `SELECT * FROM Users WHERE user_uuid = $1;`

	row := userRep.DB.QueryRow(query, id)

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
		fmt.Println(id)
		fmt.Println(err)
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
