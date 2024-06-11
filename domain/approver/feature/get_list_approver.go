package feature

import (
	"backend-nabati/domain/approver/model"
	"context"
)

func (af approverFeature) GetListOfApproverFeature(ctx context.Context) (response model.ApproverListNoFilter, err error) {
	listApprover, err := af.approverRepository.GetListOfApproverRepository(ctx)
	if err != nil {
		return
	}
	response = model.ApproverListNoFilter{
		Approver: listApprover,
	}

	return
}
