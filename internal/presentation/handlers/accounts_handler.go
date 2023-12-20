package handlers

import (
	"mongoapi/internal/domain/accounts"

	"github.com/gofiber/fiber/v2"
)

type AccountsHandler struct {
	UseCases accounts.UseCasesInterface
}

func (h *AccountsHandler) CreateAccount(c *fiber.Ctx) error {
	return nil
}
