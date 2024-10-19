package handler

import (
	"github.com/Jofich/URL-Shoritng-Service/internal/server/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, database routes.Storage) {
	app.Post("/api/v1", routes.ShortenURLHandler(database))
	app.Get("/:url", routes.ResolveURLHandler(database))
}
