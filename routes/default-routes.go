package routes

import (
	"rizaldyaristyo-fiber-boiler/handlers"

	"github.com/gofiber/fiber/v2"
)

func DefaultRoutes(app *fiber.App) {
    app.Get("/", handlers.Index)
}