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

	return u.projectsRepo.Remove(c, project)
}
