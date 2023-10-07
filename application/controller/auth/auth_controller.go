package auth

import (
	"olshop/domain"
)

type authController struct {
	userUC domain.UserUsecase
}

// NewAuthController create new auth controller instance
func NewAuthController(userUC domain.UserUsecase) domain.AuthController {
	return &authController{
		userUC: userUC,
	}
}
