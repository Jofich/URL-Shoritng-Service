package server

import (
	"fmt"

	config "github.com/Jofich/URL-Shoritng-Service/internal/config"
	"github.com/Jofich/URL-Shoritng-Service/internal/server/handler"
	storage "github.com/Jofich/URL-Shoritng-Service/internal/storage/postgres"
	"github.com/gofiber/fiber/v2"
)

func Start(conf *config.HTTPServer, database *storage.Storage) error {
	const fn = "server.Start"
	App := fiber.New()

	handler.SetupRoutes(App, database)

	err := App.Listen(conf.Address)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
