package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Protected protect routes
func Jsoncheck(c *fiber.Ctx) error {
	if c.Is("json") {
		return c.Next()
	}
	return c.SendString("Only Json Allowed!")
}
