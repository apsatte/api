package project_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"context"

	"github.com/google/uuid"
)

func (u *useCase) Remove(c context.Context, projectID uuid.UUID) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	project, err := u.projectsRepo.GetOneByID(c, projectID)
	if err != nil {
		return err
	}

	if project.CustomerID != session.CustomerID {
		return domain.ErrForbidden
	}

	err = u.projectsRepo.Remove(c, project)
	if err != nil {
		return err
	}

	u.storage.Remove(c, project.LogoURL)
	u.storage.Remove(c, project.BackgroundURL)

	return nil
}
