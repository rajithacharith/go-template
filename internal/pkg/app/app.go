package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rajithacharith/go-template/internal/pkg/routes"
)

func MakeAPIServer() (*fiber.App, error) {
	app := fiber.New()

	routes.HealthRoutes(app, "/health")

	return app, nil
}
