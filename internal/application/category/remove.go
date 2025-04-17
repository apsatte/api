package category_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) Remove(c context.Context, dto *RemoveInput) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	cstmr, err := u.customerRepo.GetOneByID(c, session.CustomerID)
	if err != nil {
		return err
	}

	prjct, err := u.projectRepo.GetOneByID(c, dto.ProjectID)
	if err != nil {
		return err
	}

	if !prjct.IsOwner(cstmr.ID) {
		return domain.ErrForbidden
	}

	category, err := u.categoryRepo.GetOneByID(c, dto.CategoryID)
	if err != nil {
		return err
	}

	if !category.HasProjectID(prjct.ID) {
		return domain.ErrForbidden
	}

	return u.categoryRepo.Remove(c, category)
}
