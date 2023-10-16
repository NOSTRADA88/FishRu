package main

import (
	"FishRu/database"
	"FishRu/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Can't load .env file: %s", err)
	}

	database.Init()

	app := fiber.New(fiber.Config{DisablePreParseMultipartForm: true, StreamRequestBody: true})
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Static("/static", "./static")

	router.SetUpRouters(app)

	if err := app.Listen(":1337"); err != nil {
		log.Fatalf("Can't run server: %s", err)
	}

}
