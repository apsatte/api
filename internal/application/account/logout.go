package account_usecase

import (
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) Logout(c context.Context) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	return u.sessionsRepo.Remove(c, session)
}
