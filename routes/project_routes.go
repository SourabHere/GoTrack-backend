package routes

import (
	"example.com/app"
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)

func SetProjectRoutes(router *gin.Engine, handlers *app.HandlersSchema) {

	apiProjects := router.Group("api/projects")
	{
		apiProjects.GET("/categories", handlers.ProjectHandler.GetProjectCategories)
		apiProjects.GET("/categories/:id", handlers.ProjectHandler.GetProjectCategoryById)

		apiProjects.Use(middlewares.AuthMiddleware())

		apiProjects.POST("/", handlers.ProjectHandler.AddProject)
		apiProjects.GET("/", handlers.ProjectHandler.GetProjects)
		apiProjects.GET("/:id", handlers.ProjectHandler.GetProjectById)
		apiProjects.PUT("/:id", handlers.ProjectHandler.UpdateProject)
		apiProjects.DELETE("/:id", handlers.ProjectHandler.DeleteProject)
	}
}
