package main

import (
	"log"

	"github.com/cecep31/learnfiber/database"
	"github.com/cecep31/learnfiber/migrations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//migrate data base
	//remove // to migrate aplikation
	db := database.Con()
	var migrate migrations.Post
	db.AutoMigrate(&migrate)

	//app route fiber
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	//routing static file
	SetupRoute(app)

	log.Fatal(app.Listen(":3131"))

}
