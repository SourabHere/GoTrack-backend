package requests

type CreateProjectRequest struct {
	ProjectName     string `json:"projectName"`
	ProjectDesc     string `json:"projectDesc"`
	ProjectURL      string `json:"projectURL"`
	ProjectCategory string `json:"projectCategory"`
	OrganisationID  string `json:"organisationID"`
	UserUUID        string `json:"userUUID"`
}
type UpdateProjectRequest struct {
	ProjectID       string `json:"projectId"`
	ProjectName     string `json:"projectName"`
	ProjectDesc     string `json:"projectDesc"`
	ProjectURL      string `json:"projectURL"`
	ProjectCategory string `json:"projectCategory"`
	OrganisationID  string `json:"organisationID"`
	UserUUID        string `json:"userUUID"`
}

type RegisterProject struct {
	ProjectID int    `json:"projectId"`
	UserUUID  string `json:"userUUID"`
}
