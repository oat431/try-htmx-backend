package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oat431/try-htmx-backend/controller"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	ping := app.Group("/api/v1")
	{
		ping.Get("/ping", controller.HealthCheckV1)
		ping.Get("/ping/:name", controller.HealthCheckV2)
	}

	return app
}
