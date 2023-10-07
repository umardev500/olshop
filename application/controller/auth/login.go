package auth

import (
	"context"
	"olshop/application/controller"
	"olshop/domain/model"
	"olshop/helper"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

// Login implements domain.AuthController.
func (a *authController) Login(c *fiber.Ctx) error {
	var creds model.LoginRequest
	if err := c.BodyParser(&creds); err != nil {
		return controller.FiberErr(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// validate
	detail, err := helper.ValidateStruct(a.validate, creds)
	if err != nil {
		return controller.FiberErr(c, fiber.StatusUnprocessableEntity, "validation error", detail)
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()
	userData, err := a.userUC.FindByUsername(ctx, creds.Username)
	if err != nil {
		return err
	}
	return c.JSON(userData)
}
