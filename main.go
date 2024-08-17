package main

import (
	"rizaldyaristyo-fiber-boiler/database"
	"rizaldyaristyo-fiber-boiler/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    database.Connect()
    app := fiber.New()

    // init routes
    routes.TaskRoutes(app)

    app.Listen(":3000")
}
