package rule

import (
	"backend-nabati/domain/rule/constant"
	"backend-nabati/domain/rule/feature"
	"backend-nabati/domain/shared/context"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/shared/response"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RuleHandler interface {
	GetRuleListsHandler(c *fiber.Ctx) error
	GetRuleHandler(c *fiber.Ctx) error
}

type ruleHandler struct {
	feature feature.RuleFeature
}

func NewRuleHandler(feature feature.RuleFeature) RuleHandler {
	return &ruleHandler{
		feature: feature,
	}
}

func (rh ruleHandler) GetRuleListsHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
	results, err := rh.feature.GetListOfRuleFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetRuleSuccess, results)
}

func (rh ruleHandler) GetRuleHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrRuleIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := rh.feature.GetRuleFeature(ctx, id)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetRuleSuccess, results)
}