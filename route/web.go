package route

import (
	"github.com/cecep31/learnfiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func WebRoute(app *fiber.App) {
	app.Static("/static", "./static")

	web := app
	web.Get("/", controllers.GetDataPosts)
}
