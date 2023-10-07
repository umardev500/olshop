package routes

import (
	"olshop/application/controller/user"
	"olshop/deps"

	"github.com/gofiber/fiber/v2"
)

func (r *router) loadUserRoutes(router fiber.Router) {
	uc := deps.UserInject(r.db)
	handler := user.NewUserController(uc, r.validate)
	router.Put("/change/pass", handler.UpdatePassword)
}
