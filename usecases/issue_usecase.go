package usecases

import (
	"example.com/domain/entities"
	"example.com/domain/enums"
)

type IssueUsecase struct {
	IssueRepo entities.IssueRepository
}

func NewIssueUsecase(issueRepo entities.IssueRepository) *IssueUsecase {
	return &IssueUsecase{
		IssueRepo: issueRepo,
	}
}

func (issueUC *IssueUsecase) CreateIssue(issue *entities.Issue) error {
	return issueUC.IssueRepo.Save(issue)
}

func (issueUC *IssueUsecase) GetAllIssues() ([]entities.Issue, error) {
	return issueUC.IssueRepo.GetAllIssues()
}

func (issueUC *IssueUsecase) GetIssueById(id int64) (*entities.Issue, error) {
	return issueUC.IssueRepo.GetIssueById(id)
}

func (issueUC *IssueUsecase) UpdateIssue(issue *entities.Issue) error {
	return issueUC.IssueRepo.Update(issue)
}

func (issueUC *IssueUsecase) GetIssueByStatus(projectId int64) (entities.IssuesMap, error) {
	var issuesMap entities.IssuesMap

	status := enums.GetAllStatus()

	for _, s := range status {
		issues, err := issueUC.IssueRepo.GetIssuesByStatus(s, projectId)
		if err != nil {
			return entities.IssuesMap{}, err
		}

		switch s {
		// case enums.To_Do:
		// 	issuesMap.ToDo = append(issuesMap.ToDo, issues...)
		case enums.Open:
			issuesMap.Open = append(issuesMap.Open, issues...)
		case enums.In_Progress:
			issuesMap.InProgress = append(issuesMap.InProgress, issues...)
		case enums.Review:
			issuesMap.Review = append(issuesMap.Review, issues...)
		case enums.Closed:
			issuesMap.Closed = append(issuesMap.Closed, issues...)
		}
	}

	return issuesMap, nil
}
