package requests

type CreateUserRequest struct {
	Data struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     []struct {
			EmailAddress string `json:"email_address"`
		} `json:"email_addresses"`

		UserUUID string `json:"id"`
	} `json:"data"`

	Type string `json:"type"`
}
