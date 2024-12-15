package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required", min=4`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"omitempty"`
}
