package usecases

import "example.com/domain"

type OrganisationUsecase struct {
	OrganisationRepo domain.OrganisationRepository
}

func NewOrganisationUsecase(orgRepo domain.OrganisationRepository) *OrganisationUsecase {
	return &OrganisationUsecase{
		OrganisationRepo: orgRepo,
	}
}

func (orgUC *OrganisationUsecase) CreateOrganisation(org *domain.Organisation) error {
	return orgUC.OrganisationRepo.Save(org)
}

func (orgUC *OrganisationUsecase) UpdateOrganisation(org *domain.Organisation) error {
	return orgUC.OrganisationRepo.Update(org)
}

func (orgUC *OrganisationUsecase) DeleteOrganisation(id int64) error {
	return orgUC.OrganisationRepo.Delete(id)
}

func (orgUC *OrganisationUsecase) GetAllOrganisations() ([]domain.Organisation, error) {
	return orgUC.OrganisationRepo.GetAllOrganisations()
}

func (orgUC *OrganisationUsecase) GetOrganisationById(id int64) (*domain.Organisation, error) {
	return orgUC.OrganisationRepo.GetOrganisationById(id)
}
