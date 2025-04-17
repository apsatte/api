package dish_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) Update(c context.Context, dto *UpdateInput) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	cstmr, err := u.customerRepo.GetOneByID(c, session.CustomerID)
	if err != nil {
		return err
	}

	// permission
	dish, err := u.dishRepo.GetOneByID(c, dto.DishID)
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

	// update
	_, err = u.categoryRepo.GetOneByID(c, dto.CategoryID)
	if err != nil {
		return err
	}

	dish.Update(dto.CategoryID, dto.Price, dto.IsAvailable, dto.Translations)
	if err := u.dishRepo.Save(c, dish); err != nil {
		return err
	}

	return nil
}
