package routes

import (
	"mongoapi/internal/main/factories"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	accountsHandler := factories.MakeAccounts()

	app.Post("/accounts", accountsHandler.CreateAccount)
}
