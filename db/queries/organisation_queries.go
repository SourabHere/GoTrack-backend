package queries

func GetAllOrganisations() string {
	return `SELECT * FROM Organisations;`
}

func GetOrganisationById() string {
	return `SELECT * FROM Organisations WHERE organisation_id = $1;`
}

func DeleteOrganisation() string {
	return `DELETE FROM Organisations WHERE organisation_id = $1;`
}

func UpdateOrganisation() string {
	return `UPDATE Organisations SET 
	organisation_name = $1, 
	organisation_type = $2, 
	organisation_url = $3, 
	organisation_logo = $4, 
	organisation_location = $5 
	WHERE organisation_id = $6;`
}

func InsertOrganisation() string {
	return `INSERT INTO Organisations 
	(organisation_name, organisation_type, organisation_url, organisation_logo, organisation_location, created_at)
	VALUES ($1,$2,$3,$4,$5,$6)
	RETURNING organisation_id;`
}
