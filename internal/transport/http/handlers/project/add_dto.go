package projects_handler

type addReq struct {
	Name        string   `json:"name" validate:"required" example:"Burger Queen"`
	Description string   `json:"description" validate:"required" example:"The Best Restaurant in the City."`
	Logo        string   `json:"logo" validate:"required" example:"aGVsbG8gd29ybGQ="`
	Background  string   `json:"background" validate:"required" example:"aGVsbG8gd29ybGQ="`
	ServiceFee  uint     `json:"service_fee" validate:"required" example:"10"`
	Languages   []string `json:"languages" validate:"required" example:"kz,en"`
}
