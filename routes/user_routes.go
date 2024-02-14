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
		apiUsers.POST("/", handlers.UserHandler.CreateUser)
		apiUsers.GET("/:uuid", handlers.UserHandler.GetUserById)

		apiUsers.Use(middlewares.VerifyUser())
		apiUsers.DELETE("/:uuid", handlers.UserHandler.DeleteUser)
		apiUsers.PUT("/:uuid", handlers.UserHandler.UpdateUser)
		apiUsers.GET("/:uuid/projects", handlers.UserHandler.GetProjectsByUserIdForOrganisation)

	}

}
