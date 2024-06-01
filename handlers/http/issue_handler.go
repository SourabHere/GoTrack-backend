package http

import (
	"net/http"

	"strconv"

	"example.com/domain/entities"
	"github.com/gin-gonic/gin"
)

type IssueHandler struct {
	IssueUsecase entities.IssueUsecase
}

func NewIssueHandler(issueUsecase entities.IssueUsecase) *IssueHandler {
	return &IssueHandler{
		IssueUsecase: issueUsecase,
	}
}

func (issueHandler *IssueHandler) GetIssues(context *gin.Context) {

	issues, err := issueHandler.IssueUsecase.GetAllIssues()

	if err != nil {

		context.JSON(http.StatusExpectationFailed, gin.H{
			"message": "could not get issues",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(200, issues)

}

func (issueHandler *IssueHandler) GetIssueById(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
			"error":   err.Error(),
		})

		return
	}

	issue, err := issueHandler.IssueUsecase.GetIssueById(id)

	if err != nil {

		context.JSON(http.StatusExpectationFailed, gin.H{
			"message": "could not get issue",
		})

		return
	}

	context.JSON(http.StatusAccepted, issue)

}

func (issueHandler *IssueHandler) AddIssue(context *gin.Context) {

	var issue entities.Issue

	err := context.ShouldBind(&issue)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": "Error in the given issue params"})

		return

	}

	err = issueHandler.IssueUsecase.CreateIssue(&issue)

	if err != nil {

		context.JSON(http.StatusExpectationFailed, gin.H{
			"message": "could not create issue",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "issue created successfully",
		"issue":   issue,
	})

}
