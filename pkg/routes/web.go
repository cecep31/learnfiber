package routes

import (
	"github.com/cecep31/learnfiber/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func WebRoute(app *fiber.App) {
	app.Static("/static", "./static")

	web := app
	web.Get("/", controllers.GetDataPosts)
}
