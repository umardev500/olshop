package user

import (
	"context"
	"olshop/domain/model"
)

func (u *userUsecase) FindByUsername(ctx context.Context, user string) (*model.UserModel, error) {
	return u.repo.FindByUser(ctx, user)
}
