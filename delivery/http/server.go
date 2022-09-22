package http

import (
	"github.com/gofiber/fiber/v2"
)

func ServeHttp(handler handler) *fiber.App {

	app := fiber.New()

	// iniate router v1
	RouterGroupV1(app, handler)

	return app
}
