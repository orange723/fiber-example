package example

import (
	"fiber-example/database"

	"github.com/gofiber/fiber/v3"
)

type Example struct {
	database.DefaultModel
	Example string `json:"example"`
}

func GetExample(c fiber.Ctx) error {
	return c.SendStatus(200)
}
