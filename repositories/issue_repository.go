package repositories

import (
	"database/sql"
	"fmt"

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

func (issueRepo *IssueRepository) Update(issue *entities.Issue) error {
	query := `UPDATE Issues SET Issue_Name = $1, Issue_Priority = $2, Issue_Status = $3, Issue_Desc = $4, Creator_ID = $5, Project_ID = $6, Issue_Type_ID = $7, Files_Attached = $8 WHERE Issue_ID = $9;`

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

	_, err = stmt.Exec(
		issue.IssueName,
		issue.IssuePriority,
		issue.IssueStatus,
		issue.IssueDesc,
		issue.CreatorID,
		issue.ProjectID,
		issue.IssueTypeID,
		formatted_string_array,
		issue.IssueID,
	)

	fmt.Print(err)

	return err

}

func (issueRepo *IssueRepository) GetAllIssues() ([]entities.Issue, error) {
	query := `SELECT * FROM Issues ORDER BY Issue_Id ASC;`

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

func (issueRepo *IssueRepository) GetIssuesByStatus(status string, projectId int64) ([]entities.Issue, error) {
	query := `SELECT * FROM Issues WHERE Issue_Status = $1 AND Project_Id = $2;`

	rows, err := issueRepo.DB.Query(query, status, projectId)

	if err != nil {
		return nil, err
	}

	var issues []entities.Issue

	for rows.Next() {
		var issue entities.Issue

		var filesAttachedTemp []byte

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
