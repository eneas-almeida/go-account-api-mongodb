package handlers

import (
	"encoding/json"
	"mongoapi/internal/domain/accounts"

	"github.com/gofiber/fiber/v2"
)

type AccountsHandler struct {
	UseCases accounts.UseCasesInterface
}

func (h *AccountsHandler) CreateAccount(ctx *fiber.Ctx) error {
	body := []byte(ctx.Body())

	var dto accounts.InputAccountDto

	json.Unmarshal(body, &dto)

	_, err := h.UseCases.Create(&dto)

	if err != nil {
		return ctx.Status(400).JSON(err)
	}

	return nil
}
