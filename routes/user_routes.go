package routes

import (
	"example.com/app"
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)

func SetUserRoutes(router *gin.Engine, handlers *app.HandlersSchema) {

	apiUsers := router.Group("api/users")
	{

		apiUsers.GET("/", handlers.UserHandler.GetUsers)
		apiUsers.POST("/", handlers.UserHandler.UserManager)
		apiUsers.GET("/:uuid", handlers.UserHandler.GetUserById)
		apiUsers.GET("/:uuid/organisations", handlers.UserHandler.GetUserOrganisationsByUUID)

		apiUsers.GET("/designation/:designationId", handlers.UserHandler.GetUserDesignationByID)

		apiUsers.Use(middlewares.VerifyUser())
		apiUsers.DELETE("/:uuid", handlers.UserHandler.DeleteUser)
		apiUsers.PUT("/:uuid", handlers.UserHandler.UpdateUser)
		apiUsers.GET("/:uuid/projects", handlers.UserHandler.GetProjectsByUserIdForOrganisation)
		apiUsers.POST("/:uuid/projects", handlers.UserHandler.RegisterUserProject)

	}

}
