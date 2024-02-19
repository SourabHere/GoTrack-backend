package entities

import (
	"time"
)

type Organisation struct {
	Organisation_ID       int       `json:"organisationId"`
	Organisation_Name     string    `json:"organisationName"`
	Organisation_Type     bool      `json:"organisationType"`
	Organisation_URL      string    `json:"organisationURL"`
	Organisation_Logo     []byte    `json:"organisationLogo"`
	Organisation_Location string    `json:"organisationLocation"`
	Created_At            time.Time `json:"createdAt"`
}

type OrganisationRepository interface {
	Save(org *Organisation) error
	Update(org *Organisation) error
	Delete(id int64) error
	GetAllOrganisations() ([]Organisation, error)
	GetOrganisationById(id int64) (*Organisation, error)
}

type OrganisationUsecase interface {
	CreateOrganisation(org *Organisation) error
	UpdateOrganisation(org *Organisation) error
	DeleteOrganisation(id int64) error
	GetAllOrganisations() ([]Organisation, error)
	GetOrganisationById(id int64) (*Organisation, error)
}
