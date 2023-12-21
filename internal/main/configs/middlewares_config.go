package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

func RequestIdConfig() requestid.Config {
	return requestid.Config{
		Next:       nil,
		Header:     fiber.HeaderXRequestID,
		Generator:  func() string { return uuid.New().String() },
		ContextKey: "requestid",
	}
}
