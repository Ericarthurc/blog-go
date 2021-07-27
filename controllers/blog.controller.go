package controllers

import (
	"blog-go/utils"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

func GetBlogIndex(c *fiber.Ctx) error {
	blogIndex := utils.BlogIndexParser()
	return c.Render("blogindex", fiber.Map{"blogIndex": blogIndex})
}

func GetBlog(c *fiber.Ctx) error {
	id := c.Params("id")
	body, meta := utils.BlogPostParser(id)
	return c.Render("blogpost", fiber.Map{"Body": template.HTML(body), "Meta": meta})
}
