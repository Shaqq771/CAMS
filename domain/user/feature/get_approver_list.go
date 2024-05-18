package feature

import (
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/shared/helper"
	shared_model "backend-nabati/domain/shared/model"
	"backend-nabati/domain/user/constant"
	"backend-nabati/domain/user/model"
	"context"
	"strings"
)

func (lf userFeature) GetApproverListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (approverList model.ApproverLists, err error) {

	// Cleaning & Set Sort query
	var (
		qsortList     []string
		qFilterList   []string
		sortby        = queryRequest.SortBy
		search        = queryRequest.Search
		totalApprovers int
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
		totalApprovers, err = lf.userRepo.GetTotalApproverWithConditionsRepository(ctx, search)
		if err != nil {
			return
		}
	} else {
		// Get Total Product Now
		totalApprovers, err = lf.userRepo.GetTotalApproverRepository(ctx)
		if err != nil {
			return
		}
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalApprovers, queryRequest.Limit, queryRequest.Page)

	// Get Lists Product
	approvers, err := lf.userRepo.GetApproverListsRepository(ctx, queryRequest.Limit, offset, sortby, search)
	if err != nil {
		return
	}

	approverList = model.ApproverLists{
		Pagination: shared_model.Pagination{
			Limit:     queryRequest.Limit,
			TotalPage: total_page,
			TotalRows: totalApprovers,
			Page:      queryRequest.Page,
		},
		Approver: approvers,
		Sort:    qsortList,
		Filter:  qFilterList,
	}

	return
}

func (lf userFeature) GetListsApproverWithFilters(ctx context.Context, filter *shared_model.Filter) (approverList model.ApproverListsByFilter, err error) {

	// Get Total Product Now
	totalApprovers, err := lf.userRepo.GetTotalApproverWithFiltersRepository(ctx, filter)
	if err != nil {
		return
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalApprovers, filter.Limit, filter.Page)

	// Get Lists Product
	approvers, err := lf.userRepo.GetApproverListsWithFiltersRepository(ctx, filter, offset)
	if err != nil {
		return
	}

	approverList = model.ApproverListsByFilter{
		Pagination: shared_model.Pagination{
			Limit:     filter.Limit,
			TotalPage: total_page,
			TotalRows: totalApprovers,
			Page:      filter.Page,
		},
		Approver: approvers,
		Filters: filter.Filters,
	}

	return
}
