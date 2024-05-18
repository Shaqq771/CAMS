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

func (lf requestFeature) GetApprovalListsFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (approvalList model.ApprovalLists, err error) {

	// Cleaning & Set Sort query
	var (
		qsortList     []string
		qFilterList   []string
		sortby        = queryRequest.SortBy
		search        = queryRequest.Search
		totalApprovals int
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
		totalApprovals, err = lf.requestRepo.GetTotalApprovalWithConditionsRepository(ctx, search)
		if err != nil {
			return
		}
	} else {
		// Get Total Product Now
		totalApprovals, err = lf.requestRepo.GetTotalApprovalRepository(ctx)
		if err != nil {
			return
		}
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalApprovals, queryRequest.Limit, queryRequest.Page)

	// Get Lists Product
	approvals, err := lf.requestRepo.GetApprovalListsRepository(ctx, queryRequest.Limit, offset, sortby, search)
	if err != nil {
		return
	}

	approvalList = model.ApprovalLists{
		Pagination: shared_model.Pagination{
			Limit:     queryRequest.Limit,
			TotalPage: total_page,
			TotalRows: totalApprovals,
			Page:      queryRequest.Page,
		},
		Approval: approvals,
		Sort:    qsortList,
		Filter:  qFilterList,
	}

	return
}

func (lf requestFeature) GetListsApprovalWithFilters(ctx context.Context, filter *shared_model.Filter) (approvalList model.ApprovalListsByFilter, err error) {

	// Get Total Product Now
	totalApprovals, err := lf.requestRepo.GetTotalApprovalWithFiltersRepository(ctx, filter)
	if err != nil {
		return
	}

	// Set Paginations for product lists
	offset, total_page := helper.GetPaginations(totalApprovals, filter.Limit, filter.Page)

	// Get Lists Product
	approvals, err := lf.requestRepo.GetApprovalListsWithFiltersRepository(ctx, filter, offset)
	if err != nil {
		return
	}

	approvalList = model.ApprovalListsByFilter{
		Pagination: shared_model.Pagination{
			Limit:     filter.Limit,
			TotalPage: total_page,
			TotalRows: totalApprovals,
			Page:      filter.Page,
		},
		Approval: approvals,
		Filters: filter.Filters,
	}

	return
}
