package example

import "github.com/gofiber/fiber/v3"

func Routes(route fiber.Router) {
	route.Get("/example", GetExample)
}
