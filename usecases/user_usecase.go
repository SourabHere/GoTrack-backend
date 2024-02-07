package usecases

import "example.com/domain"

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

func (userUC *UserUsecase) GetUserById(id string) (*domain.User, error) {
	return userUC.userRepo.GetUserById(id)
}

func (userUC *UserUsecase) UpdateUser(user *domain.User) error {
	return userUC.userRepo.Update(user)
}

func (userUC *UserUsecase) DeleteUser(id string) error {
	return userUC.userRepo.Delete(id)
}
