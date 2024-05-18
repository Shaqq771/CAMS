package feature

import (
	"backend-nabati/domain/request/constant"
	"backend-nabati/domain/request/model"
	Error "backend-nabati/domain/shared/error"
	"context"
	"errors"
	"strconv"
)

func (lf requestFeature) GetApprovalFeature(ctx context.Context, id string) (response model.Approval, err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = Error.New(constant.ErrGeneral, constant.ErrFailedConvertStringToInt, err)
		return
	}

	result, err := lf.requestRepo.GetApprovalByIdRepository(ctx, idInt)
	if err != nil {
		return
	}

	if result.Id == 0 {
		err = Error.New(constant.ErrGeneral, constant.ErrApprovalIdNotFound, errors.New(strconv.Itoa(result.Id)))
		return
	}

	response = result

	return
}