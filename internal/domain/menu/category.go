package menu

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID           uuid.UUID
	ProjectID    uuid.UUID
	Translations map[string]*CategoryTranslation
	Position     uint
	CreatedAt    time.Time
}

type CategoryTranslation struct {
	Name string
}

func NewCategory(position uint, projectID uuid.UUID, translations map[string]*CategoryTranslation) (*Category, error) {
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Category{
		ID:           ID,
		ProjectID:    projectID,
		Translations: translations,
		Position:     position,
		CreatedAt:    time.Now().UTC(),
	}, nil
}
