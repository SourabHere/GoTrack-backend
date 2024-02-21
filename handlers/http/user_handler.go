package http

import (
	"example.com/domain/entities"
	"example.com/domain/requests"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase entities.UserUsecase
}

func NewUserHandler(userUsecase entities.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (UserHandler *UserHandler) CreateUser(context *gin.Context) {

	var user entities.User

	var req requests.CreateUserRequest

	err := context.ShouldBindJSON(&req)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if req.Type == "user.deleted" {
		userUUID := req.Data.UserUUID

		err := UserHandler.userUsecase.DeleteUser(userUUID)

		if err != nil {

			context.JSON(400, gin.H{
				"message": "could not delete user",
			})

			return
		}

		context.JSON(200, gin.H{
			"message": "user deleted successfully",
		})

		return
	}

	user.FirstName = req.Data.FirstName
	user.LastName = req.Data.LastName
	user.Email = req.Data.Email[0].EmailAddress
	user.UserUUID = req.Data.UserUUID

	if req.Type == "user.updated" {

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

		return

	} else {

		user.Designation_ID = 1

		err = UserHandler.userUsecase.CreateUser(&user)

		if err != nil {
			context.JSON(400, gin.H{
				"message": "could not create user",
			})

			return
		}

		context.JSON(200, gin.H{
			"message": "user created successfully",
		})

	}
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

	user, err := userHandler.userUsecase.GetUserByUUID(uuid)

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

	var user entities.User

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

func (userHandler *UserHandler) GetUserOrganisationsByUUID(context *gin.Context) {
	userUUID := context.Param("uuid")

	if userUUID == "" {
		context.JSON(400, gin.H{
			"message": "invalid request body",
		})

		return
	}

	organisations, err := userHandler.userUsecase.GetUserOrganisationByUUID(userUUID)

	if err != nil {
		context.JSON(400, gin.H{
			"message": "could not get organisations",
		})

		return
	}

	context.JSON(200, organisations)

}
