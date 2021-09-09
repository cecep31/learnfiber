package controllers

import (
	"github.com/cecep31/learnfiber/app/models"
	"github.com/cecep31/learnfiber/database"
	"github.com/cecep31/learnfiber/migrations"
	"github.com/gofiber/fiber/v2"
)

func GetDataPosts(c *fiber.Ctx) error {
	var data []migrations.Post
	db := database.DB.Db
	db.Find(&data)
	return c.JSON(fiber.Map{
		"success": true,
		"message": "berhasil",
		"data":    data,
	})
}

func AddDataPost(c *fiber.Ctx) error {
	db := database.DB.Db
	input := new(models.Post)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	data := migrations.Post{Title: input.Title, Body: input.Body}
	db.Create(&data)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "berhasil",
		"data":    data,
	})

}
