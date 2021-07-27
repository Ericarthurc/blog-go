package routes

import (
	"blog-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	blog := app.Group("/blog")
	blog.Get("/", controllers.GetBlogIndex)
	blog.Get("/:id", controllers.GetBlog)

	// series := app.Group("/series")
	// series.Get("/")
	// series.Get("/:series")
}
