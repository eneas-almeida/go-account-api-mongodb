package main

import (
	"mongoapi/internal/main/factories"
)

func main() {
	var accountsFactory factories.AccountsFactory

	accountsFactory.MakeAccounts()

	// app := fiber.New()

	// app.Get("/login/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	_, err := configs.GetMongoDbConnection()

	// 	if err != nil {
	// 		log.Fatal("error")
	// 	}

	// 	return c.SendString("idxx=" + id)
	// })

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })

	// app.Listen(":3000")
}
