package application

import (
	"context"
	"fmt"
	"olshop/routes"
	"os"
	"time"

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
	router := routes.NewRouter(app)
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
