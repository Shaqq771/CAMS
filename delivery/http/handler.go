package http

import (
	"backend-nabati/delivery/container"
	"backend-nabati/domain/health"
	"backend-nabati/domain/logistik"

	// "backend-nabati/domain/module"
	"backend-nabati/domain/request"
	"backend-nabati/domain/sales"
	"backend-nabati/domain/user"
)

type handler struct {
	healthHandler   health.HealthHandler
	logistikHandler logistik.LogistikHandler
	salesHandler    sales.SalesHandler
	userHandler    user.UserHandler
	requestHandler    request.RequestHandler
	// moduleHandler    module.ModuleHandler
}

func SetupHandler(container container.Container) handler {
	return handler{
		healthHandler:   health.NewHealthHandler(container.HealthFeature),
		logistikHandler: logistik.NewLogistikHandler(container.LogistikFeature),
		salesHandler:    sales.NewSalesHandler(),
	}
}
