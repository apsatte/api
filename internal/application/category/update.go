package category_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"api/internal/domain/menu"
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

	category, err := u.categoryRepo.GetOneByID(c, dto.CategoryID)
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

	// update category
	translations := map[string]*menu.CategoryTranslation{}
	for lang, tr := range dto.Translations {
		translations[lang] = &menu.CategoryTranslation{
			Name: tr.Name,
		}
	}
	category.Update(translations)
	if err := u.categoryRepo.Save(c, category); err != nil {
		return err
	}

	return nil
}
