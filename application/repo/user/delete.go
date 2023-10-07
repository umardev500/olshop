package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userRepo) Delete(ctx context.Context, filter bson.D) error {
	_, err := u.user.DeleteOne(ctx, filter)
	return err
}
