package project_usecase

import (
	"api/internal/domain/customer"
	"api/internal/domain/project"
	"context"
)

type IStorage interface {
	PutImage(c context.Context, object []byte, resizeWidth int) (string, error)
	Remove(c context.Context, filename string) error
}

type useCase struct {
	projectsRepo  project.ProjectsRepository
	customersRepo customer.CustomersRepository
	storage       IStorage
	cdnBaseURL    string
}

func New(
	projectsRepo project.ProjectsRepository,
	customersRepo customer.CustomersRepository,
	storage IStorage,
	cdnBaseURL string) *useCase {
	return &useCase{
		customersRepo: customersRepo,
		projectsRepo:  projectsRepo,
		storage:       storage,
		cdnBaseURL:    cdnBaseURL,
	}
}
