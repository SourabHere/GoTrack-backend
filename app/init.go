package app

import (
	"database/sql"

	"example.com/handlers/http"
	"example.com/repositories"
	"example.com/usecases"
)

type HandlersSchema struct {
	OrgHandler     *http.Organisationhandler
	UserHandler    *http.UserHandler
	ProjectHandler *http.ProjectHandler
}

func IntitialiseHandlers(db *sql.DB) *HandlersSchema {

	orgRepo := repositories.NewOrganisationRepository(db)
	orgUsecase := usecases.NewOrganisationUsecase(orgRepo)
	orgHandler := http.NewOrganisationHandler(orgUsecase)

	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := http.NewUserHandler(userUsecase)

	projectRepo := repositories.NewProjectRepository(db)
	projectUsecase := usecases.NewProjectUsecase(projectRepo)
	projectHandler := http.NewProjectHandler(projectUsecase)

	handlers := &HandlersSchema{
		OrgHandler:     orgHandler,
		UserHandler:    userHandler,
		ProjectHandler: projectHandler,
	}

	return handlers

}
