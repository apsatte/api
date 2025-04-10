package project_usecase

import (
	"api/internal/domain/project"

	"github.com/google/uuid"
)

type UpdateInput struct {
	ProjectID    uuid.UUID
	Name         string
	ServiceFee   uint
	Languages    []string
	Translations map[string]*UpdateInputTranslation
}

type UpdateInputTranslation struct {
	Description string
}

func DTOTranslationsToDomain(dto map[string]*UpdateInputTranslation) map[string]*project.ProjectTranslation {
	output := make(map[string]*project.ProjectTranslation)

	for lang, tr := range dto {
		output[lang] = &project.ProjectTranslation{
			Description: tr.Description,
		}
	}

	return output
}
