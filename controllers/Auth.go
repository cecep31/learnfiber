package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(c *fiber.Ctx) error {
	user := c.FormValue("user")
	pass := c.FormValue("pass")

	if user != "pilput" || pass != "pilput" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claim := token.Claims.(jwt.MapClaims)
	claim["name"] = "pilput"
	claim["admin"] = true
	claim["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("pilput"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token": t,
	})
}
