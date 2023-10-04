package controller

import (
	"FishRu/database"
	"FishRu/types"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"strings"
)

func ProductList(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	ctx.Status(200)
	prodSlice, err := database.SelectAll(connection)
	if err != nil {
		return err
	}
	return ctx.JSON(prodSlice)
}

func CreateProduct(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	ctx.Status(200)
	product := types.ProductCard{}
	if err := ctx.BodyParser(&product); err != nil {
		return err
	}
	log.Println(&product)
	msg, err := database.InsertProduct(connection, &product)
	if err != nil {
		return err
	}
	msgSlice := strings.Fields(msg)
	context := fiber.Map{
		msgSlice[0]: msgSlice[2],
	}
	return ctx.JSON(context)
}

func RemoveProduct(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	ctx.Status(200)
	product := types.ProductCard{}
	if err := ctx.BodyParser(&product); err != nil {
		return err
	}
	msg, err := database.DeleteProduct(connection, &product)
	if err != nil {
		return err
	}
	msgSlice := strings.Fields(msg)
	context := fiber.Map{
		msgSlice[0]: msgSlice[1],
	}
	return ctx.JSON(context)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	ctx.Status(200)
	product := types.ProductCard{}
	if err := ctx.BodyParser(&product); err != nil {
		return err
	}
	msg, err := database.ModifyProduct(connection, &product)
	if err != nil {
		return err
	}
	msgSlice := strings.Fields(msg)
	context := fiber.Map{
		msgSlice[0]: msgSlice[1],
	}
	return ctx.JSON(context)
}

func Authorization(ctx *fiber.Ctx) error {
	connection := database.Connection(os.Getenv("DB_URL"))
	defer database.CloseConnection(connection)

	ctx.Status(200)
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
