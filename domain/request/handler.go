package request

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/feature"
	"backend-nabati/domain/shared/context"
	"backend-nabati/domain/shared/response"
	"fmt"

	Error "backend-nabati/domain/shared/error"

	"github.com/gofiber/fiber/v2"
)

type RequestHandler interface {
	GetRequestListsHandler(c *fiber.Ctx) error
	GetRequestHandler(c *fiber.Ctx) error

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

func (rh requestHandler) GetRequestHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrApprovalIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := rh.feature.GetRequestFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApprovalSuccess, results)
}
