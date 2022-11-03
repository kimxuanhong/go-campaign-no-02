package dto

type UserLogin struct {
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}
