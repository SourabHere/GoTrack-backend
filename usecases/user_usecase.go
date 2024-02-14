package usecases

import (
	"example.com/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (userUC *UserUsecase) CreateUser(user *domain.User) error {
	return userUC.userRepo.Save(user)
}

func (userUC *UserUsecase) GetUsers() ([]domain.User, error) {
	return userUC.userRepo.GetAllUsers()
}

func (userUC *UserUsecase) GetUserById(uuid string) (*domain.User, error) {
	return userUC.userRepo.GetUserById(uuid)
}

func (userUC *UserUsecase) UpdateUser(user *domain.User) error {
	return userUC.userRepo.Update(user)
}

func (userUC *UserUsecase) DeleteUser(id string) error {
	return userUC.userRepo.Delete(id)
}

func (userUC *UserUsecase) GetProjectsByUserIdForOrganisation(userUUID string, organisationId string) ([]domain.Project, error) {
	user, err := userUC.GetUserById(userUUID)

	if err != nil {
		return nil, err
	}

	return userUC.userRepo.GetProjectsByUserIdForOrganisation(user.UserID, organisationId)

}
