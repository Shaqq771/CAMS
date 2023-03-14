package http

import (
	"backend-nabati/delivery/container"

	"github.com/gofiber/fiber/v2"
)

func ServeHttp(container container.Container) *fiber.App {
	handler := SetupHandler(container)

	app := fiber.New()

	// iniate router v1
	RouterGroupV1(app, handler)

	return app
}
