package project

import (
	"context"

	"github.com/google/uuid"
)

type ProjectsRepository interface {
	GetByCustomerID(c context.Context, customerID uuid.UUID) ([]*Project, error)
	GetOneByID(c context.Context, projectID uuid.UUID) (*Project, error)
	Save(c context.Context, p *Project) error
	Remove(c context.Context, p *Project) error
}
