package menu

import (
	"context"

	"github.com/google/uuid"
)

type CategoriesRepository interface {
	GetOneByID(c context.Context, ID uuid.UUID) (*Category, error)
	GetByProjectID(c context.Context, projectID uuid.UUID) ([]*Category, error)
	Save(c context.Context, category *Category) error
	Remove(c context.Context, category *Category) error
}
