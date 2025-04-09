package menu

import (
	"api/internal/domain"

	"github.com/google/uuid"
)

type DishOption struct {
	ID           uuid.UUID
	MinQuantity  uint
	MaxQuantity  uint
	Position     uint
	Translations map[string]*DishOptionTranslation
	Items        []*DishOptionItem
}

type DishOptionTranslation struct {
	Name string
}

func NewDishOption(
	minQuantity, maxQuantity, position uint,
	translations map[string]*DishOptionTranslation,
	items []*DishOptionItem) (*DishOption, error) {
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &DishOption{
		ID:           ID,
		MinQuantity:  minQuantity,
		MaxQuantity:  maxQuantity,
		Position:     position,
		Translations: translations,
		Items:        items,
	}, nil
}

func (o *DishOption) GetItemByID(itemID uuid.UUID) (*DishOptionItem, error) {
	for _, item := range o.Items {
		if item.ID == itemID {
			return item, nil
		}
	}
	return nil, domain.ErrNotFound
}

func (o *DishOption) AddItem(item *DishOptionItem) error {
	for _, itm := range o.Items {
		if itm.ID == item.ID {
			return domain.ErrDishOptionItemAlreadyExists
		}
	}
	o.Items = append(o.Items, item)
	return nil
}
