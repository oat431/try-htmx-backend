package config

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupCors() cors.Config {
	return cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}
}
