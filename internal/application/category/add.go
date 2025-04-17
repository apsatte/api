package category_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) Add(c context.Context, dto *AddInput) error {
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

	category, err := NewCategoryFromDTO(dto)
	if err != nil {
		return err
	}

	return u.categoryRepo.Save(c, category)
}
