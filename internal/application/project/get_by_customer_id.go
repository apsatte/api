package project_usecase

import (
	"api/internal/domain/customer"
	"context"
)

func (u *useCase) GetByCustomerID(c context.Context) ([]*ProjectOutput, error) {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return nil, err
	}

	projects, err := u.projectsRepo.GetByCustomerID(c, session.CustomerID)
	if err != nil {
		return nil, err
	}

	return MapProjectsToOutput(u.cdnBaseURL, projects), nil
}
