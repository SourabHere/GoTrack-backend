package http

import (
	"net/http"
	"strconv"

	"example.com/domain/entities"
	"example.com/domain/requests"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	ProjectUsecase entities.ProjectUsecase
}

func NewProjectHandler(projectUsecase entities.ProjectUsecase) *ProjectHandler {
	return &ProjectHandler{
		ProjectUsecase: projectUsecase,
	}
}

func (projectHandler *ProjectHandler) GetProjects(context *gin.Context) {

	projects, err := projectHandler.ProjectUsecase.GetProjects()

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get projects",
			"error":   err.Error(),
		})

		return

	}

	context.JSON(200, projects)
}

func (projectHandler *ProjectHandler) GetProjectById(context *gin.Context) {
	projectId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
			"error":   err.Error(),
		})

		return

	}

	project, err := projectHandler.ProjectUsecase.GetProjectById(projectId)

	if err != nil {

		context.JSON(http.StatusNotFound, gin.H{
			"message": "invalid id",
			"error":   err.Error(),
		})

		return

	}

	context.JSON(200, project)
}

func (projectHandler *ProjectHandler) AddProject(context *gin.Context) {

	var project requests.CreateProjectRequest

	err := context.BindJSON(&project)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values found",
		})

		return
	}

	categoryName := project.ProjectCategory

	projectCategoryID, err := projectHandler.ProjectUsecase.GetProjectCategoryIDByName(categoryName)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values found",
		})

		return
	}

	organisationID, err := strconv.Atoi(project.OrganisationID)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid organisation found",
		})

		return

	}

	var projectObj = &entities.Project{
		ProjectName:         project.ProjectName,
		Project_Desc:        project.ProjectDesc,
		Project_Category_ID: projectCategoryID,
		Project_URL:         &project.ProjectURL,
		Organisation_ID:     organisationID,
	}

	addedProject, err := projectHandler.ProjectUsecase.CreateProject(projectObj)

	if err != nil {
		context.JSON(http.StatusFailedDependency, gin.H{
			"message": "could not create project",
		})

		return
	}

	context.JSON(200, addedProject)

}

func (projectHandler *ProjectHandler) UpdateProject(context *gin.Context) {

	var project requests.UpdateProjectRequest

	err := context.ShouldBind(&project)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values found in request",
		})

		return
	}

	categoryName := project.ProjectCategory

	projectCategoryID, err := projectHandler.ProjectUsecase.GetProjectCategoryIDByName(categoryName)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values found",
		})

		return
	}

	organisationID, err := strconv.Atoi(project.OrganisationID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid organisation found",
		})

		return
	}

	projectID, err := strconv.Atoi(project.ProjectID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid organisation found",
		})

		return
	}

	var projectObj = &entities.Project{
		ProjectID:           projectID,
		ProjectName:         project.ProjectName,
		Project_Desc:        project.ProjectDesc,
		Project_Category_ID: projectCategoryID,
		Project_URL:         &project.ProjectURL,
		Organisation_ID:     organisationID,
	}

	err = projectHandler.ProjectUsecase.UpdateProject(projectObj)

	if err != nil {

		context.JSON(http.StatusFailedDependency, gin.H{
			"message": "could not update project",
		})

		return
	}

	context.JSON(200, gin.H{
		"message": "project updated successfully",
	})

}

func (projectHandler *ProjectHandler) DeleteProject(context *gin.Context) {
	projectId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {

		context.JSON(http.StatusNotFound, gin.H{
			"message": "invalid id",
			"error":   err.Error(),
		})

		return

	}

	err = projectHandler.ProjectUsecase.DeleteProject(projectId)

	if err != nil {

		context.JSON(http.StatusExpectationFailed, gin.H{
			"message": "could not delete project",
		})

		return
	}

	context.JSON(200, gin.H{
		"message": "project deleted successfully",
	})
}

func (projectHandler *ProjectHandler) GetProjectCategories(context *gin.Context) {

	projectTypes, err := projectHandler.ProjectUsecase.GetProjectCategories()

	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"message": "could not get project categories",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(200, projectTypes)

}

func (projectHandler *ProjectHandler) GetProjectCategoryById(context *gin.Context) {

	categoryId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
			"error":   err.Error(),
		})

		return

	}

	projectType, err := projectHandler.ProjectUsecase.GetProjectCategoryById(int(categoryId))

	if err != nil {

		context.JSON(http.StatusNotFound, gin.H{
			"message": "invalid id",
			"error":   err.Error(),
		})

		return

	}

	context.JSON(200, projectType)
}
