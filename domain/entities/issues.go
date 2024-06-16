package entities

import "time"

type Issue struct {
	IssueID       int64     `json:"issueId"`
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

type IssuesMap struct {
	ToDo       []Issue `json:"To Do"`
	Open       []Issue `json:"Open"`
	InProgress []Issue `json:"In Progress"`
	Review     []Issue `json:"Review"`
	Closed     []Issue `json:"Closed"`
}

type IssueRepository interface {
	Save(issue *Issue) error
	Update(issue *Issue) error
	// Delete(id int64) error
	GetAllIssues() ([]Issue, error)
	GetIssueById(id int64) (*Issue, error)
	GetIssuesByStatus(status string, projectId int64) ([]Issue, error)
}

type IssueUsecase interface {
	CreateIssue(issue *Issue) error
	UpdateIssue(issue *Issue) error
	// DeleteIssue(id int64) error
	GetAllIssues() ([]Issue, error)
	GetIssueById(id int64) (*Issue, error)
	GetIssueByStatus(projectId int64) (IssuesMap, error)
}
