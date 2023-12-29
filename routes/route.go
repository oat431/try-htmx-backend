package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/oat431/try-htmx-backend/config"
	"github.com/oat431/try-htmx-backend/controller"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(config.SetupCors()))

	ping := app.Group("/api/v1")
	{
		ping.Get("/ping", controller.HealthCheckV1)
		ping.Get("/ping/:name", controller.HealthCheckV2)
	}

	return app
}
