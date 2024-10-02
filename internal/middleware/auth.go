package middleware

import (
	"go-auth/internal/modules/helper"

	"github.com/gofiber/fiber/v2"
)

func Protected(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false, "message": "invalid token",
		})
		return nil
	}

	token = token[len("Bearer "):]

	err := helper.VerifyToken(token)
	if err != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false, "message": "invalid token",
		})
		return nil
	}

	c.Next()
	return nil
}
