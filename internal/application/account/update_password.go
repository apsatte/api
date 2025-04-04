package account_usecase

import (
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) UpdatePassword(c context.Context, d *UpdatePasswordInput) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	customer, err := u.customersRepo.GetOneByID(c, session.CustomerID)
	if err != nil {
		return err
	}

	err = customer.ChangePassword(d.OldPassword, d.NewPassword)
	if err != nil {
		return err
	}

	if err := u.customersRepo.Save(c, customer); err != nil {
		return err
	}

	if d.ClearSession {
		u.sessionsRepo.RemoveAll(c, customer.ID)
		u.sessionsRepo.Create(c, session)
	}

	return nil
}
