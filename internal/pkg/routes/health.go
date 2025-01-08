package routes

import "github.com/gofiber/fiber/v2"

func HealthRoutes(app fiber.Router, path string) {
	// Health routes
	routes := app.Group(path)

	routes.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "UP",
		})
	})
}
