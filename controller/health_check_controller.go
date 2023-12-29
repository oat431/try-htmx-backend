package controller

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheckV1(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func HealthCheckV2(c *fiber.Ctx) error {
	name := c.Params("name")
	return c.JSON(fiber.Map{
		"message": "Hello " + name,
	})
}
