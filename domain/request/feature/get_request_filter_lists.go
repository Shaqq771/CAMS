package feature

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/model"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/shared/helper"
	shared_model "backend-nabati/domain/shared/model"
	"context"
	"strings"
)

func (rf requestFeature) GetRequestFilterFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (requestList model.RequestLists, err error) {

	// Cleaning & Set Sort query
	var (
		qsortList     []string
		qFilterList   []string
		sortby        = queryRequest.SortBy
		search        = queryRequest.Search
		totalProducts int
	)

	// Sort
	sortby = strings.TrimSpace(sortby)
	if sortby != "" {
		sortby, qsortList, err = helper.SortBy(sortby)
		if err != nil {
			err = Error.New(constant.ErrGeneral, constant.ErrInvalidSortBy, err)
			return
		}
	}

	// Search
	if search != "" {
		search, qFilterList, err = helper.FilterBy(search)
		if err != nil {
			err = Error.New(constant.ErrGeneral, constant.ErrInvalidFilterBy, err)
			return
		}

		// Get Total Product Now
		totalProducts, err = rf.requestRepo.GetTotalRequestWithConditionsRepository(ctx, search)
		if err != nil {
			return
		}
	} else {
		// Get Total Product Now
		totalProducts, err = rf.requestRepo.GetTotalRequestRepository(ctx)
		if err != nil {
			return
		}
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalProducts, queryRequest.Limit, queryRequest.Page)

	// Get Lists Product
	request, err := rf.requestRepo.GetRequestListsRepository(ctx, queryRequest.Limit, offset, sortby, search)
	if err != nil {
		return
	}

	requestList = model.RequestLists{
		Pagination: shared_model.Pagination{
			Limit:     queryRequest.Limit,
			TotalPage: total_page,
			TotalRows: totalProducts,
			Page:      queryRequest.Page,
		},
		Request: request,
		Sort:    qsortList,
		Filter:  qFilterList,
	}

	return
}

func (rf requestFeature) GetListsRequestWithFilters(ctx context.Context, filter *shared_model.Filter) (requestList model.RequestListsByFilter, err error) {

	// Get Total Product Now
	totalRequest, err := rf.requestRepo.GetTotalRequestWithFiltersRepository(ctx, filter)
	if err != nil {
		return
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalRequest, filter.Limit, filter.Page)

	// Get Lists Product
	request, err := rf.requestRepo.GetRequestListsWithFiltersRepository(ctx, filter, offset)
	if err != nil {
		return
	}

	requestList = model.RequestListsByFilter{
		Pagination: shared_model.Pagination{
			Limit:     filter.Limit,
			TotalPage: total_page,
			TotalRows: totalRequest,
			Page:      filter.Page,
		},
		Request: request,
		Filters: filter.Filters,
	}

	return
}
