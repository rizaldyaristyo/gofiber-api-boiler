package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Profile(c *fiber.Ctx) error {
    username := c.Locals("user")

    if username == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "user not found in context"})
    }

    return c.JSON(fiber.Map{
        "username": username,
    })
}
