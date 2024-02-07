package routes

import (
	"example.com/app"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handlers *app.HandlersSchema) {

	SetUserRoutes(router, handlers)
	SetOrganisationRoutes(router, handlers)
	SetProjectRoutes(router, handlers)

}
