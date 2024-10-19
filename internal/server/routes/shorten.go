package routes

import (
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/rand"
)

type Storage interface {
	SaveUrl(URLtoSave string, alias string) error
	GetUrl(alias string) (string, error)
}

type request struct {
	URL         string `json:"url"`
	AliasLength int    `json:"length"`
}

type responce struct {
	alias string `json:"alias"`
}

func isValidURL(str string) bool {
	url, err := url.Parse(str)
	if !(err == nil && url.Scheme != "" && url.Host != "") {
		return false
	}
	if !strings.Contains(url.Host, ".") {
		return false
	}
	return true
}

func generateAlias(aliasLen int) (string, error) {
	const minAliasLen = 6
	if aliasLen <= minAliasLen {
		aliasLen = minAliasLen
	}

	rand.Seed(uint64(time.Now().Unix()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	short := make([]byte, aliasLen)
	for i := range short {
		short[i] = charset[rand.Intn(len(charset))]
	}

	return string(short), nil
}

func ShortenURLHandler(database Storage) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		req := new(request)
		if err := c.BodyParser(&req); err != nil {
			log.Println("Cannot parse url", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}
		if !(isValidURL(req.URL)) {
			log.Println("URL is not Valid", req.URL)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid URL",
			})
		}
		alias, err := generateAlias(req.AliasLength)
		if err != nil {
			log.Println("Failed to generate alias", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		err = database.SaveUrl(req.URL, alias)
		if err != nil {
			log.Println("failed to save the record to the table", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Something get wrong, please try again",
			})
		}

		return nil
	}

}
