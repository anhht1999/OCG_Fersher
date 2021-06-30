package routes

import (
	"github.com/anhht1999/back/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigRouter(app *fiber.App) {

	productAPI := app.Group("/api/product")
	routeProduct(&productAPI)

	categoryAPI := app.Group("/api/category")
	routeCategory(&categoryAPI)
}

func routeProduct(router *fiber.Router) {
	(*router).Get("/", controller.GetProducts) //Liệt  cate
	// (*router).Get("/", controller.GetProducts1) //Liệt kê
	(*router).Get("/:id", controller.GetProductById) //Liệt kê
}

func routeCategory(router *fiber.Router) {
	(*router).Get("/", controller.GetCategories) //Liệt kê
}
