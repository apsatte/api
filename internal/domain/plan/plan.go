package plan

import (
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	ID            uuid.UUID
	ProjectsLimit uint
	Translations  map[string]*PlanTranslation
	Options       map[uuid.UUID]*PlanOption
	CreatedAt     time.Time
}

type PlanTranslation struct {
	Name        string
	Description string
}

type PlanOption struct {
	DurationDays uint
	Price        uint
}

func NewPlan(projectsLimit uint, translations map[string]*PlanTranslation, options map[uuid.UUID]*PlanOption) (*Plan, error) {
	ID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Plan{
		ID:            ID,
		ProjectsLimit: projectsLimit,
		Translations:  translations,
		Options:       options,
		CreatedAt:     time.Now().UTC(),
	}, nil
}
