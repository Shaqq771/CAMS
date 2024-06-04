package feature

import (
	"backend-nabati/domain/user/model"
	"context"
)

func (uf userFeature) GetListOfApproverFeature(ctx context.Context) (response model.ApproverListNoFilter, err error) {
	listApprover, err := uf.userRepo.GetListOfApproverRepository(ctx)
	if err != nil {
		return
	}
	response = model.ApproverListNoFilter{
		Approver: listApprover,
	}

	return
}