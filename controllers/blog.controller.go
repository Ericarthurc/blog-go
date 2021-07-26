package controllers

import (
	"blog-go/utils"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

// func GetBlogIndex(c *fiber.Ctx) error {
// 	return c.Render()
// }

func GetBlog(c *fiber.Ctx) error {
	body, _ := utils.ParserBlog()
	return c.Render("index", fiber.Map{"body": template.HTML(body)})
}
