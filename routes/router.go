package routes

import (
	"blog-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	blog := app.Group("/blog")
	blog.Get("/", controllers.GetBlog)
	// blog.Get("/:id")

	// series := app.Group("/series")
	// series.Get("/")
	// series.Get("/:series")
}
