package routes

import (
	"olshop/application/controller/auth"
	"olshop/deps"

	"github.com/gofiber/fiber/v2"
)

// loadAuthRoutes load auth routes
func (r *router) loadAuthRoutes(router fiber.Router) {
	userUC := deps.UserInject(r.db)
	handler := auth.NewAuthController(userUC)
	router.Post("/login", handler.Login)
	router.Get("/logout", handler.Logout)
	router.Post("/register", handler.Register)
	router.Post("/refresh", handler.Refresh)
}
