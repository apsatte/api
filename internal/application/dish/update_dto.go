package dish_usecase

import (
	"api/internal/domain/menu"

	"github.com/google/uuid"
)

type UpdateInput struct {
	DishID       uuid.UUID
	CategoryID   uuid.UUID
	Price        uint
	IsAvailable  bool
	Translations map[string]*menu.DishTranslation
}
