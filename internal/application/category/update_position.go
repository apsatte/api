package category_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) UpdatePosition(c context.Context, dto *UpdatePositionInput) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	cstmr, err := u.customerRepo.GetOneByID(c, session.CustomerID)
	if err != nil {
		return err
	}

	for _, cat := range dto.Categories {
		category, err := u.categoryRepo.GetOneByID(c, cat.ID)
		if err != nil {
			return err
		}

		prjct, err := u.projectRepo.GetOneByID(c, category.ID)
		if err != nil {
			return err
		}

		if !prjct.IsOwner(cstmr.ID) {
			return domain.ErrForbidden
		}

		category.UpdatePosition(cat.Position)
		if err := u.categoryRepo.Save(c, category); err != nil {
			return err
		}
	}

	return nil
}
