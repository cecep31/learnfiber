package main

import (
	"fmt"
	"log"

	"github.com/cecep31/learnfiber/controllers"
	"github.com/cecep31/learnfiber/database"
	"github.com/cecep31/learnfiber/migrations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//migrate data base
	db := database.Con()
	var migrate migrations.Post
	db.AutoMigrate(&migrate)
	route := fiber.New()
	route.Use(logger.New())
	route.Use(recover.New())

	//routing static file
	route.Static("/", "./static")

	//api routeing
	api := route.Group("/api")
	api.Get("/", controllers.GetDataPost)

	//404 page not found
	route.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("halaman %s tidak ditemukan 👋", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	log.Fatal(route.Listen(":3131"))

}
