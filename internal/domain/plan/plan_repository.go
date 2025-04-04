package plan

import (
	"context"

	"github.com/google/uuid"
)

type PlansRepository interface {
	GetAll(c context.Context) ([]*Plan, error)
	GetOneByOptionID(c context.Context, ID uuid.UUID) (*Plan, error)
}
