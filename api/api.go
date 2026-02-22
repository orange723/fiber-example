package api

import (
	"fiber-example/api/example"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")
	example.Routes(v1)
}
