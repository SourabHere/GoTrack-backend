package entities

import "time"

type Issue struct {
	IssueID       int       `json:"issueId"`
	IssueName     string    `json:"issueName"`
	IssuePriority string    `json:"issuePriority"`
	IssueStatus   string    `json:"issueStatus"`
	IssueDesc     string    `json:"issueDesc"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DueDate       time.Time `json:"dueDate"`
	CreatorID     int       `json:"creatorId"`
	ProjectID     int       `json:"projectId"`
	IssueTypeID   int       `json:"issueTypeId"`
	FilesAttached []string  `json:"filesAttached"`
}

type IssueRepository interface {
	Save(issue *Issue) error
	// Update(issue *Issue) error
	// Delete(id int64) error
	GetAllIssues() ([]Issue, error)
	GetIssueById(id int64) (*Issue, error)
}

type IssueUsecase interface {
	CreateIssue(issue *Issue) error
	// UpdateIssue(issue *Issue) error
	// DeleteIssue(id int64) error
	GetAllIssues() ([]Issue, error)
	GetIssueById(id int64) (*Issue, error)
}
