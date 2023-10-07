package user

import (
	"context"
	"olshop/domain/model"
	"olshop/helper"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userUsecase) CreateUser(ctx context.Context, req model.UserCreateRequest) error {
	var user model.UserModel = model.UserModel{
		Username: req.Username,
		Password: req.Password,
		Detail: model.UserDetailModel{
			Email: req.Email,
		},
		CreatedAt: helper.GetUnixTime(),
	}

	var payload *bson.D = new(bson.D)
	helper.ToBsonD(user, payload, false)
	err := u.repo.Create(ctx, *payload)
	return err
}
