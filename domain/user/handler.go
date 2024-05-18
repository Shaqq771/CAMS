package user

import (
	"backend-nabati/domain/shared/context"
	Error "backend-nabati/domain/shared/error"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/shared/response"
	"backend-nabati/domain/user/constant"
	"backend-nabati/domain/user/feature"
	"strconv"
	"strings"

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

func (lh userHandler) GetApproverListsHandler(c *fiber.Ctx) error {

	ctx, cancel := context.CreateContextWithTimeout()
	defer cancel()
	ctx = context.SetValueToContext(ctx, c)

	page, err := strconv.Atoi(strings.TrimSpace(c.Query(constant.PAGE)))
	if err != nil || page == 0 {
		page = constant.DefaultPage
	}

	limit, err := strconv.Atoi(strings.TrimSpace(c.Query(constant.LIMIT)))
	if err != nil || limit == 0 {
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

	resp, err := lh.feature.GetApproverListsFeature(ctx, queryRequest)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.MsgGetListsDataSuccess, resp)
}

func (lh userHandler) GetApproverListsWithFilterHandler(c *fiber.Ctx) error {

	ctx := context.CreateContext()
	ctx = context.SetValueToContext(ctx, c)

	filterRequest := new(shared_model.Filter)
	if err := c.BodyParser(filterRequest); err != nil {
		err := Error.New(constant.ErrInvalidRequest, constant.ErrInvalidRequest, err)
		return response.ResponseErrorWithContext(ctx, err)
	}

	resp, err := lh.feature.GetListsApproverWithFilters(ctx, filterRequest)
	if err != nil {
		return response.ResponseErrorWithContext(ctx, err)
	}

	return response.ResponseOK(c, constant.SUCCESS, resp)
}
