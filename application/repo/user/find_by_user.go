package user

import (
	"context"
	"olshop/domain/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userRepo) FindByUser(ctx context.Context, user string) (*model.UserModel, error) {
	var result = new(model.UserModel)

	var filter bson.M = bson.M{
		"username": user,
	}
	err := u.user.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
