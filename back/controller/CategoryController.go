package controller

import (
	"github.com/anhht1999/back/model"
	"github.com/anhht1999/back/repo"
	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	var categories []model.Category

	result := repo.Db.
		Find(&categories)

	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(categories)
	}
}