package auth

import (
	"context"
	"olshop/domain/model"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

// Login implements domain.AuthController.
func (a *authController) Login(c *fiber.Ctx) error {
	var creds model.LoginRequest
	if err := c.BodyParser(&creds); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()
	userData, err := a.userUC.FindByUsername(ctx, creds.Username)
	if err != nil {
		return err
	}
	return c.JSON(userData)
}
