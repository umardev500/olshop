package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// Create create new user
func (u *userRepo) Create(ctx context.Context, payload bson.D) error {
	_, err := u.user.InsertOne(ctx, payload)
	return err
}
