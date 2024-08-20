package handlers

import (
	"rizaldyaristyo-fiber-boiler/database"
	"rizaldyaristyo-fiber-boiler/models"

	"github.com/gofiber/fiber/v2"
)

func GetTasks(c *fiber.Ctx) error {
    rows, err := database.DB.Query("SELECT id, title, description, status FROM tasks")
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    defer rows.Close()

    tasks := []models.Task{}
    for rows.Next() {
        var task models.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
            return c.Status(500).SendString(err.Error())
        }
        tasks = append(tasks, task)
    }
    return c.JSON(tasks)
}

func CreateTask(c *fiber.Ctx) error {
    task := new(models.TaskWithoutID)
    if err := c.BodyParser(task); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    _, err := database.DB.Exec("INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)", task.Title, task.Description, task.Status)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    return c.Status(201).JSON(task)
}

func GetTask(c *fiber.Ctx) error {
    id := c.Params("id")
    task := models.Task{}
    err := database.DB.QueryRow("SELECT id, title, description, status FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Title, &task.Description, &task.Status)
    if err != nil {
        return c.Status(404).SendString("Task not found")
    }
    return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
    id := c.Params("id")
    task := new(models.TaskWithoutID)

    if err := c.BodyParser(task); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    _, err := database.DB.Exec("UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?", task.Title, task.Description, task.Status, id)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
    id := c.Params("id")

    _, err := database.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
    if err != nil {
        return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(204)
}