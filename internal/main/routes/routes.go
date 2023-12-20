package routes

import (
	"mongoapi/internal/main/factories"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	accountsFactory := factories.MakeAccountsHandler()

	app.Post("/accounts", accountsFactory.Create)
	app.Get("/accounts/:id", accountsFactory.FindOneById)
}
