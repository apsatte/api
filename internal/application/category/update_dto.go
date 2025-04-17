package category_usecase

import "github.com/google/uuid"

type UpdateInput struct {
	CategoryID   uuid.UUID
	Translations map[string]*UpdateInputTr
}

type UpdateInputTr struct {
	Name string
}
