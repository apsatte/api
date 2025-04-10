package project_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) Update(c context.Context, d *UpdateInput) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	project, err := u.projectsRepo.GetOneByID(c, d.ProjectID)
	if err != nil {
		return err
	}

	if !project.IsOwner(session.CustomerID) {
		return domain.ErrForbidden
	}

	translations := DTOTranslationsToDomain(d.Translations)
	project.Update(d.Name, d.ServiceFee, d.Languages, translations)

	return u.projectsRepo.Save(c, project)
}
