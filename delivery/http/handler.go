package http

import (
	"backend-nabati/delivery/container"
	"backend-nabati/domain/approver"
	business "backend-nabati/domain/businessunit"
	"backend-nabati/domain/health"
	"backend-nabati/domain/logistik"
	"backend-nabati/domain/module"
	"backend-nabati/domain/request"
	"backend-nabati/domain/sales"
	"backend-nabati/domain/user"
)

type handler struct {
	healthHandler   health.HealthHandler
	logistikHandler logistik.LogistikHandler
	salesHandler    sales.SalesHandler
	userHandler     user.UserHandler
	requestHandler  request.RequestHandler
	businessHandler business.BusinessHandler
	moduleHandler   module.ModuleHandler
	approverHandler approver.ApproverHandler
}

func SetupHandler(container container.Container) handler {
	return handler{
		healthHandler:   health.NewHealthHandler(container.HealthFeature),
		logistikHandler: logistik.NewLogistikHandler(container.LogistikFeature),
		salesHandler:    sales.NewSalesHandler(),
		userHandler:     user.NewUserHandler(container.UserFeature),
		requestHandler:  request.NewRequestHandler(container.RequestFeature),
		businessHandler: business.NewBusinessHandler(container.BusinessFeature),
		moduleHandler:   module.NewModuleHandler(container.ModuleFeature),
		approverHandler: approver.NewApproverHandler(container.ApproverFeature),
	}
}
