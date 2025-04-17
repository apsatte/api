package dish_usecase

import "github.com/google/uuid"

type UpdatePhotoInput struct {
	DishID uuid.UUID
	Photo  []byte
}

type UpdatePhotoOutput struct {
	PhotoURL     string `json:"photo_url"`
	PhotoMiniURL string `json:"photo_mini_url"`
}
