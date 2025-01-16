package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type HealthWebHandler struct {}

func (h HealthWebHandler) Readiness(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusOK)
}

func (h HealthWebHandler) Liveness(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusOK)
}

func NewHealthWebHandler() HealthWebHandler {
	return HealthWebHandler{}
}

func InitializeHealthWeb() HealthWebHandler {
	healthWebHandler := NewHealthWebHandler()
	return healthWebHandler
}
