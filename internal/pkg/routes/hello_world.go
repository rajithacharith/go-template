package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rajithacharith/go-template/internal/pkg/service"
)

func HelloWorldRoutes(app fiber.Router, path string, service service.HelloWorldService) {
	// Health routes
	routes := app.Group(path)

	routes.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": service.SayHello(),
		})
	})
}
