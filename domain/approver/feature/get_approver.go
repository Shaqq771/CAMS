package feature

import (
	"backend-nabati/domain/approver/constant"
	"backend-nabati/domain/approver/model"
	Error "backend-nabati/domain/shared/error"
	"context"
	"strconv"
)

func (af approverFeature) GetApproverFeature(ctx context.Context, id string) (response model.ApproverListNoFilter, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := af.approverRepository.GetApproverByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	var approverIds []int
	for _, approver := range result {
		approverIds = append(approverIds, approver.Id)
	}

	response = model.ApproverListNoFilter{
		Approver: result,
	}
	return
}
