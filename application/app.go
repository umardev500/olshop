package application

import (
	"context"
	"fmt"
	"olshop/config"
	"olshop/routes"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type application struct{}

func NewApplication() *application {
	return &application{}
}

func (a *application) Start(ctx context.Context) error {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	validate := validator.New()
	db := config.NewMongo().Database("olshop")
	router := routes.NewRouter(app, db, validate)
	router.LoadRoutes()

	ch := make(chan error, 1)
	go func() {
		port := os.Getenv("PORT")
		log.Info().Msg(fmt.Sprintf("Server running on port %s", port))
		err := app.Listen(port)
		if err != nil {
			ch <- err
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		fmt.Println()
		log.Info().Msg("Gracefully shutdown please wait...")
		app.ShutdownWithTimeout(10 * time.Second)
		log.Info().Msg("Good bye!")
	}

	return nil
}
