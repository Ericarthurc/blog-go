package controllers

import (
	"blog-go/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSeriesIndex(c *fiber.Ctx) error {
	seriesIndex := utils.SeriesIndexParser()
	return c.Render("seriesindex", fiber.Map{"seriesIndex": seriesIndex})
}

func GetSeriesPosts(c *fiber.Ctx) error {
	series := c.Params("series")
	seriesPosts := utils.SeriesPostsParser(series)
	return c.Render("series", fiber.Map{"seriesPosts": seriesPosts, "Series": series})
}
