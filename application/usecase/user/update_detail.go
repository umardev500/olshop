package user

import (
	"context"
	"olshop/domain/model"
	"olshop/helper"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userUsecase) UpdateUserDetail(ctx context.Context, id string, user model.UserModel) error {
	user.UpdatedAt = helper.GetUnixTime()
	var payload *bson.D = new(bson.D)
	helper.ToBsonD(user, payload, false)
	var filter bson.D = bson.D{
		{Key: "_id", Value: helper.GetObjID(id)},
	}
	u.repo.Update(ctx, filter, *payload)
	return nil
}
