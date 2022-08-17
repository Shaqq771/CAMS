package http

import (
	"backend-nabati/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func ServerHttp(handler handler) *fiber.App {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("server up!")
	})

	v1 := app.Group("/v1")
	{
		v1.Get("/health", handler.healthHandler.ServiceHealth)
	}

	pubProduct := v1.Group("/product")
	{
		pubProduct.Get("/get/:id", handler.logistikHandler.GetProductHandler)
		pubProduct.Get("/lists", handler.logistikHandler.GetProductListsHandler)
	}

	authProduct := v1.Group("/product")
	{
		authProduct.Use(middleware.AuthValidations())
		authProduct.Post("/add", handler.logistikHandler.AddProductHandler)
		authProduct.Put("/update/:id", handler.logistikHandler.UpdateProductHandler)
		authProduct.Delete("/delete/:id", handler.logistikHandler.DeleteProductHandler)
	}

	sales := v1.Group("/sales")
	{
		sales.Get("/health", handler.salesHandler.HealthCheck)
	}

	return app
}
