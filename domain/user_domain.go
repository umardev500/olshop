package domain

import (
	"context"
	"olshop/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserController interface{}

type UserUsecase interface {
	FindByUsername(ctx context.Context, user string) (*model.UserModel, error)
}

type UserRepo interface {
	Find(ctx context.Context, filter bson.D, opts *options.FindOptions) ([]model.UserModel, error)
	FindByID(ctx context.Context, filter bson.D) (*model.UserModel, error)
	FindByUser(ctx context.Context, user string) (*model.UserModel, error)
	Create(ctx context.Context, payload bson.D) error
	Delete(ctx context.Context, filter bson.D) error
	Update(ctx context.Context, filter bson.D, payload bson.D) error
}
