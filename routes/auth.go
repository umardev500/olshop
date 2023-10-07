package routes

import (
	"olshop/application/controller/auth"

	"github.com/gofiber/fiber/v2"
)

func (r *router) loadAuthRoutes(router fiber.Router) {
	handler := auth.NewAuthController()
	router.Post("/login", handler.Login)
	router.Get("/logout", handler.Logout)
	router.Post("/register", handler.Register)
	router.Post("/refresh", handler.Refresh)
}
