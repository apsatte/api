package category_usecase

import "github.com/google/uuid"

type UpdatePositionInput struct {
	Categories []UpdatePositionInputCategory
}

type UpdatePositionInputCategory struct {
	ID       uuid.UUID
	Position uint
}
