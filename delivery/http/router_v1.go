package http

import (
	"backend-nabati/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouterGroupV1(app *fiber.App, handler handler) {

	v1 := app.Group("/v1")
	{
		v1.Get("/health", handler.healthHandler.ServiceHealth)
		v1.Get("/ping", handler.healthHandler.Ping)
	}

	test := v1.Group("test")
	{
		test.Post("/bulk-insert-counter", handler.logistikHandler.BulkCounterHandler)
		test.Post("/get-list-product", handler.logistikHandler.GetProductListsWithFilterHandler)
	}

	pubProduct := v1.Group("/product")
	{
		pubProduct.Get("/get/:id", handler.logistikHandler.GetProductHandler)
		pubProduct.Get("/lists", handler.logistikHandler.GetProductListsHandler)
	}

	// approval := v1.Group("/request")
	// {
	// 	approval.Get("/get/:id", handler.requestHandler.GetApprovalHandler)
	// 	approval.Get("/lists", handler.requestHandler.GetApprovalListsHandler)
	// }

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

	request := v1.Group("/request")
	{
		request.Get("", handler.requestHandler.GetApprovalListsHandler)
	}

	//route > http handler > feature > repository

}
