package auth

import (
	"context"
	"olshop/application/controller"
	"olshop/constants"
	"olshop/domain/model"
	"olshop/helper"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

// Register implements domain.AuthController.
func (a *authController) Register(c *fiber.Ctx) error {
	var params model.RegisterRequest
	if err := c.BodyParser(&params); err != nil {
		return controller.FiberErr(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// validate
	detail, err := helper.ValidateStruct(a.validate, params)
	if err != nil {
		return controller.FiberErr(c, fiber.StatusUnprocessableEntity, constants.ValidationErr, detail)
	}

	// Bindig params to model
	var user model.UserModel = model.UserModel{
		Username: params.Username,
		Password: params.Password,
		Detail: model.UserDetailModel{
			Email: params.Email,
		},
	}

	// Call usecase
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()
	err = a.userUC.CreateUser(ctx, user)
	if err != nil {
		return controller.FiberErr(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return controller.FiberOK(c, fiber.StatusOK, constants.RegisterOK, nil)
}
