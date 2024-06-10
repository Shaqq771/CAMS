package feature

import (
	Error "backend-nabati/domain/shared/error"
	"backend-nabati/domain/user/constant"
	"backend-nabati/domain/user/model"
	"context"
	"strconv"
)

func (uf userFeature) GetApproverFeature(ctx context.Context, id string) (response model.ApproverListNoFilter, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := uf.userRepo.GetApproverByIdRepository(ctx, idInt)
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
