package user

import (
	"context"
	"olshop/domain/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userRepo) FindByID(ctx context.Context, filter bson.D) (*model.UserModel, error) {
	var user = new(model.UserModel)
	if err := u.user.FindOne(ctx, filter).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}
