package menu

import (
	"api/internal/domain"
	"time"

	"slices"

	"github.com/google/uuid"
)

type Dish struct {
	ID           uuid.UUID
	CategoryID   uuid.UUID
	Position     uint
	PhotoURL     string
	PhotoMiniURL string
	Price        uint
	IsAvailable  bool
	Translations map[string]*DishTranslation
	Options      []*DishOption
	CreatedAt    time.Time
}

type DishTranslation struct {
	Name        string
	Description string
}

func NewDish(
	categoryID uuid.UUID,
	position, price uint,
	photoURL, photoMiniURL string,
	isAvailable bool,
	translations map[string]*DishTranslation) (*Dish, error) {
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Dish{
		ID:           ID,
		CategoryID:   categoryID,
		Position:     position,
		PhotoURL:     photoMiniURL,
		PhotoMiniURL: photoMiniURL,
		Price:        price,
		IsAvailable:  isAvailable,
		Translations: translations,
		CreatedAt:    time.Now().UTC(),
	}, nil
}

func (d *Dish) GetOptionByID(optionID uuid.UUID) (*DishOption, error) {
	for _, option := range d.Options {
		if option.ID == optionID {
			return option, nil
		}
	}
	return nil, domain.ErrNotFound
}

func (d *Dish) AddOption(option *DishOption) error {
	for _, existing := range d.Options {
		if existing.ID == option.ID {
			return domain.ErrDishOptionAlreadyExists
		}
	}

	d.Options = append(d.Options, option)

	return nil
}

func (d *Dish) RemoveOption(optionID uuid.UUID) {
	for index, opt := range d.Options {
		if opt.ID == optionID {
			d.Options = slices.Delete(d.Options, index, index+1)
			break
		}
	}
}
