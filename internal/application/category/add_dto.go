package category_usecase

import (
	"api/internal/domain/menu"

	"github.com/google/uuid"
)

type AddInput struct {
	ProjectID    uuid.UUID
	Position     uint
	Translations map[string]*AddInputTr
}

type AddInputTr struct {
	Name string
}

func NewCategoryFromDTO(dto *AddInput) (*menu.Category, error) {
	translations := make(map[string]*menu.CategoryTranslation)

	for lang, tr := range dto.Translations {
		translations[lang] = &menu.CategoryTranslation{
			Name: tr.Name,
		}
	}

	return menu.NewCategory(dto.Position, dto.ProjectID, translations)
}
