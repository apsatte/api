package account_usecase

type UpdatePasswordInput struct {
	OldPassword  string
	NewPassword  string
	ClearSession bool
}
