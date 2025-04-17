package dish_usecase

import (
	"api/internal/domain/menu"

	"github.com/google/uuid"
)

type AddInput struct {
	CategoryID   uuid.UUID
	Position     uint
	Price        uint
	Photo        []byte
	IsAvailable  bool
	Translations map[string]*menu.DishTranslation

	Options []*AddInputOption
}

type AddInputOption struct {
	MinQuantity  uint
	MaxQuantity  uint
	Position     uint
	Translations map[string]*menu.DishOptionTranslation
	Items        []*AddInputOptionItem
}

type AddInputOptionItem struct {
	MaxQuantity   uint
	PriceModifier uint
	Position      uint
	Translations  map[string]*menu.DishOptionItemTranslation
}
