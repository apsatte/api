package menu

import "github.com/google/uuid"

type DishOptionItem struct {
	ID            uuid.UUID
	MinQuantity   uint
	MaxQuantity   uint
	PriceModifier uint
	Position      uint
	Translations  map[string]*DishOptionItemTranslation
}

type DishOptionItemTranslation struct {
	Name string
}

func NewDishOptionItem(
	maxQuantity, priceModifier, position uint,
	translations map[string]*DishOptionItemTranslation) (*DishOptionItem, error) {
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &DishOptionItem{
		ID:            ID,
		MaxQuantity:   maxQuantity,
		PriceModifier: priceModifier,
		Position:      position,
		Translations:  translations,
	}, nil
}
