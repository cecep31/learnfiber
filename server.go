package main

import (
	"log"

	"github.com/cecep31/learnfiber/database"
	"github.com/cecep31/learnfiber/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//migrate data base
	//remove // to migrate aplikation
	// db.AutoMigrate(&migrate)
	database.ConPsql()

	//app route fiber
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	//routing static file

	route.ApiRoute(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3131"))

}
