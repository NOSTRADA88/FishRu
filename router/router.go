package router

import (
	"FishRu/controller"
	"github.com/gofiber/fiber/v2"
)

// Setup routing information

func SetUpRouters(app *fiber.App) {
	app.Get("/", controller.ProductList)
	app.Post("/admin", controller.CreateProduct)
	app.Delete("/admin/:id", controller.RemoveProduct)
	app.Put("/admin/:id", controller.UpdateProduct)

	app.Get("/admin/:id", controller.ProductDetail)

	app.Post("/auth", controller.Authorization)
}
