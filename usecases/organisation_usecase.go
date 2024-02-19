package usecases

import "example.com/domain/entities"

type OrganisationUsecase struct {
	OrganisationRepo entities.OrganisationRepository
}

func NewOrganisationUsecase(orgRepo entities.OrganisationRepository) *OrganisationUsecase {
	return &OrganisationUsecase{
		OrganisationRepo: orgRepo,
	}
}

func (orgUC *OrganisationUsecase) CreateOrganisation(org *entities.Organisation) error {
	return orgUC.OrganisationRepo.Save(org)
}

func (orgUC *OrganisationUsecase) UpdateOrganisation(org *entities.Organisation) error {
	return orgUC.OrganisationRepo.Update(org)
}

func (orgUC *OrganisationUsecase) DeleteOrganisation(id int64) error {
	return orgUC.OrganisationRepo.Delete(id)
}

func (orgUC *OrganisationUsecase) GetAllOrganisations() ([]entities.Organisation, error) {
	return orgUC.OrganisationRepo.GetAllOrganisations()
}

func (orgUC *OrganisationUsecase) GetOrganisationById(id int64) (*entities.Organisation, error) {
	return orgUC.OrganisationRepo.GetOrganisationById(id)
}
