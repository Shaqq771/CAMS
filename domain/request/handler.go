package request

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/feature"
	"backend-nabati/domain/request/model"
	"backend-nabati/domain/shared/context"
	Error "backend-nabati/domain/shared/error"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/shared/response"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type RequestHandler interface {
	GetRequestListsHandler(c *fiber.Ctx) error
	GetRequestHandler(c *fiber.Ctx) error
	GetRequestFilterHandler(c *fiber.Ctx) error
	GetRequestListsWaitingHandler(c *fiber.Ctx) error
	GetRequestListsApprovedHandler(c *fiber.Ctx) error
	GetRequestListsRejectedHandler(c *fiber.Ctx) error
	GetRequestListsRevisedHandler(c *fiber.Ctx) error
	UpdateRequestHandler(c *fiber.Ctx) error
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

func (rh requestHandler) GetRequestListsWaitingHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
	results, err := rh.feature.GetRequestListsWaitingFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApprovalSuccess, results)
}

func (rh requestHandler) GetRequestListsApprovedHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
	results, err := rh.feature.GetRequestListsApprovedFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApprovalSuccess, results)
}

func (rh requestHandler) GetRequestListsRejectedHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
	results, err := rh.feature.GetRequestListsRejectedFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApprovalSuccess, results)
}

func (rh requestHandler) GetRequestListsRevisedHandler(c *fiber.Ctx) error {
	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)
	results, err := rh.feature.GetRequestListsRevisedFeature(ctx)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetApprovalSuccess, results)
}

func (rh requestHandler) GetRequestFilterHandler(c *fiber.Ctx) error {
	fmt.Println(c)

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	page, err := strconv.Atoi(strings.TrimSpace(c.Query(constant.PAGE)))
	if err != nil || page == 0 {
		fmt.Println(err)
		page = constant.DefaultPage
	}
	limit, err := strconv.Atoi(strings.TrimSpace(c.Query(constant.LIMIT)))
	if err != nil || limit == 0 {
		fmt.Println(err)
		limit = constant.DefaultLimitPerPage
	}

	sortBy := strings.TrimSpace(c.Query(constant.SORT_BY))
	search := strings.TrimSpace(c.Query(constant.SEARCH))

	queryRequest := shared_model.QueryRequest{
		Page:   page,
		Limit:  limit,
		SortBy: sortBy,
		Search: search,
	}

	resp, err := rh.feature.GetRequestFilterFeature(ctx, queryRequest)
	if err != nil {
		fmt.Println(err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetListsDataSuccess, resp)
}

func (rh requestHandler) UpdateRequestHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	id := c.Params("id")
	if id == "" || id == "0" {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrApprovalIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	request := new(model.UpdateApprovalRequest)
	if err := c.BodyParser(request); err != nil {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, fmt.Errorf(constant.ErrApprovalIdNil))
		return response.ResponseErrorWithContext(ctx, err)
	}

	results, err := rh.feature.UpdateRequestFeature(ctx, id, request)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgUpdateApprovalSuccess, results)
}