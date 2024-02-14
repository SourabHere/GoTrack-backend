package http

import (
	"example.com/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (UserHandler *UserHandler) CreateUser(context *gin.Context) {

	var user domain.User

	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{
			"error":   err,
			"message": "invalid request body",
		})

		return
	}

	err = UserHandler.userUsecase.CreateUser(&user)

	if err != nil {
		context.JSON(400, gin.H{
			"error":   err.Error(),
			"message": "could not create user",
		})

		return
	}

	context.JSON(200, gin.H{
		"message": "user created successfully",
	})

}

func (userHandler *UserHandler) GetUsers(context *gin.Context) {
	users, err := userHandler.userUsecase.GetUsers()
	if err != nil {
		context.JSON(400, gin.H{
			"message": "could not get users",
		})
	}

	context.JSON(200, users)

}

func (userHandler *UserHandler) GetUserById(context *gin.Context) {
	uuid := context.Param("uuid")

	if uuid == "" {
		context.JSON(400, gin.H{
			"message": "invalid request body",
		})
	}

	user, err := userHandler.userUsecase.GetUserById(uuid)

	if err != nil {
		context.JSON(400, gin.H{
			"message": "could not get user",
		})

		return
	}

	context.JSON(200, user)
}

func (UserHandler *UserHandler) UpdateUser(context *gin.Context) {

	uuid := context.Param("uuid")

	if uuid == "" {
		context.JSON(400, gin.H{
			"message": "invalid request body",
		})
	}

	var user domain.User

	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{
			"message": "invalid request body",
		})

		return
	}

	user.UserUUID = uuid

	err = UserHandler.userUsecase.UpdateUser(&user)

	if err != nil {
		context.JSON(400, gin.H{
			"message": "could not update user",
		})

		return
	}

	context.JSON(200, gin.H{
		"message": "user updated successfully",
	})

}

func (userHandler *UserHandler) DeleteUser(context *gin.Context) {

	uuid := context.Param("uuid")

	if uuid == "" {
		context.JSON(400, gin.H{
			"message": "invalid request body",
		})
	}

	err := userHandler.userUsecase.DeleteUser(uuid)

	if err != nil {
		context.JSON(400, gin.H{
			"message": "could not delete user",
		})

		return
	}

	context.JSON(200, gin.H{
		"message": "user deleted successfully",
	})

}

func (userHandler *UserHandler) GetProjectsByUserIdForOrganisation(context *gin.Context) {
	userUUID := context.Param("uuid")

	if userUUID == "" {
		context.JSON(400, gin.H{
			"message": "invalid request body",
		})

		return
	}

	organisationId := context.Query("organisationId")

	if organisationId == "" {
		context.JSON(400, gin.H{
			"message": "Organisation not found",
		})

		return
	}

	projects, err := userHandler.userUsecase.GetProjectsByUserIdForOrganisation(userUUID, organisationId)

	if err != nil {
		context.JSON(400, gin.H{
			"error":   err.Error(),
			"message": "could not get projects",
		})

		return
	}

	context.JSON(200, projects)

}
