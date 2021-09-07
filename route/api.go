package route

import (
	"fmt"

	"github.com/cecep31/learnfiber/controllers"
	"github.com/cecep31/learnfiber/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ApiRoute(app *fiber.App) {

	app.Post("/login", controllers.Login)
	//api routeing
	api := app.Group("/api")
	api.Use(middleware.Jsoncheck)
	api.Use(cors.New())

	api.Get("/post", controllers.GetDataPosts)
	api.Post("/post", middleware.Protected(), controllers.AddDataPost)

	//404 page not found
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("halaman %s tidak ditemukan 👋", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})
}
