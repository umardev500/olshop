package auth

import (
	"olshop/domain"

	"github.com/go-playground/validator/v10"
)

type authController struct {
	userUC   domain.UserUsecase
	validate *validator.Validate
}

// NewAuthController create new auth controller instance
func NewAuthController(userUC domain.UserUsecase, validate *validator.Validate) domain.AuthController {
	return &authController{
		userUC:   userUC,
		validate: validate,
	}
}
