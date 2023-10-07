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

func (u *userController) UpdatePassword(c *fiber.Ctx) error {
	var params model.UserUpdatePassRequest
	if err := c.BodyParser(&params); err != nil {
		return controller.FiberErr(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// validate
	detail, err := helper.ValidateStruct(u.validate, params)
	if err != nil {
		return controller.FiberErr(c, fiber.StatusUnprocessableEntity, constants.ValidationErr, detail)
	}

	// for production change password string as hashed
	var user model.UserModel = model.UserModel{
		Password: params.NewPassword,
	}

	var id string = "65217a1e7ed964eceb93af07"

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()
	err = u.usecase.UpdateUserPassword(ctx, id, user)
	if err != nil {
		return controller.FiberErr(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return controller.FiberOK(c, fiber.StatusOK, constants.UpdatePassOK, nil)
}
