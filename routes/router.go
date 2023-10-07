package routes

import "github.com/gofiber/fiber/v2"

type router struct {
	app *fiber.App
}

func NewRouter(app *fiber.App) *router {
	return &router{
		app: app,
	}
}

// LoadRoutes load all router for web and api
func (r *router) LoadRoutes() {
	r.app.Route("api", r.loadApiRoutes)
}
