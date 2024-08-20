package main

import (
	"rizaldyaristyo-fiber-boiler/database"
	"rizaldyaristyo-fiber-boiler/routes"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    database.Connect()
    app := fiber.New()
    app.Static("/assets", "./public/assets")

    // init routes
    routes.DefaultRoutes(app)
    routes.TaskRoutes(app)
    routes.UserRoutes(app)

    app.Listen(":3000")
}
