package http

import (
	"backend-nabati/delivery/container"
	"backend-nabati/domain/health"
	"backend-nabati/domain/logistik"
	"backend-nabati/domain/sales"
)

type handler struct {
	healthHandler   health.HealthHandler
	logistikHandler logistik.LogistikHandler
	salesHandler    sales.SalesHandler
}

func SetupHandler(container container.Container) handler {
	return handler{
		healthHandler:   health.NewHealthHandler(container.HealthFeature),
		logistikHandler: logistik.NewLogistikHandler(container.LogistikFeature),
		salesHandler:    sales.NewSalesHandler(),
	}
}
