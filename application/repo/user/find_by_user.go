package user

import (
	"context"
	"olshop/domain/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userRepo) FindByUser(ctx context.Context, filter bson.D) (*model.UserModel, error) {
	var result = new(model.UserModel)
	err := u.user.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
