package routes

import (
	"example.com/app"
	"github.com/gin-gonic/gin"
)

func SetUserRoutes(router *gin.Engine, handlers *app.HandlersSchema) {

	apiUsers := router.Group("/users")
	{

		apiUsers.GET("/", handlers.UserHandler.GetUsers)
		apiUsers.POST("/", handlers.UserHandler.CreateUser)
		apiUsers.GET("/:uuid", handlers.UserHandler.GetUserById)

		apiUsers.DELETE("/:uuid", handlers.UserHandler.DeleteUser)
		// apiUsers.Use(middlewares.AuthMiddleware())

		apiUsers.PUT("/:uuid", handlers.UserHandler.UpdateUser)

	}

}
