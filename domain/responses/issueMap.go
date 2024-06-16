package responses

import "example.com/domain/entities"

type IssueMapResponse struct {
	IssuesMap entities.IssuesMap `json:"issuesMap"`
	ProjectID int                `json:"project_id"`
}
