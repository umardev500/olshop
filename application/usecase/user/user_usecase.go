package user

import (
	"olshop/domain"
)

type userUsecase struct {
	repo domain.UserRepo
}

func NewUserUsecase(repo domain.UserRepo) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}
