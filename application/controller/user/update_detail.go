package user

import (
	"context"
	"olshop/application/controller"
	"olshop/constants"
	"olshop/domain/model"
	"olshop/helper"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (u *userController) UpdateDetail(c *fiber.Ctx) error {
	var params model.UserUpdateDetailRequest
	if err := c.BodyParser(&params); err != nil {
		return controller.FiberErr(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// validate
	detail, err := helper.ValidateStruct(u.validate, params)
	if err != nil {
		return controller.FiberErr(c, fiber.StatusUnprocessableEntity, constants.ValidationErr, detail)
	}
	var users model.UserModel = model.UserModel{
		Detail: model.UserDetailModel{
			Name:    params.Name,
			Address: params.Address,
			Email:   params.Email,
			Phone:   params.Phone,
		},
	}

	var id string = "65217a1e7ed964eceb93af07"
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()
	err = u.usecase.UpdateUserDetail(ctx, id, users)
	if err != nil {
		return controller.FiberErr(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return controller.FiberOK(c, fiber.StatusOK, constants.UpdatePassOK, nil)
}
