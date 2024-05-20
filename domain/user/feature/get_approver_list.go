package feature

// import (
// 	"backend-nabati/domain/shared/helper"
// 	shared_model "backend-nabati/domain/shared/model"
// 	"backend-nabati/domain/user/model"
// 	"context"
// )

// // func (uf userFeature) GetApproverListsFeature(ctx context.Context) (approverList model.ApproverLists, err error) {

// // 	listRequest, err := uf.userRepo.GetApproverListsRepository(ctx)
// // 	if err != nil {
// // 		return
// // 	}
// // 	approverList = model.RequestListNoFilter{
// // 		Request: listRequest,
// // 	}

// // 	return
// // }

// func (lf userFeature) GetListsApproverWithFilters(ctx context.Context, filter *shared_model.Filter) (approverList model.ApproverListsByFilter, err error) {

// 	// Get Total Product Now
// 	totalApprovers, err := lf.userRepo.GetTotalApproverWithFiltersRepository(ctx, filter)
// 	if err != nil {
// 		return
// 	}

// 	// Set Paginations for product lists
// 	offset, total_page := helper.GetPaginations(totalApprovers, filter.Limit, filter.Page)

// 	// Get Lists Product
// 	approvers, err := lf.userRepo.GetApproverListsWithFiltersRepository(ctx, filter, offset)
// 	if err != nil {
// 		return
// 	}

// 	approverList = model.ApproverListsByFilter{
// 		Pagination: shared_model.Pagination{
// 			Limit:     filter.Limit,
// 			TotalPage: total_page,
// 			TotalRows: totalApprovers,
// 			Page:      filter.Page,
// 		},
// 		Approver: approvers,
// 		Filters: filter.Filters,
// 	}

// 	return
// }
