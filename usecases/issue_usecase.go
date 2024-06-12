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

func (issueUC *IssueUsecase) GetIssueByStatus(projectId int64) (map[string][]entities.Issue, error) {
	statusMap := map[string][]entities.Issue{}

	status := enums.GetAllStatus()

	for index := range status {
		issues, err := issueUC.IssueRepo.GetIssuesByStatus(status[index], projectId)
		if err != nil {
			return nil, err
		}
		statusMap[status[index]] = issues
	}

	return statusMap, nil

}
