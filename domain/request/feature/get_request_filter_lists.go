package feature

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/model"
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/shared/helper"
	shared_model "backend-nabati/domain/shared/model"
	"context"
	"errors"
	"strings"
)

func (rf requestFeature) GetRequestFilterFeature(ctx context.Context, queryRequest shared_model.QueryRequest) (requestList model.RequestLists, err error) {

	// Cleaning & Set Sort query
	var (
		qsortList     []string
		qFilterList   []string
		sortby        = queryRequest.SortBy
		search        = queryRequest.Search
		totalRequests int
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

		// Get Total Request Now
		totalRequests, err = rf.requestRepo.GetTotalRequestWithConditionsRepository(ctx, search)
		if err != nil {
			return
		}
	} else {
		// Get Total Request Now
		totalRequests, err = rf.requestRepo.GetTotalRequestRepository(ctx)
		if err != nil {
			return
		}
	}

	// Set Paginations for request lists
	offset, total_page := helper.GetPaginations(totalRequests, queryRequest.Limit, queryRequest.Page)

	// Get Lists Product
	request, err := rf.requestRepo.GetRequestListsRepository(ctx, queryRequest.Limit, offset, sortby, search)
	if err != nil {
		return
	}

	requestList = model.RequestLists{
		Pagination: shared_model.Pagination{
			Limit:     queryRequest.Limit,
			TotalPage: total_page,
			TotalRows: totalRequests,
			Page:      queryRequest.Page,
		},
		Request: request,
		Sort:    qsortList,
		Filter:  qFilterList,
	}

	return
}

func (rf requestFeature) GetListsRequestWithFilters(ctx context.Context, filter *shared_model.Filter) (requestList model.RequestListsByFilter, err error) {

	// Get Total Request Now
	totalRequest, err := rf.requestRepo.GetTotalRequestWithFiltersRepository(ctx, filter)
	if err != nil {
		return
	}

	// Set Paginations for request lists
	offset, total_page := helper.GetPaginations(totalRequest, filter.Limit, filter.Page)

	// Get Lists Request
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

// GetApprovalRequestStatsFeature retrieves request counts for each approval status
func (rf requestFeature) GetApprovalRequestStatsFeature(ctx context.Context) (map[string]int, error) {
	// Initialize an empty map to store status counts
	stats := map[string]int{
	  constant.ApprovalStatusWaiting: 0,
	  constant.ApprovalStatusApproved: 0,
	  constant.ApprovalStatusRevised: 0,
	  constant.ApprovalStatusRejected: 0,
	}
  
	// Perform queries to retrieve request counts for each status
	var err error
  
	// Call repository methods to get request counts by status
	for status, count := range stats {
	  count, err = rf.requestRepo.GetTotalRequestByStatusRepository(ctx, status)
	  if err != nil {
		return nil, errors.New("failed to retrieve request count for status " + status + ": " + err.Error())
	  }
	  stats[status] = count
	}
  
	// Return the map containing status counts
	return stats, nil
  }
  
  


