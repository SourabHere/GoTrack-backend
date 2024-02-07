package http

import (
	"net/http"
	"strconv"

	"example.com/domain"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	ProjectUsecase domain.ProjectUsecase
}

func NewProjectHandler(projectUsecase domain.ProjectUsecase) *ProjectHandler {
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
	var project *domain.Project

	err := context.BindJSON(&project)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values found",
			"error":   err.Error(),
		})

		return
	}

	err = projectHandler.ProjectUsecase.CreateProject(project)

	if err != nil {

		context.JSON(http.StatusFailedDependency, gin.H{
			"message": "could not create project",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(200, project)

}

func (projectHandler *ProjectHandler) UpdateProject(context *gin.Context) {

	projectId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})

		return
	}

	var project domain.Project

	err = context.ShouldBind(&project)

	project.ProjectID = int(projectId)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid values found",
			"error":   err.Error(),
		})

		return
	}

	err = projectHandler.ProjectUsecase.UpdateProject(&project)

	if err != nil {

		context.JSON(http.StatusFailedDependency, gin.H{
			"message": "could not update project",
			"error":   err.Error(),
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
