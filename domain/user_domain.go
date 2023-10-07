package domain

import (
	"context"
	"olshop/domain/model"
)

type UserController interface{}

type UserUsecase interface {
	FindByUsername(ctx context.Context, user string) (*model.UserModel, error)
}

type UserRepo interface {
	FindByUser(ctx context.Context, user string) (*model.UserModel, error)
}
