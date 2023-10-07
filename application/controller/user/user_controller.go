package user

import (
	"olshop/domain"
)

type userController struct{}

func NewUserController() domain.UserController {
	return &userController{}
}
