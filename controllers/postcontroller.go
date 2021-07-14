package controllers

import (
	"github.com/cecep31/learnfiber/database"
	"github.com/cecep31/learnfiber/migrations"
	"github.com/gofiber/fiber/v2"
)

func GetDataPost(c *fiber.Ctx) error {
	var data []migrations.Post
	db := database.Con()
	db.Find(&data)
	return c.JSON(fiber.Map{
		"success": true,
		"message": "berhasil",
		"data":    data,
	})
}
