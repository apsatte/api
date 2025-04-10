package account_handler

type updateReq struct {
	Name string `json:"name" example:"John Doe" vaildate:"required"`
}
