package main

import (
	"fmt"

	"github.com/cecep31/learnfiber/controllers"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func SetupRoute(app *fiber.App) {
	app.Static("/", "./static")

	app.Post("/login", controllers.Login)
	//api routeing
	api := app.Group("/api")

	api.Get("/post", controllers.GetDataPosts)
	api.Post("/post", controllers.AddDataPost)

	//auth jwt
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("pilput"),
	}))

	//404 page not found
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("halaman %s tidak ditemukan 👋", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})
}
