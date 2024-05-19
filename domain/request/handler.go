package request

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/feature"
	"backend-nabati/domain/shared/context"
	"backend-nabati/domain/shared/response"

	"github.com/gofiber/fiber/v2"
)

type RequestHandler interface {
	GetRequestListsHandler(c *fiber.Ctx) error
}

type requestHandler struct {
	feature feature.RequestFeature
}

func NewRequestHandler(feature feature.RequestFeature) RequestHandler {
	return &requestHandler{
		feature: feature,
	}
}

func (rh requestHandler) GetRequestListsHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
	results, err := rh.feature.GetListOfRequestFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApprovalSuccess, results)
}
