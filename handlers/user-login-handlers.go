package handlers

import (
	"fmt"
	"os"
	"rizaldyaristyo-fiber-boiler/database"
	"rizaldyaristyo-fiber-boiler/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {

    // load secret from env
    var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

    // parse request body to user struct
    var user models.User
    fmt.Println(c.BodyParser(&user))
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    // check if user exists
    var storedPassword string
    err := database.DB.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&storedPassword)
    if err != nil {
        return c.Status(401).SendString("Invalid username or password")
    }

    // compare hash
    err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
    if err != nil {
        return c.Status(401).SendString("Invalid username or password")
    }

    // generate token
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = user.Username
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // 72 hours

    t, err := token.SignedString(jwtSecret)
    if err != nil {
        return c.Status(500).SendString("Could not login")
    }

    return c.JSON(fiber.Map{
        "token": t,
    })
}
