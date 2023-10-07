package user

import (
	"context"
	"olshop/domain/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *userUsecase) FindByUsername(ctx context.Context, user string) (*model.UserModel, error) {
	filter := bson.D{{Key: "username", Value: user}}
	return u.repo.FindByUser(ctx, filter)
}
