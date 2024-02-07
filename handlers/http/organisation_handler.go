package http

import (
	"net/http"
	"strconv"

	"example.com/domain"
	"github.com/gin-gonic/gin"
)

type Organisationhandler struct {
	OrganisationUsecase domain.OrganisationUsecase
}

func NewOrganisationHandler(orgUsecase domain.OrganisationUsecase) *Organisationhandler {
	return &Organisationhandler{
		OrganisationUsecase: orgUsecase,
	}
}

func (orgHandler *Organisationhandler) GetOrganisations(context *gin.Context) {

	organisations, err := orgHandler.OrganisationUsecase.GetAllOrganisations()

	if err != nil {

		context.JSON(http.StatusExpectationFailed, gin.H{
			"message": "could not get organisations",
		})

		return
	}

	context.JSON(200, organisations)

}

func (orgHandler *Organisationhandler) GetOrganisationById(context *gin.Context) {
	orgId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})

		return
	}

	organisation, err := orgHandler.OrganisationUsecase.GetOrganisationById(orgId)

	if err != nil {

		context.JSON(http.StatusNotFound, gin.H{
			"message": "invalid id",
		})

		return
	}

	context.JSON(200, organisation)
}

func (orgHandler *Organisationhandler) CreateOrganisation(context *gin.Context) {
	var org domain.Organisation

	err := context.ShouldBindJSON(&org)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = orgHandler.OrganisationUsecase.CreateOrganisation(&org)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "organisation created successfully", "organisation": org})

}

func (orgHandler *Organisationhandler) UpdateOrganisation(context *gin.Context) {

	orgId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})

		return
	}

	var organisation domain.Organisation

	err = context.ShouldBindJSON(&organisation)

	organisation.Organisation_ID = int(orgId)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid data",
		})

		return
	}

	err = orgHandler.OrganisationUsecase.UpdateOrganisation(&organisation)

	if err != nil {

		context.JSON(http.StatusExpectationFailed, gin.H{
			"message": "could not update organisation",
		})

		return
	}

	context.JSON(200, gin.H{
		"message": "organisation updated successfully",
	})

}

func (orgHandler *Organisationhandler) DeleteOrganisation(context *gin.Context) {
	organisation_id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
			"error":   err.Error(),
		})
	}

	err = orgHandler.OrganisationUsecase.DeleteOrganisation(organisation_id)

	if err != nil {

		context.JSON(http.StatusFailedDependency, gin.H{
			"message": "could not delete data",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(200, gin.H{
		"message": "organisation deleted successfully",
	})

}
