package category_usecase

import (
	"api/internal/domain/customer"
	"api/internal/domain/menu"
	"api/internal/domain/project"
)

type useCase struct {
	categoryRepo menu.CategoriesRepository
	customerRepo customer.CustomersRepository
	projectRepo  project.ProjectsRepository
}

func New(
	categoryRepo menu.CategoriesRepository,
	customerRepo customer.CustomersRepository,
	projectRepo project.ProjectsRepository) *useCase {
	return &useCase{
		categoryRepo: categoryRepo,
		customerRepo: customerRepo,
		projectRepo:  projectRepo,
	}
}
