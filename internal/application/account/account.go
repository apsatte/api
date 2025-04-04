package account_usecase

import "api/internal/domain/customer"

type useCase struct {
	customersRepo customer.CustomersRepository
	sessionsRepo  customer.SessionsRepository
}

func New(
	customersRepo customer.CustomersRepository,
	sessionsRepo customer.SessionsRepository) *useCase {
	return &useCase{
		customersRepo: customersRepo,
		sessionsRepo:  sessionsRepo,
	}
}
