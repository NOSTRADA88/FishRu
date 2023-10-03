package router

import (
	"FishRu/controller"
	"github.com/gofiber/fiber/v2"
)

// Setup routing information

func SetUpRouters(app *fiber.App) {
	// list => get
	// add => post

	// update => put in work -_-

	// delete => delete
	app.Get("/", controller.ProductList)
	app.Post("/", controller.CreateProduct)
	app.Delete("/", controller.RemoveProduct)
	app.Put("/", controller.UpdateProduct)
}
