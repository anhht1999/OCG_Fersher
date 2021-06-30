package controller

import (
	"github.com/anhht1999/back/model"
	"github.com/anhht1999/back/repo"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

func GetProducts(c *fiber.Ctx) error {
	var products []model.Product
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	offset := (page - 1) * limit
	sort := c.Query("sort")
	order := c.Query("order")
	q := c.Query("q")
	// categoryId,_:=strconv.Atoi(c.Query("categoryId"))

	text := "%" + q + "%"
	s := []string{sort, order}

	result := repo.Db.Preload("Category").
		Order(strings.Join(s, " ")).
		Where("Name LIKE ?", text).Or("Author LIKE ?", text).
		Offset(offset).
		Limit(limit).
		Find(&products)

	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(products)
	}
}

func GetProductById(c *fiber.Ctx) error {
	var product model.Product
	id, _ := strconv.Atoi(c.Params("id"))
	result := repo.Db.
		Preload("Category").
		Where("ID = ?", id).
		Find(&product)

	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(product)
	}
}
