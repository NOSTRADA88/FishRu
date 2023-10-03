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
	// if we need to drop table
	//if err := database.DropTable(connection); err != nil {
	//	log.Fatalf("Can't drop table: %s", err)
	//}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.SetUpRouters(app)

	if err := app.Listen(":1337"); err != nil {
		log.Fatalf("Can't run server: %s", err)
	}

}
