package menu

import (
	"context"

	"github.com/google/uuid"
)

type DishesRepository interface {
	GetOneByID(c context.Context, dishID uuid.UUID) (*Dish, error)
	GetByCategoryID(c context.Context, categoryID uuid.UUID) ([]*Dish, error)
	Save(c context.Context, dish *Dish) error
}
