package user

import (
	"context"
	"olshop/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (u *userRepo) Find(ctx context.Context, filter bson.D, opts *options.FindOptions) ([]model.UserModel, error) {
	cur, err := u.user.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var users = make([]model.UserModel, 0)
	for cur.Next(ctx) {
		var each model.UserModel
		if err := cur.Decode(&each); err != nil {
			return nil, err
		}
		users = append(users, each)
	}

	return users, nil
}
