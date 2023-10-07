package user

import (
	"olshop/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	user *mongo.Collection
}

func NewUserRepo(user *mongo.Collection) domain.UserRepo {
	return &userRepo{
		user: user,
	}
}
