package main

import (
	"github.com/cecep31/learnfiber/database"
	"github.com/cecep31/learnfiber/pkg/routes"
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

	routes.ApiRoute(app)
	routes.WebRoute(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	err := app.Listen(":3131")

	if err != nil {
		panic(err)
	}

}
