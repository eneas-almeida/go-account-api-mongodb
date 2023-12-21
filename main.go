package main

import (
	"log"
	"mongoapi/internal/main/configs"
	"mongoapi/internal/main/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	err := configs.InitDB()

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	defer configs.CloseDB()

	app := fiber.New()

	app.Use(cors.New())
	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(limiter.New())
	app.Use(requestid.New(configs.RequestIdConfig()))

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
