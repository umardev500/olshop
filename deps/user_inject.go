package deps

import (
	userRepository "olshop/application/repo/user"
	userUsecase "olshop/application/usecase/user"
	"olshop/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

func UserInject(db *mongo.Database) domain.UserUsecase {
	userCollection := db.Collection("users")
	userRepo := userRepository.NewUserRepo(userCollection)
	userUC := userUsecase.NewUserUsecase(userRepo)
	return userUC
}
