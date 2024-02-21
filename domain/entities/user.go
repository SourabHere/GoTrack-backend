package entities

type User struct {
	UserID         int    `json:"userId"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Designation_ID int    `json:"designationId"`
	DateOfJoining  string `json:"dateOfJoining"`
	UserUUID       string `json:"userUUID"`
}

type UserRepository interface {
	Save(user *User) error
	Update(user *User) error
	Delete(id string) error
	GetAllUsers() ([]User, error)
	GetUserByUUID(id string) (*User, error)
	GetProjectsByUserIdForOrganisation(userid int, organisationId string) ([]Project, error)
	GetUserIdByUUID(userUUID string) (int, error)
	GetUserOrganisationByUUID(userUUID string) ([]Organisation, error)
}

type UserUsecase interface {
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id string) error
	GetUsers() ([]User, error)
	GetUserByUUID(id string) (*User, error)
	GetProjectsByUserIdForOrganisation(userUUID string, organisationId string) ([]Project, error)
	GetUserOrganisationByUUID(userUUID string) ([]Organisation, error)
}
