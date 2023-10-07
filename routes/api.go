package routes

import "github.com/gofiber/fiber/v2"

func (r *router) loadApiRoutes(router fiber.Router) {
	router.Route("auth", r.loadAuthRoutes)
}
