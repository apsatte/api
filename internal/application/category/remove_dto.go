package category_usecase

import "github.com/google/uuid"

type RemoveInput struct {
	ProjectID  uuid.UUID
	CategoryID uuid.UUID
}
