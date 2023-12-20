package routes

import (
	"mongoapi/internal/main/factories"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	accountsFactory := factories.MakeAccounts()

	app.Post("/accounts", accountsFactory.CreateHandle)
}
