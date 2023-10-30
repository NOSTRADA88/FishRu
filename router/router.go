package router

import (
	"FishRu/controller"
	"github.com/gofiber/fiber/v2"
)

// Setup routing information

func SetUpRouters(app *fiber.App) {
	app.Get("/categories/:slug", controller.CategoryDetailBySlug)
	app.Get("/categories-cards", controller.CategoryCardsList)

	app.Get("/products", controller.ProductList)
	app.Get("/products/:slug", controller.ProductDetailSlug)

	app.Post("/auth", controller.Authorization)
	app.Get("/admin/:id", controller.ProductDetailID)
	app.Post("/admin", controller.CreateProduct)
	app.Delete("/admin/:id", controller.RemoveProduct)
	app.Put("/admin/:id", controller.UpdateProduct)

}
