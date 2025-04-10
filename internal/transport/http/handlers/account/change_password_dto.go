package account_handler

type changePasswordReq struct {
	OldPassword  string `json:"old_password" example:"badboy2015" validate:"required"`
	NewPassword  string `json:"new_password"  example:"QcU+,317QL[t" validate:"required"`
	ClearSession bool   `json:"clear_session" example:"false"`
}
