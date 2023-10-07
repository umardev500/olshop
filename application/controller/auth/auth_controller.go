package auth

import (
	"olshop/domain"
)

type authController struct{}

// NewAuthController create new auth controller instance
func NewAuthController() domain.AuthController {
	return &authController{}
}
