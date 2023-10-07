package model

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=8"`
}

// UserCreateRequest is a body to create new user
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,min=6"`
}
