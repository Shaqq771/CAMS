package user

import (
	"backend-nabati/domain/shared/context"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/shared/response"
	"backend-nabati/domain/user/constant"
	"backend-nabati/domain/user/feature"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetApproverListsHandler(c *fiber.Ctx) error
	GetApproverHandler(c *fiber.Ctx) error
}

type userHandler struct {
	feature feature.UserFeature
}

func NewUserHandler(feature feature.UserFeature) UserHandler {
	return &userHandler{
		feature: feature,
	}
}

// func (uh userHandler) GetApproverListsHandler(c *fiber.Ctx) error {

// 	ctx, cancel := context.CreateContextWithTimeout()
// 	defer cancel()
// 	ctx = context.SetValueToContext(ctx, c)
// 	results, err := uh.feature.GetApproverListsFeature(ctx)
// 	if err != nil {
// 		return response.ResponseErrorWithContext(ctx, err)
// 	}

// 	return response.ResponseOK(c, constant.MsgGetApproverSuccess, results)
// }


func (uh userHandler) GetApproverListsHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	results, err := uh.feature.GetListOfApproverFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApproverSuccess, results)
}

func (uh userHandler) GetApproverHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrApproverIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := uh.feature.GetApproverFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApproverSuccess, results)
}
