package project_usecase

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"api/internal/domain/project"
	"context"
)

func (u *useCase) Add(c context.Context, d *AddProjectInput) error {
	session, err := customer.GetSessionFromContext(c)
	if err != nil {
		return err
	}

	projects, err := u.projectsRepo.GetByCustomerID(c, session.CustomerID)
	if err != nil {
		return err
	}

	customer, err := u.customersRepo.GetOneByID(c, session.CustomerID)
	if err != nil {
		return err
	}

	if len(projects) >= int(customer.Subscription.ProjectsLimit) {
		return domain.ErrProjectsLimit
	}

	logoURL, err := u.storage.PutImage(c, d.Logo, 500)
	if err != nil {
		return err
	}

	backgroundURL, err := u.storage.PutImage(c, d.Logo, 800)
	if err != nil {
		return err
	}

	project, err := project.New(session.CustomerID, d.Name, logoURL, backgroundURL, d.ServiceFee, d.Languages)
	if err != nil {
		u.storage.Remove(c, logoURL)
		u.storage.Remove(c, backgroundURL)
		return err
	}

	if err := u.projectsRepo.Save(c, project); err != nil {
		u.storage.Remove(c, logoURL)
		u.storage.Remove(c, backgroundURL)
		return err
	}

	return nil
}
