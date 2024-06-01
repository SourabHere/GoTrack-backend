package repositories

import (
	"database/sql"

	"example.com/domain/entities"
	"example.com/utils"
)

type IssueRepository struct {
	DB *sql.DB
}

func NewIssueRepository(db *sql.DB) *IssueRepository {
	return &IssueRepository{
		DB: db,
	}
}

func (issueRepo *IssueRepository) Save(issue *entities.Issue) error {
	query := `INSERT INTO Issues (Issue_Name,Issue_Desc,Due_Date,Creator_ID,Project_ID,Issue_Type_ID,Files_Attached) 
	VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING Issue_ID;`

	stmt, err := issueRepo.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	incoming_string_array := issue.FilesAttached
	formatted_string_array := ""

	if len(incoming_string_array) > 0 {

		for i := 0; i < len(incoming_string_array); i++ {
			if i == 0 {
				formatted_string_array = "{" + incoming_string_array[i] + ","
			} else {
				formatted_string_array += incoming_string_array[i] + ","
			}
		}

		formatted_string_array = formatted_string_array[:len(formatted_string_array)-1] + "}"

	} else {
		formatted_string_array = "{}"
	}

	result := stmt.QueryRow(
		issue.IssueName,
		issue.IssueDesc,
		issue.DueDate,
		issue.CreatorID,
		issue.ProjectID,
		issue.IssueTypeID,
		formatted_string_array,
	)

	err = result.Scan(&issue.IssueID)

	return err

}

func (issueRepo *IssueRepository) GetIssueById(id int64) (*entities.Issue, error) {
	query := `SELECT * FROM Issues WHERE Issue_ID = $1;`

	stmt := issueRepo.DB.QueryRow(query, id)

	var issue entities.Issue

	var filesAttached []byte

	err := stmt.Scan(
		&issue.IssueID,
		&issue.IssueName,
		&issue.IssuePriority,
		&issue.IssueStatus,
		&issue.IssueDesc,
		&issue.CreatedAt,
		&issue.UpdatedAt,
		&issue.DueDate,
		&issue.CreatorID,
		&issue.ProjectID,
		&issue.IssueTypeID,
		&filesAttached,
	)

	if err != nil {
		return nil, err
	}

	issue.FilesAttached = utils.ParseIssueFilesAttached(filesAttached)

	return &issue, nil

}

func (issueRepo *IssueRepository) GetAllIssues() ([]entities.Issue, error) {
	query := `SELECT * FROM Issues;`

	rows, err := issueRepo.DB.Query(query)

	if err != nil {
		return nil, err
	}

	var issues []entities.Issue

	var filesAttachedTemp []byte

	for rows.Next() {
		var issue entities.Issue
		err := rows.Scan(
			&issue.IssueID,
			&issue.IssueName,
			&issue.IssuePriority,
			&issue.IssueStatus,
			&issue.IssueDesc,
			&issue.CreatedAt,
			&issue.UpdatedAt,
			&issue.DueDate,
			&issue.CreatorID,
			&issue.ProjectID,
			&issue.IssueTypeID,
			&filesAttachedTemp,
		)

		if err != nil {
			return nil, err
		}

		issue.FilesAttached = utils.ParseIssueFilesAttached(filesAttachedTemp)

		issues = append(issues, issue)

	}

	return issues, nil

}
