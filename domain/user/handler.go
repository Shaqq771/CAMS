package user

import (
	"backend-nabati/domain/shared/context"
	"backend-nabati/domain/shared/response"
	"backend-nabati/domain/user/constant"
	"backend-nabati/domain/user/feature"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetApproverListsHandler(c *fiber.Ctx) error
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
