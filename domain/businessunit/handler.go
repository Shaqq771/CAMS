package business

import (
	"backend-nabati/domain/businessunit/constant"
	"backend-nabati/domain/businessunit/feature"
	"backend-nabati/domain/shared/context"
	"backend-nabati/domain/shared/response"

	"github.com/gofiber/fiber/v2"
)

type BusinessHandler interface {
	GetBusinessListsHandler(c *fiber.Ctx) error
}

type businessHandler struct {
	feature feature.BusinessFeature
}

func NewBusinessHandler(feature feature.BusinessFeature) BusinessHandler {
	return &businessHandler{
		feature: feature,
	}
}

func (bh businessHandler) GetBusinessListsHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
	results, err := bh.feature.GetListOfBusinessFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetBusinessSuccess, results)
}