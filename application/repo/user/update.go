package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userRepo) Update(ctx context.Context, filter bson.D, payload bson.D) error {
	var update bson.M = bson.M{"$set": payload}
	_, err := u.user.UpdateOne(ctx, filter, update)
	return err
}
