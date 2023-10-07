package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type router struct {
	app *fiber.App
	db  *mongo.Database
}

func NewRouter(app *fiber.App, db *mongo.Database) *router {
	return &router{
		app: app,
		db:  db,
	}
}

// LoadRoutes load all router for web and api
func (r *router) LoadRoutes() {
	r.app.Route("api", r.loadApiRoutes)
}
