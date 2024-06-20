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
	UpdateDelegateStatusHandler(c *fiber.Ctx) error
	UpdateSkipStatusHandler(c *fiber.Ctx) error
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
		fmt.Println(err, "err")
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	} else if request.Name == "" ||
		request.Email == "" ||
		request.Department == "" ||
		request.JobTitle == "" ||
		request.BusinessUnit == "" ||
		request.ApproverUserId == 0 ||
		request.Role == "" ||
		request.Location == "" ||
		request.Description == "" {
		err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := ah.feature.AddApproverFeature(ctx, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgInsertDataSuccess, results)
}

func (ah approverHandler) UpdateDelegateStatusHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
  
	// Extract approver ID from request path
	id := c.Params("id")
	if id == "" || id == "0" {
	  err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrApproverIdNil))
	  return response.ResponseErrorWithContext(ctx, err)
	}
  
	// Parse request body for delegate status
	var request struct {
	  IsDelegate bool `json:"is_delegate"`
	}
	if err := c.BodyParser(&request); err != nil {
	  err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
	  return response.ResponseErrorWithContext(ctx, err)
	}
  
	// Call feature to update delegate status
	err := ah.feature.UpdateDelegateStatusFeature(ctx, id, request.IsDelegate)
	if err != nil {
	  return response.ResponseErrorWithContext(ctx, err)
	}
  
	// Success response
	return response.ResponseOK(c, constant.MsgUpdateSuccess, nil)
  }

  func (ah approverHandler) UpdateSkipStatusHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
  
	// Extract approver ID from request path
	id := c.Params("id")
	if id == "" || id == "0" {
	  err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrApproverIdNil))
	  return response.ResponseErrorWithContext(ctx, err)
	}
  
	// Parse request body for skip status
	var request struct {
	  IsSkip bool `json:"is_skip"`
	}
	if err := c.BodyParser(&request); err != nil {
	  err = Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
	  return response.ResponseErrorWithContext(ctx, err)
	}
  
	// Call feature to update delegate status
	err := ah.feature.UpdateSkipStatusFeature(ctx, id, request.IsSkip)
	if err != nil {
	  return response.ResponseErrorWithContext(ctx, err)
	}
  
	// Success response
	return response.ResponseOK(c, constant.MsgUpdateSuccess, nil)
  }