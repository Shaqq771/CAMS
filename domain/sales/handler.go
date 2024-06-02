package sales

import (
	"fmt"

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
	fmt.Println("masuk sini")
	return nil
}
