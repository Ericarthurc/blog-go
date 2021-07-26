package main

import (
	"blog-go/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./public")

	routes.SetupRoutes(app)

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Render("404", nil)
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
