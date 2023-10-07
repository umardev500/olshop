package domain

import (
	"context"
	"olshop/domain/model"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type AuthUsecase interface {
	LoginUser(ctx context.Context, creds model.LoginRequest) error
}
