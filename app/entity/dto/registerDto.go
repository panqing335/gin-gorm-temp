package dto

type RegisterDto struct {
	Username string `json:"username" binding:"required,min=3,max=10"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}
