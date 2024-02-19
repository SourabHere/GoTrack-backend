package usecases

import (
	"example.com/domain/entities"
)

type UserUsecase struct {
	userRepo entities.UserRepository
}

func NewUserUsecase(userRepo entities.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (userUC *UserUsecase) CreateUser(user *entities.User) error {
	return userUC.userRepo.Save(user)
}

func (userUC *UserUsecase) GetUsers() ([]entities.User, error) {
	return userUC.userRepo.GetAllUsers()
}

func (userUC *UserUsecase) GetUserById(uuid string) (*entities.User, error) {
	return userUC.userRepo.GetUserById(uuid)
}

func (userUC *UserUsecase) UpdateUser(user *entities.User) error {

	userUUID := user.UserUUID

	storedUser, err := userUC.GetUserById(userUUID)

	if err != nil {
		return err
	}

	user.Designation_ID = storedUser.Designation_ID

	return userUC.userRepo.Update(user)
}

func (userUC *UserUsecase) DeleteUser(id string) error {
	return userUC.userRepo.Delete(id)
}

func (userUC *UserUsecase) GetProjectsByUserIdForOrganisation(userUUID string, organisationId string) ([]entities.Project, error) {
	user, err := userUC.GetUserById(userUUID)

	if err != nil {
		return nil, err
	}

	return userUC.userRepo.GetProjectsByUserIdForOrganisation(user.UserID, organisationId)

}
