package dish_usecase

import (
	"api/internal/domain/customer"
	"api/internal/domain/menu"
	"api/internal/domain/project"
	"context"

	"github.com/google/uuid"
)

type UseCase interface {
	Add(c context.Context, dto *AddInput) error
	Update(c context.Context, dto *UpdateInput) error
	Remove(c context.Context, dishID uuid.UUID) error
	UpdatePhoto(c context.Context, dto *UpdatePhotoInput) (*UpdatePhotoOutput, error)
}

type iStorage interface {
	PutImage(c context.Context, object []byte, resizeWidth int) (string, error)
	Remove(c context.Context, objectName string) error
}

type useCase struct {
	categoryRepo menu.CategoriesRepository
	customerRepo customer.CustomersRepository
	projectRepo  project.ProjectsRepository
	dishRepo     menu.DishesRepository

	storage iStorage
}

func New(
	categoryRepo menu.CategoriesRepository,
	customerRepo customer.CustomersRepository,
	projectRepo project.ProjectsRepository,
	dishRepo menu.DishesRepository,
	storage iStorage) *useCase {
	return &useCase{
		categoryRepo: categoryRepo,
		customerRepo: customerRepo,
		projectRepo:  projectRepo,
		dishRepo:     dishRepo,
		storage:      storage,
	}
}
