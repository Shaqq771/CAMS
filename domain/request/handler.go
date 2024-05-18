package request

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/feature"
	"backend-nabati/domain/shared/context"
	Error "backend-nabati/domain/shared/error"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/shared/response"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RequestHandler interface {
	GetApprovalListsHandler(c *fiber.Ctx) error
	GetApprovalHandler(c *fiber.Ctx) error
	GetApprovalListsWithFilterHandler(c *fiber.Ctx) error
}

type requestHandler struct {
	feature feature.RequestFeature
}

func NewRequestHandler(feature feature.RequestFeature) RequestHandler {
	return &requestHandler{
		feature: feature,
	}
}

func (lh requestHandler) GetApprovalListsHandler(c *fiber.Ctx) error {
	nil
}

func (lh requestHandler) GetApprovalHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrApprovalIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := lh.feature.GetApprovalFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApprovalSuccess, results)
}

func (lh requestHandler) GetApprovalListsWithFilterHandler(c *fiber.Ctx) error {

	ctx := context.CreateContext()
	ctx = context.SetValueToContext(ctx, c)

	filterRequest := new(shared_model.Filter)
	if err := c.BodyParser(filterRequest); err != nil {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	resp, err := lh.feature.GetListsApprovalWithFilters(ctx, filterRequest)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}
