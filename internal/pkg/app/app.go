package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rajithacharith/go-template/internal/pkg/routes"
	"github.com/rajithacharith/go-template/internal/pkg/service"
)

type AppDependencies struct {
	HelloWorldService service.HelloWorldService
}

func MakeAPIServer() (*fiber.App, error) {
	deps, err := initializeDependencies()
	if err != nil {
		return nil, fmt.Errorf("failed to create app dependencies: %v", err)
	}

	app := fiber.New()

	routes.HealthRoutes(app, "/health")
	routes.HelloWorldRoutes(app, "/hello", deps.HelloWorldService)

	return app, nil
}

func initializeDependencies() (*AppDependencies, error) {
	return &AppDependencies{
		HelloWorldService: service.NewHelloWorldService(),
	}, nil
}
