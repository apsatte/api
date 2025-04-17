package dish_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"

	"github.com/google/uuid"
)

func (u *useCase) Remove(c context.Context, dishID uuid.UUID) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	cstmr, err := u.customerRepo.GetOneByID(c, session.CustomerID)
	if err != nil {
		return err
	}

	// permission
	dish, err := u.dishRepo.GetOneByID(c, dishID)
	if err != nil {
		return err
	}

	category, err := u.categoryRepo.GetOneByID(c, dish.CategoryID)
	if err != nil {
		return err
	}

	prjct, err := u.projectRepo.GetOneByID(c, category.ProjectID)
	if err != nil {
		return err
	}

	if !prjct.IsOwner(cstmr.ID) {
		return domain.ErrForbidden
	}

	// remove
	return u.dishRepo.Remove(c, dish)
}
