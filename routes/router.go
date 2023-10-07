package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type router struct {
	app      *fiber.App
	db       *mongo.Database
	validate *validator.Validate
}

func NewRouter(app *fiber.App, db *mongo.Database, validate *validator.Validate) *router {
	return &router{
		app:      app,
		db:       db,
		validate: validate,
	}
}

// LoadRoutes load all router for web and api
func (r *router) LoadRoutes() {
	r.app.Route("api", r.loadApiRoutes)
}
