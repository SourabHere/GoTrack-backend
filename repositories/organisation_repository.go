package repositories

import (
	"database/sql"
	"time"

	"example.com/db/queries"
	"example.com/domain/entities"
)

type OrganisationRepository struct {
	DB *sql.DB
}

func NewOrganisationRepository(db *sql.DB) *OrganisationRepository {
	return &OrganisationRepository{
		DB: db,
	}
}

func (organisationRep *OrganisationRepository) Save(org *entities.Organisation) error {
	query := queries.InsertOrganisation()

	stmt, err := organisationRep.DB.Prepare(query)

	if err != nil {
		panic(err.Error())
	}

	org.Created_At = time.Now()

	defer stmt.Close()

	result := stmt.QueryRow(
		org.Organisation_Name,
		org.Organisation_Type,
		org.Organisation_URL,
		org.Organisation_Logo,
		org.Organisation_Location,
		org.Created_At,
	)

	err = result.Scan(&org.Organisation_ID)

	return err

}

func (organisationRep *OrganisationRepository) Update(org *entities.Organisation) error {
	query := queries.UpdateOrganisation()

	stmt, err := organisationRep.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(org.Organisation_Name, org.Organisation_Type, org.Organisation_URL, org.Organisation_Logo, org.Organisation_Location, org.Organisation_ID)

	return err

}

func (organisationRep *OrganisationRepository) Delete(id int64) error {

	org, err := organisationRep.GetOrganisationById(id)

	if err != nil {
		return err
	}

	query := queries.DeleteOrganisation()

	stmt, err := organisationRep.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(org.Organisation_ID)

	return err

}

func (organisationRep *OrganisationRepository) GetOrganisationById(id int64) (*entities.Organisation, error) {
	query := queries.GetOrganisationById()

	row := organisationRep.DB.QueryRow(query, id)

	var organisation entities.Organisation

	err := row.Scan(&organisation.Organisation_ID, &organisation.Organisation_Name, &organisation.Organisation_Type, &organisation.Organisation_URL, &organisation.Organisation_Logo, &organisation.Organisation_Location, &organisation.Created_At)

	if err != nil {
		return nil, err
	}

	return &organisation, nil

}

func (organisationRep *OrganisationRepository) GetAllOrganisations() ([]entities.Organisation, error) {
	query := queries.GetAllOrganisations()

	rows, err := organisationRep.DB.Query(query)

	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var organisations []entities.Organisation

	for rows.Next() {
		var organisation entities.Organisation

		err := rows.Scan(&organisation.Organisation_ID, &organisation.Organisation_Name, &organisation.Organisation_Type, &organisation.Organisation_URL, &organisation.Organisation_Logo, &organisation.Organisation_Location, &organisation.Created_At)

		if err != nil {
			panic(err.Error())
		}

		organisations = append(organisations, organisation)

	}

	return organisations, nil

}
