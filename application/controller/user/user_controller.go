package user

import (
	"olshop/domain"

	"github.com/go-playground/validator/v10"
)

type userController struct {
	usecase  domain.UserUsecase
	validate *validator.Validate
}

func NewUserController(usecase domain.UserUsecase, validate *validator.Validate) domain.UserController {
	return &userController{
		usecase:  usecase,
		validate: validate,
	}
}
