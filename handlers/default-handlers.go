package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.SendFile("./public/index.html")
	// return c.Status(200).SendString("Hello, World!")
}
