package routes

import (
	"example.com/app"
	"github.com/gin-gonic/gin"
)

func SetIssueRoutes(router *gin.Engine, handlers *app.HandlersSchema) {

	apiIssues := router.Group("api/issues")
	{

		// apiIssues.Use(middlewares.AuthMiddleware())
		apiIssues.GET("/", handlers.IssueHandler.GetIssues)
		apiIssues.POST("/", handlers.IssueHandler.AddIssue)
		apiIssues.GET("/:id", handlers.IssueHandler.GetIssueById)
		apiIssues.PUT("/:id", handlers.IssueHandler.UpdateIssue)
		apiIssues.GET("/status/:projectId", handlers.IssueHandler.GetIssueByStatus)
	}

}
