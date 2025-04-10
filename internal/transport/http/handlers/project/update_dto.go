package projects_handler

type updateReq struct {
	Name         string                           `json:"name" validate:"required"`
	ServiceFee   uint                             `json:"service_fee" validate:"required"`
	Languages    []string                         `json:"languages" validate:"required"`
	Translations map[string]*updateReqTranslation `json:"translations" validate:"required"`
}

type updateReqTranslation struct {
	Description string `json:"description" validate:"required"`
}
