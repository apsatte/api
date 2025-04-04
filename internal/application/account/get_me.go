package account_usecase

import (
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) GetMe(c context.Context) (*AuthOutput, error) {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return nil, err
	}

	customer, err := u.customersRepo.GetOneByID(c, session.CustomerID)
	if err != nil {
		return nil, err
	}

	return NewAuthOutput(customer), nil
}
