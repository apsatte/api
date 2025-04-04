package account_handler

type loginReq struct {
	Email    string `json:"email" example:"user@example.com" validate:"required"`
	Password string `json:"password"  example:"qwerty123" validate:"required"`
}
