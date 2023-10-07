package controller

import (
	"olshop/domain/model"

	"github.com/gofiber/fiber/v2"
)

func FiberOK(c *fiber.Ctx, code int, message string, data interface{}) error {
	payload := model.OK{
		Code:    code,
		Success: true,
		Message: message,
		Data:    data,
	}
	return c.Status(code).JSON(payload)
}

func FiberErr(c *fiber.Ctx, code int, message string, detail []interface{}) error {
	payload := model.Err{
		Code:    code,
		Success: false,
		Message: message,
		Detail:  detail,
	}
	return c.Status(code).JSON(payload)
}
