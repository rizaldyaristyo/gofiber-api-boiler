package handlers

import (
	"rizaldyaristyo-fiber-boiler/database"
	"rizaldyaristyo-fiber-boiler/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
    
    // parse request body to user struct
    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    // bcrypt hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(500).SendString("Could not hash the password")
    }
    user.Password = string(hashedPassword)

    // query
    _, err = database.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
    if err != nil {
        return c.Status(500).SendString("Could not create user")
    }

    return c.Status(201).SendString("User registered successfully")
}
