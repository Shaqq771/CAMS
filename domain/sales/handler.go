package sales

import (
	"github.com/gofiber/fiber/v2"
)

type SalesHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type salesHandler struct {
}

func NewSalesHandler() SalesHandler {
	return &salesHandler{}
}

func (sh salesHandler) HealthCheck(c *fiber.Ctx) error {

	return nil
}
