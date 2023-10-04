package router

import (
	"FishRu/controller"
	"github.com/gofiber/fiber/v2"
)

// Setup routing information

func SetUpRouters(app *fiber.App) {
	app.Get("/", controller.ProductList)
	app.Post("/", controller.CreateProduct)
	app.Delete("/", controller.RemoveProduct)
	app.Put("/", controller.UpdateProduct)

	app.Post("/auth", controller.Authorization)
}
