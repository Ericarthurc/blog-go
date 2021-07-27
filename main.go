package main

import (
	"blog-go/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		ProxyHeader: "X-Real-IP",
		Views:       engine,
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.Render("404", fiber.Map{"Url": c.OriginalURL(), "IP": c.IP()})
		},
	})

	app.Use(recover.New())

	app.Static("/", "./public/root")
	app.Static("/static", "./public")

	routes.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/blog")
	})

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Render("404", fiber.Map{"Url": c.OriginalURL(), "IP": c.IP()})
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
