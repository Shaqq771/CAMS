package approver

import (
	"backend-nabati/domain/approver/feature"
	"backend-nabati/domain/approver/model"
	"backend-nabati/domain/shared/context"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/shared/response"

	"backend-nabati/domain/approver/constant"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ApproverHandler interface {
	GetApproverListsHandler(c *fiber.Ctx) error
	GetApproverHandler(c *fiber.Ctx) error
	AddApproverHandler(c *fiber.Ctx) error
}

type approverHandler struct {
	feature feature.ApproverFeature
}

func NewApproverHandler(feature feature.ApproverFeature) ApproverHandler {
	return &approverHandler{
		feature: feature,
	}
}

func (ah approverHandler) GetApproverListsHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	results, err := ah.feature.GetListOfApproverFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApproverSuccess, results)
}

func (ah approverHandler) GetApproverHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrApproverIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := ah.feature.GetApproverFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApproverSuccess, results)
}

func (ah approverHandler) AddApproverHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	request := new(model.AddApproverRequest)
	if err := c.BodyParser(request); err != nil {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	} else if request.Name == "" ||
		request.Email == "" ||
		request.Department == "" ||
		request.Job_title == "" ||
		request.Business_unit == "" ||
		request.Approver_user_id == "" {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := ah.feature.AddApproverFeature(ctx, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgInsertDataSuccess, results)
}
