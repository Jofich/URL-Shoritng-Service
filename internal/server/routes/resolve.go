package routes

import (
	"fmt"
	"log"

	storage "github.com/Jofich/URL-Shoritng-Service/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func ResolveURLHandler(database Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		alias := c.Params("url")
		url, err := database.GetUrl(alias)
		fmt.Println("url:", url)
		if err == storage.ErrURLNotExists {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "shortened link does not exist",
			})
		}
		if err != nil {
			log.Println(err.Error()+", alias:", alias)
			return nil
		}
		fmt.Println("url:", url)
		return c.Redirect(url, 301)
	}
}
