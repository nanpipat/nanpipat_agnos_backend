package requests

type PasswordRequest struct {
	InitPassword string `json:"init_password" binding:"required,min=1,max=40"`
}
