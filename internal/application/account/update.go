package account_usecase

import (
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) Update(c context.Context, d *UpdateInput) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	customer, err := u.customersRepo.GetOneByID(c, session.CustomerID)
	if err != nil {
		return err
	}

	customer.Update(d.Name)
	return u.customersRepo.Save(c, customer)
}
