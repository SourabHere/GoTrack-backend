package routes

import (
	"example.com/app"
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)

func SetOrganisationRoutes(router *gin.Engine, handlers *app.HandlersSchema) {

	apiOrganisations := router.Group("api/organisations")
	{

		apiOrganisations.Use(middlewares.AuthMiddleware())
		apiOrganisations.POST("/", handlers.OrgHandler.CreateOrganisation)
		apiOrganisations.GET("/", handlers.OrgHandler.GetOrganisations)
		apiOrganisations.GET("/:id", handlers.OrgHandler.GetOrganisationById)
		apiOrganisations.PUT("/:id", handlers.OrgHandler.UpdateOrganisation)
		apiOrganisations.DELETE("/:id", handlers.OrgHandler.DeleteOrganisation)

	}

}
