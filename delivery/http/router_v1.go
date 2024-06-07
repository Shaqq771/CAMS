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
		request.Get("/", handler.requestHandler.GetRequestListsHandler)
		request.Get("/get/:id", handler.requestHandler.GetRequestHandler)
		request.Get("/lists", handler.requestHandler.GetRequestFilterHandler)
		request.Get("/waiting", handler.requestHandler.GetRequestListsWaitingHandler)
		request.Get("/approved", handler.requestHandler.GetRequestListsApprovedHandler)
		request.Get("/rejected", handler.requestHandler.GetRequestListsRejectedHandler)
		request.Get("/revised", handler.requestHandler.GetRequestListsRevisedHandler)

	}

	approver := v1.Group("/approver")
	{
		approver.Get("/", handler.userHandler.GetApproverListsHandler)
		approver.Get("/get/:id", handler.userHandler.GetApproverHandler)
	}

	business := v1.Group("/businessunit")
	{
		business.Get("/", handler.businessHandler.GetBusinessListsHandler)
		business.Get("/get/:id", handler.businessHandler.GetBusinessHandler)
	}

	module := v1.Group("/module")
	{
		module.Get("/", handler.moduleHandler.GetModuleListsHandler)
		module.Get("/get/:id", handler.moduleHandler.GetModuleHandler)
	}
	

	//route > http handler > feature > repository

}
