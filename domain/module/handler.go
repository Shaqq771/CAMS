package module

import (
	"backend-nabati/domain/module/constant"
	"backend-nabati/domain/module/feature"
	"backend-nabati/domain/shared/context"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/shared/response"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ModuleHandler interface {
	GetModuleListsHandler(c *fiber.Ctx) error
	GetModuleHandler(c *fiber.Ctx) error
	// GetRequestFilterHandler(c *fiber.Ctx) error
}

type moduleHandler struct {
	feature feature.ModuleFeature
}

func NewModuleHandler(feature feature.ModuleFeature) ModuleHandler {
	return &moduleHandler{
		feature: feature,
	}
}

func (mh moduleHandler) GetModuleListsHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
	results, err := mh.feature.GetListOfModuleFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetModuleSuccess, results)
}

func (mh moduleHandler) GetModuleHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrModuleIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := mh.feature.GetModuleFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetModuleSuccess, results)
}