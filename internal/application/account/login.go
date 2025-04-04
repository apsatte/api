package account_usecase

import (
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) Login(c context.Context, d *LoginInput) (*AuthOutput, string, error) {
	cstmr, err := u.customersRepo.GetOneByEmail(c, d.Email)
	if err != nil {
		return nil, "", err
	}

	if err := cstmr.ComparePassword(d.Password); err != nil {
		return nil, "", err
	}

	session, err := customer.NewSession(cstmr.ID, d.UserAgent, d.IP)
	if err != nil {
		return nil, "", err
	}

	if err := u.sessionsRepo.Create(c, session); err != nil {
		return nil, "", err
	}

	return NewAuthOutput(cstmr), session.AccessToken, nil
}
