package routes

import (
	"rizaldyaristyo-fiber-boiler/handlers"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App) {
    app.Get("/tasks", handlers.GetTasks)
    app.Post("/tasks", handlers.CreateTask)
    app.Get("/tasks/:id", handlers.GetTask)
    app.Put("/tasks/:id", handlers.UpdateTask)
    app.Delete("/tasks/:id", handlers.DeleteTask)
}