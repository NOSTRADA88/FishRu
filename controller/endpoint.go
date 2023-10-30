package controller

import (
	"FishRu/database"
	"FishRu/types"
	"github.com/gofiber/fiber/v2"
	"os"
)

func ProductList(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	context := fiber.Map{}

	prodSlice, err := database.SelectAll(connection)
	if err != nil {
		context["status"] = fiber.StatusNotFound
		context["error"] = err
		return ctx.JSON(context)
	}
	context["status"] = fiber.StatusOK
	context["data"] = prodSlice
	return ctx.JSON(context)
}

func CategoryCardsList(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	context := fiber.Map{}

	prodSlice, err := database.SelectCategoryCard(connection)
	if err != nil {
		context["status"] = fiber.StatusNotFound
		context["error"] = err
		return ctx.JSON(context)
	}
	context["status"] = fiber.StatusOK
	context["data"] = prodSlice
	return ctx.JSON(context)
}

func CreateProduct(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	context := fiber.Map{}
	product := types.ProductCard{}

	form, err := ctx.MultipartForm()
	if err != nil {
		context["status"] = fiber.StatusNotFound
		context["error"] = err
		return ctx.JSON(context)
	}

	for _, fileHeaders := range form.File {
		for _, fileHeader := range fileHeaders {
			filePath := "./static/uploads/" + fileHeader.Filename
			if err := ctx.SaveFile(fileHeader, filePath); err != nil {
				context["status"] = fiber.StatusBadRequest
				context["error"] = err
				return ctx.JSON(context)
			}
			product.Photos = append(product.Photos, filePath)
		}
	}

	if err := ctx.BodyParser(&product); err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		return ctx.JSON(context)
	}

	msg, err := database.InsertProduct(connection, &product)
	if err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		return ctx.JSON(context)
	}
	context["status"] = fiber.StatusOK
	context["insert"] = msg
	return ctx.JSON(context)
}

func RemoveProduct(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	context := fiber.Map{}
	id, err := ctx.ParamsInt("id")

	if err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		return ctx.JSON(context)
	}

	msg, err := database.DeleteProduct(connection, id)
	if err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		return ctx.JSON(context)
	}
	context["status"] = fiber.StatusOK
	context["delete"] = msg
	return ctx.JSON(context)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	context := fiber.Map{}
	id, err := ctx.ParamsInt("id")
	product := types.ProductCard{}

	if err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		context["message"] = "Can't parse params in url"
		return ctx.JSON(context)
	}

	//form, err := ctx.MultipartForm()
	//
	//if err != nil {
	//	context["status"] = fiber.StatusNotFound
	//	context["error"] = err
	//	context["message"] = "Can't parse form entries from binary"
	//	return ctx.JSON(context)
	//}
	//
	//for _, fileHeaders := range form.File {
	//	for _, fileHeader := range fileHeaders {
	//		filePath := "./static/uploads/" + fileHeader.Filename
	//		if err := ctx.SaveFile(fileHeader, filePath); err != nil {
	//			context["status"] = fiber.StatusBadRequest
	//			context["error"] = err
	//			return ctx.JSON(context)
	//		}
	//		product.Photos = append(product.Photos, filePath)
	//	}
	//}

	if err := ctx.BodyParser(&product); err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		context["message"] = "Can't parse request body"
		return ctx.JSON(context)
	}

	msg, err := database.ModifyProduct(connection, &product, id)
	if err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		context["message"] = "Can't update product"
		return ctx.JSON(context)
	}
	context["status"] = fiber.StatusOK
	context["update"] = msg
	context["message"] = "Product updated successfully"
	return ctx.JSON(context)
}

func Authorization(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	user := types.User{}
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}
	msg, err := database.VerifyUser(connection, &user)
	if err != nil {
		return err
	}

	var context bool
	if "SELECT 1" == msg {
		context = true
	}
	return ctx.JSON(context)
}

func ProductDetailSlug(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	context := fiber.Map{}

	slug := ctx.Params("slug")

	product, err := database.SelectBySlugName(connection, slug)

	if err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		return ctx.JSON(context)
	}
	context["status"] = fiber.StatusOK
	context["data"] = product
	return ctx.JSON(context)
}

func CategoryDetailBySlug(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	context := fiber.Map{}

	slug := ctx.Params("slug")

	productSlice, err := database.SelectBySlugCategory(connection, slug)
	if err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		return ctx.JSON(context)
	}
	context["status"] = fiber.StatusOK
	context["data"] = productSlice
	return ctx.JSON(productSlice)
}

func ProductDetailID(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	context := fiber.Map{}
	id, err := ctx.ParamsInt("id")

	if err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		return ctx.JSON(context)
	}

	product, err := database.SelectByID(connection, id)
	if err != nil {
		context["status"] = fiber.StatusBadRequest
		context["error"] = err
		return ctx.JSON(context)
	}
	context["status"] = fiber.StatusOK
	context["data"] = product
	return ctx.JSON(context)
}
